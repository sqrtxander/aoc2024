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

func isDescending(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= nums[i] {
			return false
		}
	}
	return true
}

func isAscending(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] <= nums[i] {
			return false
		}
	}
	return true
}

func correctDiff(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if diff := utils.Abs(nums[i-1] - nums[i]); diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")

	count := 0
	for _, line := range lines {
		nums := utils.Map(strings.Fields(line), utils.HandledAtoi)
		if (isAscending(nums) || isDescending(nums)) && correctDiff(nums) {
			count++
		}
	}

	return count
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
