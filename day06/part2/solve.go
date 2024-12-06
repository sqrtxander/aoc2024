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

type state struct {
	p utils.Point
	d utils.Direction
}

func causesLoop(newPoint utils.Point, guard utils.Point, blocked utils.HashGrid) bool {
	seen := map[state]bool{}
	dir := utils.UP
	for _, ok := blocked[guard]; ok; _, ok = blocked[guard] {
		s := state{
			p: guard,
			d: dir,
		}
		if seen[s] {
			return true
		}
		seen[s] = true
		for !blocked[guard] && guard != newPoint {
			guard.MoveInDir(dir, 1)
			if _, ok := blocked[guard]; !ok {
				return false
			}
		}
		dir.Rotate180()
		guard.MoveInDir(dir, 1)
		dir.RotateLeft()
		guard.MoveInDir(dir, 1)
	}
	return false
}

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
	checkGuard := guard
	_ = dir
	blocked := utils.GetHashGrid(strings.Replace(s, "^", ".", -1), '.', '#')

	toCheck := map[utils.Point]bool{}
	for {
		toCheck[checkGuard] = true
		checkGuard.MoveInDir(dir, 1)
		if blocked[checkGuard] {
			dir.Rotate180()
			checkGuard.MoveInDir(dir, 1)
			dir.RotateLeft()
			checkGuard.MoveInDir(dir, 1)
		}
		if _, ok := blocked[checkGuard]; !ok {
			break
		}
	}

	count := 0
	for p := range toCheck {
		if causesLoop(p, guard, blocked) {
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
