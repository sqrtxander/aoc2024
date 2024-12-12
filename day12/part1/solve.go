package main

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
)

func perimiter(plot []utils.Point) int {
	result := 0
	for _, p := range plot {
		contribution := 0
		for _, q := range utils.Adjacent4(p) {
			if !slices.Contains(plot, q) {
				contribution++
			}
		}
		result += contribution
	}
	return result
}

func area(plot []utils.Point) int {
	return len(plot)
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")

	garden := map[utils.Point]rune{}

	for y, line := range lines {
		for x, ch := range line {
			garden[utils.Point{X: x, Y: y}] = ch
		}
	}
	plots := [][]utils.Point{}
	for len(garden) > 0 {
		var p utils.Point
		var ch rune
		nextPlot := []utils.Point{}
		for p, ch = range garden {
			nextPlot = append(nextPlot, p)
			delete(garden, p)
			break
		}
		queue := utils.Queue[utils.Point]{p}
		for len(queue) > 0 {
			var q utils.Point
			queue, q = queue.Pop()
			for _, r := range utils.Adjacent4(q) {
				if c, ok := garden[r]; ok && c == ch {
					nextPlot = append(nextPlot, r)
					queue = append(queue, r)
					delete(garden, r)
				}
			}
		}
		plots = append(plots, nextPlot)
	}

	return utils.Sum(utils.Map(plots, func(plot []utils.Point) int {
		return area(plot) * perimiter(plot)
	})...)
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
