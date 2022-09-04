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
			expected: []string{"aaa", "bbb", "ccc", "ddd", "eee"},
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

func TestModeA(t *testing.T) {
	testTable := []struct {
		input struct {
			data   []string
			search string
			N      int
		}
		expected []string
	}{
		{
			input: struct {
				data   []string
				search string
				N      int
			}{
				data:   []string{"alex", "who", "fill", "kolya", "stepan", "kolya", "ivan", "fill", "kolya", "stepan"},
				search: "alex",
				N:      2,
			},
			expected: []string{"who", "fill"},
		},
	}

	//Act
	for _, testCase := range testTable {
		result := modeA(testCase.input.data, testCase.input.search, testCase.input.N)

		t.Logf("Calling modeA(%v), result %s\n", testCase.input, result)

		//Assert
		assert.Equal(t, testCase.expected, result,
			fmt.Sprintf("Incorrext result. Expect %s, got %s", testCase.expected, result))
	}
}

func TestModeB(t *testing.T) {
	testTable := []struct {
		input struct {
			data   []string
			search string
			N      int
		}
		expected []string
	}{
		{
			input: struct {
				data   []string
				search string
				N      int
			}{
				data:   []string{"alex", "who", "fill", "kolya", "stepan", "kolya", "ivan", "fill", "kolya", "stepan"},
				search: "kolya",
				N:      2,
			},
			expected: []string{"who", "fill"},
		},
	}

	//Act
	for _, testCase := range testTable {
		result := modeB(testCase.input.data, testCase.input.search, testCase.input.N)

		t.Logf("Calling modeB(%v), result %s\n", testCase.input, result)

		//Assert
		assert.Equal(t, testCase.expected, result,
			fmt.Sprintf("Incorrext result. Expect %s, got %s", testCase.expected, result))
	}
}

func TestModeC(t *testing.T) {
	testTable := []struct {
		input struct {
			data   []string
			search string
			N      int
		}
		expected []string
	}{
		{
			input: struct {
				data   []string
				search string
				N      int
			}{
				data:   []string{"alex", "who", "fill", "kolya", "stepan", "ivan"},
				search: "kolya",
				N:      2,
			},
			expected: []string{"who", "fill", "stepan", "ivan"},
		},
	}

	//Act
	for _, testCase := range testTable {
		result := modeC(testCase.input.data, testCase.input.search, testCase.input.N)

		t.Logf("Calling modeC(%v), result %s\n", testCase.input, result)

		//Assert
		assert.Equal(t, testCase.expected, result,
			fmt.Sprintf("Incorrext result. Expect %s, got %s", testCase.expected, result))
	}
}
