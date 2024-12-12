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

type pointnside struct {
	p    utils.Point
	side utils.Point
}

func sides(plot []utils.Point) int {
	result := 0
	seen := map[pointnside]bool{}
	dirs := []utils.Point{
		{X: 0, Y: -1},
		{X: 1, Y: 0},
		{X: 0, Y: 1},
		{X: -1, Y: 0},
	}
	for _, p := range plot {
		if !slices.Contains(plot, utils.Add(p, dirs[0])) && !seen[pointnside{p: p, side: dirs[0]}] { // up
			result++
			seen[pointnside{p: p, side: dirs[0]}] = true
			newP := utils.Point{X: p.X, Y: p.Y}
			for !slices.Contains(plot, utils.Add(newP, dirs[0])) && slices.Contains(plot, newP) {
				newP.MoveInDir(utils.LEFT, 1)
				seen[pointnside{p: newP, side: dirs[0]}] = true
			}
			newP = utils.Point{X: p.X, Y: p.Y}
			for !slices.Contains(plot, utils.Add(newP, dirs[0])) && slices.Contains(plot, newP) {
				newP.MoveInDir(utils.RIGHT, 1)
				seen[pointnside{p: newP, side: dirs[0]}] = true
			}
		}
		if !slices.Contains(plot, utils.Add(p, dirs[1])) && !seen[pointnside{p: p, side: dirs[1]}] { // right
			result++
			seen[pointnside{p: p, side: dirs[1]}] = true
			newP := utils.Point{X: p.X, Y: p.Y}
			for !slices.Contains(plot, utils.Add(newP, dirs[1])) && slices.Contains(plot, newP) {
				newP.MoveInDir(utils.UP, 1)
				seen[pointnside{p: newP, side: dirs[1]}] = true
			}
			newP = utils.Point{X: p.X, Y: p.Y}
			for !slices.Contains(plot, utils.Add(newP, dirs[1])) && slices.Contains(plot, newP) {
				newP.MoveInDir(utils.DOWN, 1)
				seen[pointnside{p: newP, side: dirs[1]}] = true
			}
		}
		if !slices.Contains(plot, utils.Add(p, dirs[2])) && !seen[pointnside{p: p, side: dirs[2]}] { // down
			result++
			seen[pointnside{p: p, side: dirs[2]}] = true
			newP := utils.Point{X: p.X, Y: p.Y}
			for !slices.Contains(plot, utils.Add(newP, dirs[2])) && slices.Contains(plot, newP) {
				newP.MoveInDir(utils.LEFT, 1)
				seen[pointnside{p: newP, side: dirs[2]}] = true
			}
			newP = utils.Point{X: p.X, Y: p.Y}
			for !slices.Contains(plot, utils.Add(newP, dirs[2])) && slices.Contains(plot, newP) {
				newP.MoveInDir(utils.RIGHT, 1)
				seen[pointnside{p: newP, side: dirs[2]}] = true
			}
		}
		if !slices.Contains(plot, utils.Add(p, dirs[3])) && !seen[pointnside{p: p, side: dirs[3]}] { // left
			result++
			seen[pointnside{p: p, side: dirs[3]}] = true
			newP := utils.Point{X: p.X, Y: p.Y}
			for !slices.Contains(plot, utils.Add(newP, dirs[3])) && slices.Contains(plot, newP) {
				newP.MoveInDir(utils.UP, 1)
				seen[pointnside{p: newP, side: dirs[3]}] = true
			}
			newP = utils.Point{X: p.X, Y: p.Y}
			for !slices.Contains(plot, utils.Add(newP, dirs[3])) && slices.Contains(plot, newP) {
				newP.MoveInDir(utils.DOWN, 1)
				seen[pointnside{p: newP, side: dirs[3]}] = true
			}
		}
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
		return area(plot) * sides(plot)
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
