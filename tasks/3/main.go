package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
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
	k          int
	n, r, u, m bool
}

var months = map[string]int{
	"january":   1,
	"february":  2,
	"mart":      3,
	"april":     4,
	"may":       5,
	"june":      6,
	"july":      7,
	"august":    8,
	"september": 9,
	"october":   10,
	"november":  11,
	"december":  12,
}

// Реализация структуры для взаимодействия с ней
func Init(k *int, n, r, u, m *bool) Flags {
	return Flags{
		k: *k,
		n: *n,
		r: *r,
		u: *u,
		m: *m,
	}
}

func main() {

	// Установка флагов
	k := flag.Int("k", 0, "Указание колонки для сортировки")
	n := flag.Bool("n", false, "Сортировать по числовому значению")
	r := flag.Bool("r", false, "Сортировать в обратном порядке")
	u := flag.Bool("u", false, "Не выводить повторяющиеся строки")

	m := flag.Bool("m", false, "сортировать по названию месяца")

	flag.Parse()

	flags := Init(k, n, r, u, m)

	//fmt.Print(flags)

	name := "test.txt"
	data := readFile(name)

	//fmt.Println(data)

	if flags.n == true {
		sortNumber(data, flags.k, flags.r)
	}
	if flags.m == true {
		sortMonth(data, flags.k, flags.r)
	}

	fmt.Print(data)
}

func readFile(name string) (data [][]string) {

	slice := make([]string, 0)

	file, err := os.Open(name)

	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	// read line by line
	for fileScanner.Scan() {
		//fmt.Println(fileScanner.Text())
		slice = append(slice, fileScanner.Text())
	}
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	//var matrix [][]string
	for _, str := range slice {
		data = append(data, strings.Fields(str))
	}

	fmt.Println(data)
	return
}

func sortNumber(data [][]string, k int, r bool) [][]string {
	sort.Slice(data, func(i, j int) bool {

		a := toInt(data[i][k-1])
		b := toInt(data[j][k-1])

		if r == true {
			return a > b
		}
		return a < b
	})

	return data
}

func sortMonth(data [][]string, k int, r bool) [][]string {

	sort.Slice(data, func(i, j int) bool {
		var a, b int
		a = months[data[i][k]]
		b = months[data[j][k]]

		if r == true {
			return a > b
		}
		return a < b
	})

	return data
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Неверный формат столбца: %v\n", err)
	}
	return i
}
