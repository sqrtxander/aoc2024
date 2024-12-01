package main

import (
	"aoc2024/support"
	"flag"
	"fmt"
	"log"
)

var partFlag = flag.Int("p", 0, "Part 1 or 2")

func main() {
	flag.Parse()
	part := *partFlag
	if part != 1 && part != 2 {
		fmt.Println(part)
		log.Fatalln("Must provide a part of either 1 or 2")
	}
	support.SubmitSolution(part)
}
