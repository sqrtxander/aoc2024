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

type file struct {
	full    bool
	touched bool
	id      int
	space   int
}

func solve(s string) int {
	s = strings.TrimSpace(s)

	chsums := utils.Deque[file]{}
	for i, ch := range s {
		full := i%2 == 0
		chsums = chsums.PushRight(file{
			full:    full,
			touched: false,
			id:      i / 2,
			space:   utils.HandledAtoi(string(ch)),
		})
	}
	i := len(chsums)
	for i >= 0 {
		i--
		for i > 0 && (!chsums[i].full || chsums[i].touched) {
			i--
		}
		if i < 0 {
			break
		}
		from := chsums[i]
		chsums[i].touched = true
		var j int
		var to file
		found := false
		for j, to = range chsums {
			if j >= i {
				break
			}
			if !to.full && to.space >= from.space {
				chsums[i], chsums[j] = chsums[j], chsums[i]
				found = true
				break
			}
		}
		if found && chsums[i].space > chsums[j].space {
			tmp := chsums[i].space
			chsums[i].space = chsums[j].space
			chsums = slices.Insert(chsums, j+1, file{
				full:  false,
				space: tmp - chsums[j].space,
			})
		}
	}

	i = 0
	total := 0
	for _, chsum := range chsums {
		oldI := i
		if chsum.full {
			for i < chsum.space+oldI {
				total += chsum.id * i
				i++
			}
		} else {
			i += chsum.space
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
