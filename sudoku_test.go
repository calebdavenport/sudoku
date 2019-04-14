package main

import (
	"reflect"
	"testing"
)

func TestSolveExclusivePair(t *testing.T) {
	expected := [][]int{{1, 2}, {1, 2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}
	testSource := [][]int{{1, 2}, {1, 2}, {3}, {4}, {5}, {6}, {2, 7}, {1, 8}, {1, 2, 9}}
	testPointer := []*[]int{}
	for i := 0; i < 9; i++ {
		testPointer = append(testPointer, &testSource[i])
	}

	exclusive_pair(testPointer)

	for i := 0; i < 9; i++ {
		if reflect.DeepEqual(testSource[i], expected[i]) {
			continue
		} else {
			t.Log("Test Source: ", testSource)
			t.Log("Expected: ", expected)
			t.Error("Test Failed")
		}
	}
}
