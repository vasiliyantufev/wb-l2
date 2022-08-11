package main

import (
	"errors"
	"fmt"
	"unicode"
)

/*
Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""

Дополнительно
Реализовать поддержку escape-последовательностей.
Например:
qwe\4\5 => qwe45 (*)
qwe\45 => qwe44444 (*)
qwe\\5 => qwe\\\\\ (*)
*/

var ErrFirstDigit = errors.New("Пустая строка")

func main() {

	var str string

	fmt.Print("Ведите строку: ")
	fmt.Scan(&str)

	str, err := pakg(str)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}

func pakg(s string) (string, error) {

	//var runes []rune
	var outString []rune
	inputStr := []rune(s)

	for _, val := range inputStr {
		if unicode.IsDigit(val) {
			return "", ErrFirstDigit
		}
	}
	return string(outString), nil
}
