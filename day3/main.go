// Advent of Code 2022, Day 3
package main

import (
	"fmt"
	"strings"

	"github.com/ghonzo/advent2022/common"
)

// Day 3: Rucksack Reorganization
// Part 1 answer: 8401
// Part 2 answer: 2641
func main() {
	fmt.Println("Advent of Code 2022, Day 3")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	var sum int
	for _, s := range entries {
		for _, r := range s[:len(s)/2] {
			if strings.ContainsRune(s[len(s)/2:], r) {
				sum += val(r)
				break
			}
		}
	}
	return sum
}

func val(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r - 'a' + 1)
	}
	return int(r - 'A' + 27)
}

func part2(entries []string) int {
	var sum int
	for i := 0; i < len(entries); i += 3 {
		for _, r := range entries[i] {
			if strings.ContainsRune(entries[i+1], r) && strings.ContainsRune(entries[i+2], r) {
				sum += val(r)
				break
			}
		}
	}
	return sum
}
