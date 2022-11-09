package minheap

import (
	"golang.org/x/exp/constraints"
)

type minHeap[T constraints.Ordered] []T

func New[T constraints.Ordered]() *minHeap[T] {
	return &minHeap[T]{}
}

func (h *minHeap[T]) Push(value interface{}) {
	*h = append(*h, value.(T))
}

func (h *minHeap[T]) Pop() interface{} {
	length := len(*h)
	result := (*h)[length-1]
	*h = (*h)[:length-1]
	return result
}

func (h minHeap[T]) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h minHeap[T]) Len() int {
	return len(h)
}

func (h minHeap[T]) Less(i int, j int) bool {
	return h[i] < h[j]
}
