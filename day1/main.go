// Advent of Code 2022, Day 1
package main

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/ghonzo/advent2022/common"
)

// Day 1: Calorie Counting
// Part 1 answer: 67633
// Part 2 answer: 199628
func main() {
	fmt.Println("Advent of Code 2022, Day 1")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	var max, accum int
	// add a synthetic blank line to the end
	for _, s := range append(entries, "") {
		if len(s) == 0 {
			if accum > max {
				max = accum
			}
			accum = 0
		} else {
			accum += common.Atoi(s)
		}
	}
	return max
}

func part2(entries []string) int {
	var elves []int
	var accum int
	// add a synthetic blank line to the end
	for _, s := range append(entries, "") {
		if len(s) == 0 {
			elves = append(elves, accum)
			accum = 0
		} else {
			accum += common.Atoi(s)
		}
	}
	//sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	slices.SortFunc(elves, func(a, b int) int {
		return -cmp.Compare(a, b)
	})
	return elves[0] + elves[1] + elves[2]
}
