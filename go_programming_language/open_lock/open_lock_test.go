package main

import "testing"

type testCase struct {
	deadends []string
	target   string
	expected int
}

var testCases = []testCase{
	{
		deadends: []string{"0201", "0101", "0102", "1212", "2002"},
		target:   "0202",
		expected: 6,
	},
	{
		deadends: []string{"8888"},
		target:   "0009",
		expected: 1,
	},
	{
		deadends: []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"},
		target:   "8888",
		expected: -1,
	},
	{
		deadends: []string{"0000"},
		target:   "8888",
		expected: -1,
	},
	{
		deadends: []string{"1002", "1220", "0122", "0112", "0121"},
		target:   "1200",
		expected: 3,
	},
	{
		deadends: []string{"5557", "5553", "5575", "5535", "5755", "5355", "7555", "3555", "6655", "6455", "4655", "4455", "5665", "5445", "5645", "5465", "5566", "5544", "5564", "5546", "6565", "4545", "6545", "4565", "5656", "5454", "5654", "5456", "6556", "4554", "4556", "6554"},
		target:   "5555",
		expected: 10,
	},
}

func TestOpenLock(t *testing.T) {
	for _, tc := range testCases {
		result := openLock(tc.deadends, tc.target)
		if result != tc.expected {
			t.Errorf("openLock(%v, %s) called. Result %d != expected %d", tc.deadends, tc.target, result, tc.expected)
		}
	}
}
