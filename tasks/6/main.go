package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"bufio"
)

//Реализовать утилиту аналог консольной команды cut (man cut).
//Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB)
//на колонки и выводить запрошенные.
//
//Реализовать поддержку утилитой следующих ключей:
//-f - "fields" - выбрать поля (колонки)
//-d - "delimiter" - использовать другой разделитель
//-s - "separated" - только строки с разделителем

const (
	separate = "\t"
	column   = 1
)

type Flags struct {
	field int
	separated bool
	delimiter string
	name string
}

func main() {

	flagField := flag.Int("f", column, "fields")
	flagDelimiter := flag.String("d", separate, "delimiter")
	flagSeparated := flag.Bool("s", false, "separated")
	flagName := flag.String("name", "test.txt", "")

	flag.Parse()

	flag := Flags{
		field: *flagField,
		delimiter: *flagDelimiter,
		separated: *flagSeparated,
		name: *flagName,
	}

	fileData := readFile(flag.name)

    //fmt.Println(fileData)

 	cut := Cut(fileData, flag.field, flag.delimiter, flag.separated)
 	Out(cut)
}

//Чтение файла*/
func readFile(name string) (data []string) {

	file, err := os.Open(name)
	if err != nil {
		log.Fatalf("Ошибка открытия файла")
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalf("Ошибка чтения файла: %v\n", err)
				return
			}
		}
		data = append(data, strings.TrimSpace(line))
	}
	return
}

func Cut(data []string, fields int, delimiter string, separated bool) ([]string) {

    var d []string
	for _, line := range data {

		cols := strings.Split(line, delimiter)

				if fields > 0 && fields <= len(cols) {
        			//fmt.Println(cols[fields-1])
        			d = append(d, cols[fields-1])
        		} else if !separated {
        			//fmt.Println(line)
        			d = append(d, line)
        		}
	}

	//fmt.Println(d)
	return d
}

//Output method. [3]
func Out(data []string) {
	for _, val := range data {
		fmt.Println(val)
	}
}