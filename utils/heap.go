package utils

import (
	"log"
)

type Heap[T any] interface {
	Push(T) Heap[T]
	Pop() (Heap[T], T)
	Peek() T
	Heapify(int) Heap[T]
}
type HeapFunc[T any] struct {
	Heap []T
	Cmp  func(T, T) int
}

func (h HeapFunc[T]) Push(num T) HeapFunc[T] {
	currIdx := len(h.Heap)
	parentIdx := (currIdx - 1) / 2
	h.Heap = append(h.Heap, num)
	for currIdx > 0 && h.Cmp(h.Heap[currIdx], h.Heap[parentIdx]) > 0 {
		h.Heap[currIdx], h.Heap[parentIdx] = h.Heap[parentIdx], h.Heap[currIdx]
		currIdx = parentIdx
		parentIdx = (currIdx - 1) / 2
	}
	return h
}

func (h HeapFunc[T]) Pop() (HeapFunc[T], T) {
	l := len(h.Heap)
	if l == 0 {
		log.Fatalln("Attempt to pop from empty heap")
	}
	result := h.Heap[0]
	h.Heap[0] = h.Heap[l-1]
	h.Heap = h.Heap[:l-1]

	return h.Heapify(0), result
}

func (h HeapFunc[T]) Peek() T {
	if len(h.Heap) == 0 {
		log.Fatalln("Attempt to peek at empty heap")
	}
	return h.Heap[0]
}

func (h HeapFunc[T]) Heapify(idx int) HeapFunc[T] {
	n := len(h.Heap)
	if n == 0 {
		return h
	}
	lo := idx
	l := 2*idx + 1
	r := 2*idx + 2
	if l < n && h.Cmp(h.Heap[l], h.Heap[lo]) > 0 {
		lo = l
	}
	if r < n && h.Cmp(h.Heap[r], h.Heap[lo]) > 0 {
		lo = r
	}
	if lo == idx {
		return h
	}
	h.Heap[lo], h.Heap[idx] = h.Heap[idx], h.Heap[lo]
	return h.Heapify(lo)
}

func IntGreater(a, b int) int {
	return a - b
}

func MaxHeapInt(contents []int) HeapFunc[int] {
	return HeapFunc[int]{
		Heap: contents,
		Cmp:  IntGreater,
	}
}

func IntLower(a, b int) int {
	return b - a
}

func MinHeapInt(contents []int) HeapFunc[int] {
	return HeapFunc[int]{
		Heap: contents,
		Cmp:  IntLower,
	}
}
