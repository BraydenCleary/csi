package main

import "testing"

type testCase struct {
	matrix   [][]int
	target   int
	expected bool
}

var testCases = []testCase{
	// {
	// 	matrix: [][]int{
	// 		{1, 3, 5, 7},
	// 		{10, 11, 16, 20},
	// 		{23, 30, 34, 60},
	// 	},
	// 	target:   13,
	// 	expected: false,
	// },
	// {
	// 	matrix: [][]int{
	// 		{1, 3, 5, 7},
	// 		{10, 11, 16, 20},
	// 		{23, 30, 34, 60},
	// 	},
	// 	target:   3,
	// 	expected: true,
	// },
	// {
	// 	matrix: [][]int{
	// 		{1, 1},
	// 	},
	// 	target:   0,
	// 	expected: false,
	// },
	// {
	// 	matrix: [][]int{
	// 		{1, 3},
	// 	},
	// 	target:   3,
	// 	expected: true,
	// },
	// {
	// 	matrix: [][]int{
	// 		{1},
	// 		{3},
	// 	},
	// 	target:   0,
	// 	expected: false,
	// },
	// {
	// 	matrix: [][]int{
	// 		{1},
	// 		{3},
	// 	},
	// 	target:   2,
	// 	expected: false,
	// },
	// {
	// 	matrix: [][]int{
	// 		{1},
	// 		{3},
	// 	},
	// 	target:   3,
	// 	expected: true,
	// },
	{
		matrix: [][]int{
			{-8, -8, -7, -7, -6, -5, -3, -2},
			{0, 0, 1, 3, 4, 6, 8, 8},
			{11, 12, 14, 16, 18, 18, 19, 19},
			{22, 23, 25, 27, 28, 30, 30, 31},
			{34, 35, 37, 39, 40, 42, 43, 45},
			{48, 50, 51, 51, 53, 54, 55, 57},
			{58, 60, 62, 62, 62, 63, 63, 65},
			{68, 69, 71, 72, 72, 72, 74, 76},
		},
		target:   76,
		expected: true,
	},
}

func TestSearchMatrix(t *testing.T) {
	for _, tc := range testCases {
		result := searchMatrix(tc.matrix, tc.target)
		if result != tc.expected {
			t.Errorf("searchMatrix(%v, %d) called. Result %t != expected %t", tc.matrix, tc.target, result, tc.expected)
		}
	}
}
