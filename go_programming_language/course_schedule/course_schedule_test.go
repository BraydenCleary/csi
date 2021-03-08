package main

import (
	"testing"
)

type testCase struct {
	numCourses    int
	prerequisites [][]int
	expected      bool
}

var testCases = []testCase{
	// {numCourses: 2, prerequisites: [][]int{[]int{1, 0}}, expected: true},
	{numCourses: 2, prerequisites: [][]int{[]int{1, 0}, []int{0, 1}}, expected: false},
}

func TestCourseSchedule(t *testing.T) {
	for _, testCase := range testCases {
		actualResult := canFinish(testCase.numCourses, testCase.prerequisites)
		if actualResult != testCase.expected {
			t.Errorf("canFinish called with %d %v. Result %t != %t", testCase.numCourses, testCase.prerequisites, actualResult, testCase.expected)
		}
	}
}
