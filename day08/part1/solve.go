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

func getAntinodePos(p1 utils.Point, p2 utils.Point) (utils.Point, utils.Point) {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	a1 := utils.Point{X: p1.X + dx, Y: p1.Y + dy}
	a2 := utils.Point{X: p2.X - dx, Y: p2.Y - dy}
	return a1, a2
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
				a1, a2 := getAntinodePos(p1, p2)
				if a1.X >= 0 && a1.X < width && a1.Y >= 0 && a1.Y < height {
					antinodes[a1] = true
				}
				if a2.X >= 0 && a2.X < width && a2.Y >= 0 && a2.Y < height {
					antinodes[a2] = true
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
