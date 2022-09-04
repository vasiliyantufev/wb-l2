package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	//Arrange тестовая выборка
	testTable := []struct {
		input    []string
		expected map[string][]string
	}{
		{
			input: []string{"пятка", "пятак", "тяпка", "листок", "слиток",
				"столик", "кот", "ток"},
			expected: map[string][]string{
				"кот":    {"кот", "ток"},
				"листок": {"листок", "слиток", "столик"},
				"пятак":  {"пятак", "пятка", "тяпка"},
			},
		},
	}

	//Act расчет результата
	for _, testCase := range testTable {
		result := FindAnagrams(testCase.input)

		t.Logf("Calling FindAnagrams(%s), result %s\n", testCase.input, result)

		//Assert сравнение результата с ожиданием
		assert.Equal(t, testCase.expected, result,
			fmt.Sprintf("Incorrext result. Expect %s, got %s", testCase.expected, result))
	}
}
