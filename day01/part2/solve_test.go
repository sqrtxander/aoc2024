package main

import "testing"

var INPUT string = `
3   4
4   3
2   5
1   3
3   9
3   3
`[1:]

var EXPECTED int = 31

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}
