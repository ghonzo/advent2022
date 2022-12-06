// Advent of Code 2022, Day 6
package main

import (
	"fmt"

	"github.com/ghonzo/advent2022/common"
)

// Day 6: Tuning Trouble
// Part 1 answer: 1909
// Part 2 answer: 3380
func main() {
	fmt.Println("Advent of Code 2022, Day 6")
	str := common.ReadStringsFromFile("input.txt")[0]
	fmt.Printf("Part 1: %d\n", part1(str))
	fmt.Printf("Part 2: %d\n", part2(str))
}

func part1(str string) int {
	for i := 4; i < len(str); i++ {
		rMap := make(map[rune]bool)
		for _, r := range str[i-4 : i] {
			rMap[r] = true
		}
		if len(rMap) == 4 {
			return i
		}
	}
	return -1
}

func part2(str string) int {
	for i := 14; i < len(str); i++ {
		rMap := make(map[rune]bool)
		for _, r := range str[i-14 : i] {
			rMap[r] = true
		}
		if len(rMap) == 14 {
			return i
		}
	}
	return -1
}
