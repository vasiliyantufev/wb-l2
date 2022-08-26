package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortNumber(t *testing.T) {
	//Arrange тестовая выборка
	testTable := []struct {
		input struct {
			data [][]string
			k    int
			r    bool
			n    bool
		}

		expected [][]string
	}{
		{
			input: struct {
				data [][]string
				k    int
				r    bool
				n    bool
			}{
				data: [][]string{{"111", "666", "777", "december"}, {"222", "555", "888", "june"}, {"333", "444", "999", "january"}},
				k:    2,
				r:    false,
				n:    true,
			},
			expected: [][]string{{"333", "444", "999", "january"}, {"222", "555", "888", "june"}, {"111", "666", "777", "december"}},
		},
		{
			input: struct {
				data [][]string
				k    int
				r    bool
				n    bool
			}{
				data: [][]string{{"111", "666", "777", "december"}, {"222", "555", "888", "june"}, {"333", "444", "999", "january"}},
				k:    2,
				r:    true,
				n:    true,
			},
			expected: [][]string{{"111", "666", "777", "december"}, {"222", "555", "888", "june"}, {"333", "444", "999", "january"}},
		},
	}

	//Act расчет результата
	for _, testCase := range testTable {
		result := sortNumber(testCase.input.data, testCase.input.k, testCase.input.r)

		//Assert сравнение результата с ожиданием
		assert.Equal(t, testCase.expected, result,
			fmt.Sprintf("Incorrext result. Expect %s, got %s", testCase.expected, result))
	}
}

func TestSortMonth(t *testing.T) {
	//Arrange тестовая выборка
	testTable := []struct {
		input struct {
			data [][]string
			k    int
			r    bool
			m    bool
		}

		expected [][]string
	}{
		{
			input: struct {
				data [][]string
				k    int
				r    bool
				m    bool
			}{
				data: [][]string{{"111", "666", "777", "december"}, {"222", "555", "888", "june"}, {"333", "444", "999", "january"}},
				k:    3,
				r:    false,
				m:    true,
			},
			expected: [][]string{{"333", "444", "999", "january"}, {"222", "555", "888", "june"}, {"111", "666", "777", "december"}},
		},
		{
			input: struct {
				data [][]string
				k    int
				r    bool
				m    bool
			}{
				data: [][]string{{"111", "666", "777", "december"}, {"333", "444", "999", "january"}, {"222", "555", "888", "june"}},
				k:    3,
				r:    true,
				m:    true,
			},
			expected: [][]string{{"111", "666", "777", "december"}, {"222", "555", "888", "june"}, {"333", "444", "999", "january"}},
		},
	}

	//Act расчет результата
	for _, testCase := range testTable {
		result := sortMonth(testCase.input.data, testCase.input.k, testCase.input.r)

		//Assert сравнение результата с ожиданием
		assert.Equal(t, testCase.expected, result,
			fmt.Sprintf("Incorrext result. Expect %s, got %s", testCase.expected, result))
	}
}
