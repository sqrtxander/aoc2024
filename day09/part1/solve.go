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
	full  bool
	id    int
	space int
}

func solve(s string) int {
	s = strings.TrimSpace(s)

	chsums := utils.Deque[file]{}
	for i, ch := range s {
		full := i%2 == 0
		chsums = chsums.PushRight(file{
			full:  full,
			id:    i / 2,
			space: utils.HandledAtoi(string(ch)),
		})
	}
	i := 1
	for i < len(chsums) && !chsums[i].full {
		for !chsums.PeekRight().full {
			chsums, _ = chsums.PopRight()
		}
		var from file
		chsums, from = chsums.PopRight()
		var initFromSpace int
		for {
			initFromSpace = from.space
			if i >= len(chsums) {
				chsums = chsums.PushRight(file{
					full:  true,
					id:    from.id,
					space: initFromSpace,
				})
				break
			}
			fitAmt := min(chsums[i].space, from.space)
			from.space -= fitAmt
			chsums[i].id = from.id               // to
			if chsums[i].space < initFromSpace { // if filled up to
				chsums[i].full = true
				i += 2
			} else if chsums[i].space == initFromSpace {
				chsums[i].full = true
				break
			} else {
				break
			}
		}
		if i >= len(chsums) {
			break
		}
		if !chsums[i].full {
			chsums[i].full = true
			initToSpace := chsums[i].space
			chsums[i].space = initFromSpace
			chsums = slices.Insert(chsums, i+1, file{
				full:  false,
				space: initToSpace - initFromSpace,
			})
			i++
		} else {
			i += 2
		}
	}

	i = 0
	total := 0
	for _, chsum := range chsums {
		oldI := i
		for i < chsum.space+oldI {
			total += chsum.id * i
			i++
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
