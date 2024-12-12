package main

import (
	"aoc2024/utils"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"

	"github.com/gdamore/tcell/v2"
)

type updateState struct {
	x     int
	y     int
	new   bool
	style tcell.Style
	perim int
	area  int
}

func randStyle() tcell.Style {
	r := rand.Int31n(256)
	g := rand.Int31n(256)
	b := rand.Int31n(256)
	fg := tcell.NewHexColor(0xffffff)
	if (float32(r)*0.299 + float32(g)*0.587 + float32(b)*0.114) > 150 {
		fg = tcell.NewHexColor(0x000000)
	}
	return tcell.StyleDefault.Background(tcell.NewRGBColor(r, g, b)).Foreground(fg)
}

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

func solve(s string, resultChan *chan updateState) int {
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
		style := randStyle()
		for p, ch = range garden {
			nextPlot = append(nextPlot, p)
			delete(garden, p)
			*resultChan <- updateState{
				x:     p.X,
				y:     p.Y,
				new:   true,
				style: style,
				perim: perimiter(nextPlot),
				area:  area(nextPlot),
			}
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
					*resultChan <- updateState{
						x:     r.X,
						y:     r.Y,
						new:   false,
						style: style,
						perim: perimiter(nextPlot),
						area:  area(nextPlot),
					}
				}
			}
		}
		plots = append(plots, nextPlot)
	}

	return utils.Sum(utils.Map(plots, func(plot []utils.Point) int {
		return area(plot) * perimiter(plot)
	})...)
}

func getContents() string {
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
		return ""
	}
	return string(contents)
}
