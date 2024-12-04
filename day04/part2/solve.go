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

func foundXMas(grid map[utils.Point]byte, p utils.Point) bool {
	if grid[p] != 'A' {
		return false
	}
	mCount := 0
	sCount := 0
	for _, dir := range utils.Adjacent4Corners(utils.Point{X: 0, Y: 0}) {
		q := utils.Add(p, dir)
		if grid[q] == 'M' {
			mCount++
		}
		if grid[q] == 'S' {
			sCount++
		}
	}
	if mCount != 2 || sCount != 2 {
		return false
	}
	tl := utils.Add(p, utils.Point{X: -1, Y: -1})
	br := utils.Add(p, utils.Point{X: +1, Y: +1})
	return grid[tl] != grid[br]
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
		if foundXMas(grid, p) {
			count++
		}
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
