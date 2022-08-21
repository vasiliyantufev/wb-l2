package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

//Реализовать утилиту фильтрации по аналогии с консольной утилитой
//(man grep — смотрим описание и основные параметры).
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
	A    int
	B    int
	C    int
	c    bool
	i    bool
	v    bool
	F    bool
	n    bool
	name string
}

//func main() {
//
//	A := flag.Int("A", 0, "")
//	B := flag.Int("B", 0, "")
//	C := flag.Int("C", 0, "")
//}

func main() {

	var /*name,*/ search string

	for {
		data := make([]byte, 1)
		_, err := os.Stdin.Read(data)
		search += string(data)
		if err != nil {
			break
		}
	}

	flagA := flag.Int("A", 0, "")
	flagB := flag.Int("B", 0, "")
	flagC := flag.Int("C", 0, "")
	flagc := flag.Bool("c", false, "")
	flagi := flag.Bool("i", false, "")
	flagv := flag.Bool("v", false, "")
	flagF := flag.Bool("F", false, "")
	flagn := flag.Bool("n", false, "")

	flagName := flag.String("name", "test.txt", "")

	flag.Parse()

	f := Flags{
		A:    *flagA,
		B:    *flagB,
		C:    *flagC,
		c:    *flagc,
		i:    *flagi,
		v:    *flagv,
		F:    *flagF,
		n:    *flagn,
		name: *flagName,
	}

	//fmt.Println(f.name)

	fileData := readFile(f.name)

	if f.A != 0 {
		//fmt.Println("a")
		//fmt.Print(modeA(fileData, search, f.A))
		modeA(fileData, search, f.A)
	} else if f.B != 0 {
		//fmt.Println("b")
		modeB(fileData, search, f.B)
	} else if f.C != 0 {
		//fmt.Println("c")
		modeC(fileData, search, f.C)
	}
}

//Чтение файла
func readFile(name string) (data []string) {

	file, err := os.OpenFile(name, os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("Ошибка открытия файла")
	}
	defer file.Close()

	data = make([]string, 0, 4)
	buf := bufio.NewReader(file)

	for {
		line, err := buf.ReadString('\n')
		//fmt.Println(line)
		//fmt.Println(err)

		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalf("Ошибка чтения файла: %v\n", err)
				return
			}
		}
		//data :=
		data = append(data, line)
	}
	return data
}

//Вывод N строк после совпадения
func modeA(data []string, searchS string, N int) []string {

	newData := make([]string, 10)

	for index, value := range data {
		fmt.Print(searchS)
		fmt.Print(value)
		if value == searchS {
			fmt.Print("test")
			if index+N < len(data) {
				newData = append(newData, data[index])
			}
		}
	}

	return newData
}

//Вывод N строк до совпадения
func modeB(data []string, searchS string, N int) []string {

	newData := make([]string, 10)

	return newData
}

//Вывод N строк после совпадения и до совпадения
func modeC(data []string, searchS string, N int) []string {

	newData := make([]string, 10)

	return newData
}
