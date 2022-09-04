package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	search = strings.TrimSpace(string(data))

	//fmt.Println(len(search))
	//fmt.Println(len("l"))

	//search = strings.TrimSpace(in)

	//print(search)

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
		fmt.Print(modeA(fileData, search, f.A))
		//modeA(fileData, search, f.A)
	} else if f.B != 0 {
		//fmt.Println("b")
		fmt.Print(modeB(fileData, search, f.B))
		//modeB(fileData, search, f.B)
	} else if f.C != 0 {
		//fmt.Println("c")
		fmt.Println(modeC(fileData, search, f.C))
	}
}

//Чтение файла
func readFile(name string) (data []string) {

	file, err := os.Open(name)
	//file, err := os.ReadFile(name, os.O_RDWR, 0666)
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
		data = append(data, strings.TrimSpace(line))
	}
	return data
}

//Вывод N строк после совпадения
func modeA(data []string, searchS string, N int) []string {

	fmt.Println(data)
	var newData []string

	for index, value := range data {
		val := strings.TrimSpace(value)
		var count int
		if val == searchS {
			if index+N >= len(data) {
				count = len(data) - 1
			} else {
				count = N
			}
			for j := index + 1; j <= count; j++ {
				newData = append(newData, strings.TrimSpace(data[j]))
			}
			break
		}
	}
	return newData
}

//Вывод N строк до совпадения
func modeB(data []string, searchS string, N int) []string {

	//fmt.Println(data)
	var newData []string

	for index, value := range data {
		val := strings.TrimSpace(value)
		var count int
		if val == searchS {
			if index-N <= 0 {
				count = 0
			} else {
				count = index - N
			}
			for j := count; j < index; j++ {
				newData = append(newData, strings.TrimSpace(data[j]))
			}
			break
		}
	}
	return newData
}

//Вывод N строк после совпадения и до совпадения
func modeC(data []string, searchS string, N int) []string {

	//fmt.Println(data)
	var newData []string

	for index, value := range data {
		val := strings.TrimSpace(value)
		var count int
		var count2 int
		if val == searchS {

			if index-N <= 0 {
				count2 = 0
			} else {
				count2 = index - N
			}
			for j := count2; j < index; j++ {
				newData = append(newData, strings.TrimSpace(data[j]))
			}

			if index+N >= len(data) {
				count = len(data) - 1
			} else {
				count = index + N
			}
			for j := index + 1; j <= count; j++ {
				newData = append(newData, strings.TrimSpace(data[j]))
			}
			break
		}
	}
	return newData
}
