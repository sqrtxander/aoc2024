package main

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type equation struct {
	test      int
	remaining []int
}

func isValid(eq equation, idx int, result int) bool {
	if idx == len(eq.remaining) && result == eq.test {
		return true
	}
	if idx >= len(eq.remaining) || result > eq.test {
		return false
	}
	if isValid(eq, idx+1, result+eq.remaining[idx]) {
		return true
	}
	if isValid(eq, idx+1, result*eq.remaining[idx]) {
		return true
	}
	concatenated := result
	for digitCounter := eq.remaining[idx]; digitCounter > 0; digitCounter /= 10 {
		concatenated *= 10
	}
	concatenated += eq.remaining[idx]
	if isValid(eq, idx+1, concatenated) {
		return true
	}
	return false
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")

	equations := make([]equation, 0, len(lines))
	for _, line := range lines {
		test, rest, _ := strings.Cut(line, ": ")
		remaining := utils.Map(strings.Split(rest, " "), utils.HandledAtoi)
		equations = append(equations, equation{
			test:      utils.HandledAtoi(test),
			remaining: remaining,
		})
	}

	total := 0
	for _, eq := range equations {
		if isValid(eq, 1, eq.remaining[0]) {
			total += eq.test
		}
	}

	return total
}

func main() {
	var inputPath string
	if len(os.Args) > 1 {
		inputPath = os.Args[1]
	} else {
		_, currentFilePath, _, _ := runtime.Caller(0)
		dir := filepath.Dir(currentFilePath)
		dir = filepath.Dir(dir)
		inputPath = filepath.Join(dir, "input.in")
	}
	contents, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatalf("Error reading file %s:\n%v\n", inputPath, err)
		return
	}
	fmt.Println(solve(string(contents)))
}