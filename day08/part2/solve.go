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

func getAntinodePos(p1 utils.Point, p2 utils.Point, w int, h int) []utils.Point {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	as := []utils.Point{}
	a := utils.Point{X: p1.X, Y: p1.Y}
	for a.X >= 0 && a.X < w && a.Y >= 0 && a.Y < h {
		as = append(as, a)
		a = utils.Point{X: a.X + dx, Y: a.Y + dy}
	}
	a = utils.Point{X: p2.X, Y: p2.Y}
	for a.X >= 0 && a.X < w && a.Y >= 0 && a.Y < h {
		as = append(as, a)
		a = utils.Point{X: a.X - dx, Y: a.Y - dy}
	}
	return as
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")

	width, height := len(lines[0]), len(lines)
	antennas := map[rune][]utils.Point{}
	for y, line := range lines {
		for x, ch := range line {
			if ch != '.' {
				antennas[ch] = append(antennas[ch], utils.Point{X: x, Y: y})
			}
		}
	}
	antinodes := map[utils.Point]bool{}
	_, _ = width, height
	for _, ps := range antennas {
		for i, p1 := range ps {
			for _, p2 := range ps[i+1:] {
				as := getAntinodePos(p1, p2, width, height)
				for _, a := range as {
					antinodes[a] = true
				}
			}
		}
	}

	return len(antinodes)
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
