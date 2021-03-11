package main

import "testing"

type testCase struct {
	grid     [][]byte
	expected int
}

var testCases = []testCase{
	// {
	// 	grid: [][]byte{
	// 		{'1', '1', '1', '1', '0'},
	// 		{'1', '1', '0', '1', '0'},
	// 		{'1', '1', '0', '0', '0'},
	// 		{'0', '0', '0', '0', '0'},
	// 	},
	// 	expected: 1,
	// },
	{
		grid: [][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		},
		expected: 3,
	},
}

func TestNumIslands(t *testing.T) {
	for _, tc := range testCases {
		result := numIslands(tc.grid)
		if result != tc.expected {
			t.Errorf("numIslands(%v) called. Result %d != expected %d", tc.grid, result, tc.expected)
		}
	}
}
