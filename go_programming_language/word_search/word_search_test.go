package main

import (
	"testing"
)

type testCase struct {
	board    [][]byte
	word     string
	expected bool
}

var testCases = []testCase{
	{
		board:    [][]byte{[]byte{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
		word:     "SEE",
		expected: true,
	},
	{
		board:    [][]byte{[]byte{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
		word:     "ABCB",
		expected: false,
	},
	{
		board:    [][]byte{[]byte{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
		word:     "ABCCED",
		expected: true,
	},
	{
		board:    [][]byte{[]byte{'a'}},
		word:     "a",
		expected: true,
	},
	{
		board:    [][]byte{[]byte{'b', 'a', 'a', 'b', 'a', 'b'}, {'a', 'b', 'a', 'a', 'a', 'a'}, {'a', 'b', 'a', 'a', 'a', 'b'}, {'a', 'b', 'a', 'b', 'b', 'a'}, {'a', 'a', 'b', 'b', 'a', 'b'}, {'a', 'a', 'b', 'b', 'b', 'a'}, {'a', 'a', 'b', 'a', 'a', 'b'}},
		word:     "aabbbbabbaababaaaabababbaaba",
		expected: true,
	},
}

func TestExist(t *testing.T) {
	for _, testCase := range testCases {
		output := exist(testCase.board, testCase.word)
		if output != testCase.expected {
			t.Errorf("exist(%v, %v) called. output %t != expected %t", testCase.board, testCase.word, output, testCase.expected)
		}
	}
}
