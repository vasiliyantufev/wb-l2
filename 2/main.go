package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

/*
Создать Go-функцию, осуществляющую примитивную распаковку строки,
содержащую повторяющиеся символы/руны, например:
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

var ErrFirstDigit = errors.New("Первый символ не может быть число!")
var ErrLastBackSlash = errors.New("Бэкслеш не может быть последним символом!")
var ErrConverter = errors.New("Ошибка конвертации!")

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
	backslash := false

	for i, val := range inputStr {

		if i == 0 && unicode.IsDigit(val) {
			return "", ErrFirstDigit
		}

		if unicode.IsLetter(val) {
			outString = append(outString, val)
		}

		if val == '\\' {
			if backslash {
				outString = append(outString, inputStr[i-1])
				backslash = false
				continue
			} else if i+1 == len(inputStr) {
				return "", ErrLastBackSlash
			}
			backslash = true
		}

		if unicode.IsDigit(val) {

			st := string(val)
			count, err := strconv.Atoi(string(st))

			if err != nil {
				return "", ErrConverter
			}

			if backslash == true {
				outString = append(outString, val)
				backslash = false
			} else if count > 1 {
				for j := 0; j < count-1; j++ {
					outString = append(outString, inputStr[i-1])
				}
			}
		}

	}
	return string(outString), nil
}
