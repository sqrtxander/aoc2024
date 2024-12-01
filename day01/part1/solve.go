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

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")

	leftIDs := make([]int, 0, len(lines))
	rightIDs := make([]int, 0, len(lines))
	for _, line := range lines {
		nums := strings.Fields(line)
		leftIDs = append(leftIDs, utils.HandledAtoi(nums[0]))
		rightIDs = append(rightIDs, utils.HandledAtoi(nums[1]))
	}
	slices.Sort(leftIDs)
	slices.Sort(rightIDs)
	distance := 0

	for i := range leftIDs {
		distance += utils.Abs(leftIDs[i] - rightIDs[i])
	}
	return distance
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
