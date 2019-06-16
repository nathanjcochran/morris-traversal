package main

import (
	"reflect"
	"testing"
)

func TestTraverse(t *testing.T) {
	testCases := []struct {
		inserts  []int
		expected []int
	}{
		{
			inserts:  nil,
			expected: nil,
		},
		{
			inserts:  []int{1},
			expected: []int{1},
		},
		{
			inserts:  []int{1, 1},
			expected: []int{1, 1},
		},
		{
			inserts:  []int{1, 2},
			expected: []int{1, 2},
		},
		{
			inserts:  []int{2, 1},
			expected: []int{1, 2},
		},
		{
			inserts:  []int{1, 2, 3, 4, 5, 6, 7},
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			inserts:  []int{7, 6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			inserts:  []int{4, 2, 6, 1, 3, 5, 7},
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			inserts:  []int{4, 3, 2, 1, 5, 6, 7},
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
	}

	for _, testCase := range testCases {
		var tree Tree
		for _, i := range testCase.inserts {
			tree.Add(i)
		}

		actual := tree.Traverse()
		if !reflect.DeepEqual(testCase.expected, actual) {
			t.Errorf("Expected: %v\nActual: %v\n", testCase.expected, actual)
		}
	}
}
