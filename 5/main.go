package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

//Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).
//
//Реализовать поддержку утилитой следующих ключей:
//-A - "after" печатать +N строк после совпадения
//-B - "before" печатать +N строк до совпадения
//-C - "context" (A+B) печатать ±N строк вокруг совпадения
//-c - "count" (количество строк)
//-i - "ignore-case" (игнорировать регистр)
//-v - "invert" (вместо совпадения, исключать)
//-F - "fixed", точное совпадение со строкой, не паттерн
//-n - "line num", напечатать номер строки

type Flags struct {
	A int
	B int
	C int
	c bool
	i bool
	v bool
	F bool
	n bool
}

//func main() {
//
//	A := flag.Int("A", 0, "")
//	B := flag.Int("B", 0, "")
//	C := flag.Int("C", 0, "")
//}

func main() {

	var name string

	flagA := flag.Int("A", 0, "")
	flagB := flag.Int("B", 0, "")
	flagC := flag.Int("C", 0, "")
	flagc := flag.Bool("c", false, "")
	flagi := flag.Bool("i", false, "")
	flagv := flag.Bool("v", false, "")
	flagF := flag.Bool("F", false, "")
	flagn := flag.Bool("n", false, "")

	flag.StringVar(&name, "name", "test.txt", "file-path")

	flag.Parse()

	f := Flags{
		A: *flagA,
		B: *flagB,
		C: *flagC,
		c: *flagc,
		i: *flagi,
		v: *flagv,
		F: *flagF,
		n: *flagn,
	}

	//fmt.Print(f)

	data := readFile(name, f.i, f.F)

	fmt.Println(data)
}

//Чтение файла
func readFile(name string, ig, F bool) (data []string) {

	file, err := os.OpenFile(name, os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("Ошибка открытия файла")
	}
	defer file.Close()
	buf := bufio.NewReader(file)

	for {
		line, err := buf.ReadString('\n')
		fmt.Println(line)
		fmt.Println(err)

		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalf("Ошибка чтения файла: %v\n", err)
				return
			}
		}
	}

	return
}
