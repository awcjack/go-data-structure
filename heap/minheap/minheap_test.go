package minheap_test

import (
	"container/heap"
	"testing"

	"github.com/awcjack/go-data-structure/heap/minheap"
)

func TestMinHeap_Integrated(t *testing.T) {
	type testcase struct {
		test             string
		push             []int
		expectedPopOrder []int
	}

	testcases := []testcase{
		{
			test:             "correct order insert",
			push:             []int{0, 1, 2, 3, 4, 5, 6, 7},
			expectedPopOrder: []int{0, 1, 2, 3, 4, 5, 6, 7},
		},
		{
			test:             "reverse order insert",
			push:             []int{7, 6, 5, 4, 3, 2, 1, 0},
			expectedPopOrder: []int{0, 1, 2, 3, 4, 5, 6, 7},
		},
		{
			test:             "random order insert",
			push:             []int{0, 7, 2, 4, 5, 3, 6, 1},
			expectedPopOrder: []int{0, 1, 2, 3, 4, 5, 6, 7},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			minHeap := minheap.New[int]()
			for _, v := range tc.push {
				heap.Push(minHeap, v)
			}
			for i, expectedValue := range tc.expectedPopOrder {
				value := heap.Pop(minHeap).(int)
				if expectedValue != value {
					t.Errorf("Expected get %v at index %v, but got %v", expectedValue, i, value)
				}
			}
		})
	}
}
