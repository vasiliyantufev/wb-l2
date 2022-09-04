package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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
	s bool
	f int
	d string
}

func main() {

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(string(data))

	flags := flag.Bool("s", false, "separated")
	flagf := flag.Int("f", column, "fields")
	flagd := flag.String("d", separate, "delimiter")

	flag.Parse()

	f := Flags{
		s: *flags,
		f: *flagf,
		d: *flagd,
	}

	//fmt.Print(f)

	//	arr := make([][]string, 0)
	str := strings.Split(string(data), " | ")
	//str := strings.Split("yyy bbb", " ")

	//fmt.Print(f.d)

	for _, line := range str {
		cols := strings.Split(line, f.d)

		if f.f > 0 && f.f <= len(cols) {
			fmt.Println(cols[f.f-1])
		} else if !f.s {
			fmt.Println(line)
		}
	}

	//for _, val := range arr {
	//	fmt.Printf("%s\n", val)
	//}

} /**/
