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

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	var guard utils.Point
	dir := utils.UP
	for y, line := range lines {
		for x, ch := range line {
			if ch == '^' {
				guard = utils.Point{X: x, Y: y}
			}
		}
	}
	blocked := utils.GetHashGrid(strings.Replace(s, "^", ".", -1), '.', '#')

	seen := map[utils.Point]bool{}
	for _, ok := blocked[guard]; ok; _, ok = blocked[guard] {
		seen[guard] = true
		guard.MoveInDir(dir, 1)
		if blocked[guard] {
			dir.Rotate180()
			guard.MoveInDir(dir, 1)
			dir.RotateLeft()
			guard.MoveInDir(dir, 1)
		}
	}

	return len(seen)
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
