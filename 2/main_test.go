package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPakg(t *testing.T) {
	//Arrange тестовая выборка
	testTable := []struct {
		input    string
		expected string
	}{
		{
			input:    "abcd",
			expected: "abcd",
		},
		{
			input:    `qwe\4\5`,
			expected: "qwe45",
		},
		{
			input:    `qwe\45`,
			expected: `qwe44444`,
		},
		{
			input:    `qwe\\5`,
			expected: `qwe\\\\\`,
		},
	}
	//Act расчет результата
	for _, testCase := range testTable {
		result := pakg(testCase.input)

		t.Logf("Calling Unpack(%v), result %s\n", testCase.input, result)

		//Assert сравнение результата с ожиданием
		assert.Equal(t, testCase.expected, result,
			fmt.Sprintf("Incorrext result. Expect %s, got %s", testCase.expected, result))
	}
}
