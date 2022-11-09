package maxheap

import (
	"golang.org/x/exp/constraints"
)

type maxHeap[T constraints.Ordered] []T

func New[T constraints.Ordered]() *maxHeap[T] {
	return &maxHeap[T]{}
}

func (h *maxHeap[T]) Push(value interface{}) {
	*h = append(*h, value.(T))
}

func (h *maxHeap[T]) Pop() interface{} {
	length := len(*h)
	result := (*h)[length-1]
	*h = (*h)[:length-1]
	return result
}

func (h maxHeap[T]) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h maxHeap[T]) Len() int {
	return len(h)
}

func (h maxHeap[T]) Less(i int, j int) bool {
	return h[i] > h[j]
}
