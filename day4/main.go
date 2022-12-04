// Advent of Code 2022, Day 4
package main

import (
	"fmt"
	"regexp"

	"github.com/ghonzo/advent2022/common"
)

// Day 4: Camp Cleanup
// Part 1 answer: 562
// Part 2 answer: 924
func main() {
	fmt.Println("Advent of Code 2022, Day 4")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

type assignment struct {
	min, max int
}

var lineRegex = regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)

func part1(entries []string) int {
	var count int
	for _, s := range entries {
		group := lineRegex.FindStringSubmatch(s)
		a1 := assignment{common.Atoi(group[1]), common.Atoi(group[2])}
		a2 := assignment{common.Atoi(group[3]), common.Atoi(group[4])}
		if (a1.min <= a2.min && a1.max >= a2.max) || (a2.min <= a1.min && a2.max >= a1.max) {
			count++
		}
	}
	return count
}

func part2(entries []string) int {
	var count int
	for _, s := range entries {
		group := lineRegex.FindStringSubmatch(s)
		a1 := assignment{common.Atoi(group[1]), common.Atoi(group[2])}
		a2 := assignment{common.Atoi(group[3]), common.Atoi(group[4])}
		if (a1.min >= a2.min && a1.min <= a2.max) || (a1.max >= a2.min && a1.max <= a2.max) ||
			(a2.min >= a1.min && a2.min <= a1.max) || (a2.max >= a1.min && a2.max <= a1.max) {
			count++
		}
	}
	return count
}
