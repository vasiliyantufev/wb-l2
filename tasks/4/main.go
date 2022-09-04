package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
Написать функцию поиска всех множеств анаграмм по словарю.


Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Требования:
Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке
в кодировке utf8
Выходные данные: ссылка на мапу множеств анаграмм
Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый
элемент которого,
слово из множества.
Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

*/

func main() {
	var values = []string{"пятка", "пятак", "тяпка", "листок", "слиток",
		"столик", "кот", "ток"}
	fmt.Println("Входные данные:", values)
	fmt.Println("Анаграмы:", FindAnagrams(values))
}

func FindAnagrams(slice []string) (anagrams map[string][]string) {

	anagrams = make(map[string][]string)
	m := make(map[string]map[string]struct{})
	//m := make(map[string][]string, 0)

	//ищем категории
	for _, w := range slice {

		lower := strings.ToLower(w)         //приводим к нижнему регистру
		symbols := strings.Split(lower, "") //дробим на символы
		sort.Strings(symbols)               //сортируем символы

		_, ok := m[strings.Join(symbols, "")]
		if ok != true { //если ключа ещё нет, то инициализировать мапу для второго ключа
			m[strings.Join(symbols, "")] = make(map[string]struct{})
		}

		//m = append(slice2, k)
		//m[strings.Join(symbols, "")] = make(map[string]struct{})
		m[strings.Join(symbols, "")][lower] = struct{}{}
		//m[strings.Join(symbols, "")] = make(map[string]struct{})
	}

	fmt.Println(m)

	var temp []string
	for _, keys := range m {

		//fmt.Println(key)
		//fmt.Println(word)
		//temp = make([]string, )
		//for k := range keys {
		//	temp = append(temp, k)
		//}
		temp = make([]string, 0, len(keys))
		for k := range keys {
			temp = append(temp, k)
		}
		sort.Strings(temp)
		for _, word := range temp {
			anagrams[temp[0]] = append(anagrams[temp[0]], word)
		}

	}

	return
}
