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

	leftIDs := make([]int, 0, len(lines))
	rightIDs := map[int]int{}
	for _, line := range lines {
		nums := strings.Fields(line)
		leftIDs = append(leftIDs, utils.HandledAtoi(nums[0]))
		rightIDs[utils.HandledAtoi(nums[1])] += 1
	}

	similarity := 0
	for _, id := range leftIDs {
		similarity += id * rightIDs[id]
	}
	return similarity
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
