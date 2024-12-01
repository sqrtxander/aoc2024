package utils

import (
	"log"
	"strings"
)

type (
	HashGrid        map[Point]bool
	BoundedHashGrid struct {
		Grid HashGrid
		W    int
		H    int
	}
)

func GetHashGrid(s string, falsy rune, truthy rune) HashGrid {
	result := HashGrid{}
	lines := strings.Split(s, "\n")
	for y, line := range lines {
		for x, char := range line {
			switch char {
			case falsy:
				result[Point{X: x, Y: y}] = false
			case truthy:
				result[Point{X: x, Y: y}] = true
			default:
				log.Fatalf("Invalid hashdot character: '%c'\n", char)
			}
		}
	}
	return result
}

func ParseBoundedHashGrid(s string, falsy rune, truthy rune) (result BoundedHashGrid) {
	lines := strings.Split(s, "\n")
	result.Grid = GetHashGrid(s, falsy, truthy)
	result.H = len(lines)
	result.W = len(lines[0])
	return
}

func (g BoundedHashGrid) GetBoundedHash() string {
	result := ""
	for y := range g.H {
		for x := range g.W {
			if g.Grid[Point{X: x, Y: y}] {
				result += "#"
			} else {
				result += " "
			}
		}
		result += "\n"
	}
	return result[:len(result)-1]
}
