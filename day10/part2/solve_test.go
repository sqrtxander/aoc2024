package main

import "testing"

var INPUT string = `
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
`[1:]

var EXPECTED int = 81

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}
