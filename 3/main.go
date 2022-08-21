package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

/*
Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные параметры):
на входе подается файл из несортированными строками, на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:

-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительно

Реализовать поддержку утилитой следующих ключей:

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учетом суффиксов
*/

// Cтруктура хранения флагов
type Flags struct {
	k       int
	n, r, u bool
}

// Реализация структуры для взаимодействия с ней
func Init(k *int, n, r, u *bool) Flags {
	return Flags{
		k: *k,
		n: *n,
		r: *r,
		u: *u,
	}
}

func main() {

	//var name string
	//var k int

	// Установка флагов
	k := flag.Int("k", -1, "Указание колонки для сортировки")
	n := flag.Bool("n", false, "Сортировать по числовому значению")
	r := flag.Bool("r", false, "Сортировать в обратном порядке")
	u := flag.Bool("u", false, "Не выводить повторяющиеся строки")
	flag.Parse()

	flags := Init(k, n, r, u)

	//fmt.Print(flags)

	name := "/home/walle/Projects/Go/wb-l2/3/test.txt"
	data := readFile(name, flags)

	fmt.Print(data)
}

func readFile(name string, flags Flags) (dataT [][]string) {

	file, err := os.Open(name)

	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	// read line by line
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
	//file.Close()

	//file, err := os.OpenFile(name, os.O_RDWR, 0666)
	//if err != nil {
	//	panic(err)
	//
	//}
	//
	return
}
