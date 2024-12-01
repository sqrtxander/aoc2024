package utils

import "log"

type Queue[T any] []T

func (q Queue[T]) Push(v T) Queue[T] {
	return append(q, v)
}

func (q Queue[T]) Pop() (Queue[T], T) {
	if len(q) == 0 {
		log.Fatalln("Attempt to pop from empty queue")
	}
	return q[1:], q[0]
}

func (q Queue[T]) Peek() T {
	if len(q) == 0 {
		log.Fatalln("Attempt to peek at empty queue")
	}
	return q[0]
}

func (q Queue[T]) Clear() Queue[T] {
	for len(q) > 0 {
		q, _ = q.Pop()
	}
	return q
}
