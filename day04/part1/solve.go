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

func foundCount(grid map[utils.Point]byte, p utils.Point, word string) int {
	if grid[p] != word[0] {
		return 0
	}
	if len(word) == 1 {
		return 1
	}
	count := 0
	for _, dir := range utils.Adjacent8(utils.Point{X: 0, Y: 0}) {
		q := p
		found := true
		for _, ch := range word[1:] {
			q = utils.Add(q, dir)
			if grid[q] != byte(ch) {
				found = false
				break
			}
		}
		if found {
			count++
		}
	}
	return count
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")

	grid := map[utils.Point]byte{}

	for y, line := range lines {
		for x, ch := range line {
			grid[utils.Point{X: x, Y: y}] = byte(ch)
		}
	}
	count := 0
	for p := range grid {
		count += foundCount(grid, p, "XMAS")
	}

	return count
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
