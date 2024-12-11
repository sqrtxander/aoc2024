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

	grid := map[utils.Point]int{}
	for y, line := range lines {
		for x, ch := range line {
			grid[utils.Point{X: x, Y: y}] = utils.HandledAtoi(string(ch))
		}
	}

	total := 0
	for p, h := range grid {
		if h != 0 {
			continue
		}
		seen := map[utils.Point]bool{}
		queue := utils.Queue[utils.Point]{p}
		for len(queue) > 0 {
			var q utils.Point
			queue, q = queue.Pop()
			if grid[q] == 9 {
				total++
				continue
			}
			for _, r := range utils.Adjacent4(q) {
				if !seen[r] && grid[r] == grid[q]+1 {
					queue = queue.Push(r)
					seen[r] = true
				}
			}
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
