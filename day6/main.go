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
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	str := entries[0]
	for i := 4; i < len(str); i++ {
		cMap := make(map[rune]bool)
		for _, r := range str[i-4 : i] {
			cMap[r] = true
		}
		if len(cMap) == 4 {
			return i
		}
	}
	return -1
}

func part2(entries []string) int {
	str := entries[0]
	for i := 14; i < len(str); i++ {
		cMap := make(map[rune]bool)
		for _, r := range str[i-14 : i] {
			cMap[r] = true
		}
		if len(cMap) == 14 {
			return i
		}
	}
	return -1
}
