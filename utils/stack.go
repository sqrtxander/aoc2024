package utils

import "log"

type Stack[T any] []T

func (s Stack[T]) Push(v T) Stack[T] {
	return append(s, v)
}

func (s Stack[T]) Pop() (Stack[T], T) {
	l := len(s)
	if l == 0 {
		log.Fatalln("Attempt to pop from empty stack")
	}
	return s[:l-1], s[l-1]
}

func (s Stack[T]) Peek() T {
	l := len(s)
	if l == 0 {
		log.Fatalln("Attempt to peek at empty stack")
	}
	return s[l-1]
}

func (s Stack[T]) Clear() Stack[T] {
	for len(s) > 0 {
		s, _ = s.Pop()
	}
	return s
}
