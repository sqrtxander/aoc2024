package utils

import (
	"log"
	"strconv"
)

type Pair[T, U any] struct {
	K T
	V U
}

func HandledAtoi(numStr string) int {
    num, err := strconv.Atoi(numStr)
    if err != nil {
        log.Fatalf("Invalid number: %q\n", numStr)
    }
    return num
}

func Sum(is ...int) int {
	result := 0
	for _, i := range is {
		result += i
	}
	return result
}

func MostFrequent[T comparable](slice []T) T {
	frequency := make(map[T]int)
	for _, elem := range slice {
		frequency[elem]++
	}
	count := 0
	var mostFrequent T
	for elem, freq := range frequency {
		if freq > count {
			mostFrequent = elem
			count = freq
		}
	}
	return mostFrequent
}

func LeastFrequent[T comparable](slice []T) T {
	frequency := make(map[T]int)
	for _, elem := range slice {
		frequency[elem]++
	}
	count := len(slice) + 1
	var mostFrequent T
	for elem, freq := range frequency {
		if freq < count {
			mostFrequent = elem
			count = freq
		}
	}
	return mostFrequent
}
