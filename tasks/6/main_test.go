package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {

	//Arrange
	testTable := []struct {
		input struct {
			name string
			i    bool
			F    bool
		}
		expected []string
	}{
		{
			input: struct {
				name string
				i    bool
				F    bool
			}{name: "test.txt",
				i: false,
				F: false,
			},
			expected: []string{"aaayyy/lll/kkk", "bbb/ppp/rrr", "ccc", "ddd", "eee"},
		},
	}

	//Act
	for _, testCase := range testTable {
		result := readFile(testCase.input.name)

		t.Logf("Calling readFile(%v), result %s\n", testCase.input, result)

		//Assert
		assert.Equal(t, testCase.expected, result,
			fmt.Sprintf("Incorrext result. Expect %s, got %s", testCase.expected, result))
	}
}

func TestCutWithoutSeparated(t *testing.T) {
	testTable := []struct {
		input struct {
			data   []string
			fields int
			delimiter string
			separated bool
		}
		expected []string
	}{
		{
			input: struct {
				data   []string
				fields int
				delimiter string
				separated bool
			}{
				data:   []string{"aaayyy/lll/kkk", "bbb/ppp/rrr", "ccc", "ddd", "eee"},
				fields:     3,
				delimiter:  "/",
				separated:  false,
			},
			expected: []string{"kkk", "rrr", "ccc", "ddd", "eee"},
			//expected: []string{"kkk", "rrr"},
		},
	}

	//Act
	for _, testCase := range testTable {
		result := Cut(testCase.input.data, testCase.input.fields, testCase.input.delimiter, testCase.input.separated)

		//Assert
		assert.Equal(t, testCase.expected, result,
			fmt.Sprintf("Incorrext result. Expect %s, got %s", testCase.expected, result))
	}
}

func TestCutWithSeparated(t *testing.T) {
	testTable := []struct {
		input struct {
			data   []string
			fields int
			delimiter string
			separated bool
		}
		expected []string
	}{
		{
			input: struct {
				data   []string
				fields int
				delimiter string
				separated bool
			}{
				data:   []string{"aaayyy/lll/kkk", "bbb/ppp/rrr", "ccc", "ddd", "eee"},
				fields:     3,
				delimiter:  "/",
				separated:  true,
			},
			expected: []string{"kkk", "rrr"},
		},
	}

	//Act
	for _, testCase := range testTable {
		result := Cut(testCase.input.data, testCase.input.fields, testCase.input.delimiter, testCase.input.separated)

		//Assert
		assert.Equal(t, testCase.expected, result,
			fmt.Sprintf("Incorrext result. Expect %s, got %s", testCase.expected, result))
	}
}
