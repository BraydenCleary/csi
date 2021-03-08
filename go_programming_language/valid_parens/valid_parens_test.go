package main

import "testing"

type testCase struct {
	input  string
	output bool
}

var testCases = []testCase{
	{input: "()", output: true},
	{input: "()[]{}", output: true},
	{input: "(]", output: false},
	{input: "([)]", output: false},
	{input: "{[]}", output: true},
}

func TestIsValid(t *testing.T) {
	for _, testCase := range testCases {
		result := isValid(testCase.input)
		if result != testCase.output {
			t.Errorf("isValid(%s) called. result %t != expected %t", testCase.input, result, testCase.output)
		}
	}
}
