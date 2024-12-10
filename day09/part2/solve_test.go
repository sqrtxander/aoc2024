package main

import "testing"

var INPUT string = `
2333133121414131402
`[1:]

var EXPECTED int = 2858

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}
