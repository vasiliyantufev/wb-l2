package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	signal "os/signal"
	"strings"
	syscall "syscall"
	"time"
	
	"github.com/reiver/go-telnet"
)

/*Реализовать простейший telnet-клиент.

Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Требования:
Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP. После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться. При подключении к несуществующему сервер, программа должна завершаться через timeout
*/

func main() {
	var host, port string
	var conn *telnet.Conn
	var err error
	
	timeout, err := time.ParseDuration(os.Args[2])
	if err != nil {
		log.Fatalln("Error timeout parsing", err)
	}
	
	host = os.Args[6]
	port = os.Args[7]
	
	fmt.Print(host)
	fmt.Print(port)
	fmt.Print(timeout)
	
	over := time.After(timeout)
	
	completed := make(chan bool)
	
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	
	go func() {
		for {
			select {
			case <-over:
				log.Println("TimeOut")
				completed <- true
				return
			case <-done:
				log.Println("system event")
				completed <- true
				return
			default:
				out := make([]string, 0, 4)
				conn, err = telnet.DialTo(os.Args[6] + ":" + os.Args[7])
				if err != nil {
					log.Println("dial error", err)
					continue
				}
				defer conn.Close()
				
				buf := bufio.NewReader(conn)
				for {
					line, err := buf.ReadBytes(byte('\n'))
					if err != nil {
						if err == io.EOF {
							break
						} else {
							log.Printf("Error socet : %v\n", err)
							return
						}
					}
					out = append(out, string(line))
					
				}
				
				fmt.Println(strings.Join(out, ""))
				completed <- true
				
			}
		}
	}()
	<-completed
}
