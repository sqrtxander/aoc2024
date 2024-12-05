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

func isValid(nums []int, rules map[int]map[int]bool) bool {
	seen := map[int]bool{}
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		for after := range seen {
			if !rules[num][after] {
				return false
			}
		}
		seen[num] = true
	}
	return true
}

func makeValid(nums []int, rules map[int]map[int]bool) []int {
	for target := 0; target < len(nums); target++ {
		targetNum := nums[target]
		for i := target; i >= 1; i-- {
			if utils.All(nums[:i], func(after int) bool {
				return !rules[targetNum][after]
			}) {
				break
			}
			nums[i], nums[i-1] = nums[i-1], nums[i]
		}
	}
	return nums
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	chunks := strings.Split(s, "\n\n")

	rules := map[int]map[int]bool{}
	for _, line := range strings.Split(chunks[0], "\n") {
		fstStr, sndStr, _ := strings.Cut(line, "|")
		fst := utils.HandledAtoi(fstStr)
		snd := utils.HandledAtoi(sndStr)
		if _, ok := rules[fst]; !ok {
			rules[fst] = map[int]bool{}
		}
		rules[fst][snd] = true
	}

	total := 0
	for _, line := range strings.Split(chunks[1], "\n") {
		nums := utils.Map(strings.Split(line, ","), utils.HandledAtoi)
		if !isValid(nums, rules) {
			nums = makeValid(nums, rules)
			total += nums[len(nums)/2]
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
