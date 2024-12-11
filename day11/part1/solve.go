package main

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func solve(s string) int {
	s = strings.TrimSpace(s)
	nums := strings.Fields(s)

	stones := map[int]int{}

	for _, num := range nums {
		stones[utils.HandledAtoi(num)]++
	}

	for range 25 {
		newStones := map[int]int{}
		for stone, count := range stones {
			if stone == 0 {
				newStones[1] += count
				continue
			}
			num := strconv.Itoa(stone)
			if len(num)%2 == 0 {
				s1 := num[len(num)/2:]
				s2 := num[:len(num)/2]
				newStones[utils.HandledAtoi(s1)] += count
				newStones[utils.HandledAtoi(s2)] += count
			} else {
				newStones[stone*2024] += count
			}
		}
		stones = newStones
	}

	total := 0
	for _, count := range stones {
		total += count
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
