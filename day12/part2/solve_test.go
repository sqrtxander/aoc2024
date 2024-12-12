package main

import "testing"

var INPUT string = `
AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA
`[1:]

var EXPECTED int = 368

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}
