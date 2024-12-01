package utils

import "log"

type Deque[T any] []T

func (s Deque[T]) PushRight(v T) Deque[T] {
	return append(s, v)
}

func (s Deque[T]) PushLeft(v T) Deque[T] {
	return append([]T{v}, s...)
}

func (s Deque[T]) PopRight() (Deque[T], T) {
	l := len(s)
	if l == 0 {
		log.Fatalln("Attempt to pop from empty deque")
	}
	return s[:l-1], s[l-1]
}

func (s Deque[T]) PopLeft() (Deque[T], T) {
	l := len(s)
	if l == 0 {
		log.Fatalln("Attempt to pop from empty deque")
	}
	return s[1:], s[0]
}

func (s Deque[T]) PeekRight() T {
	l := len(s)
	if l == 0 {
		log.Fatalln("Attempt to peek at empty deque")
	}
	return s[l-1]
}

func (s Deque[T]) PeekLeft() T {
	l := len(s)
	if l == 0 {
		log.Fatalln("Attempt to peek at empty deque")
	}
	return s[0]
}

func (s Deque[T]) Clear() Deque[T] {
	for len(s) > 0 {
		s, _ = s.PopLeft()
	}
	return s
}
