// Advent of Code 2022, Day 2
package main

import (
	"fmt"

	"github.com/ghonzo/advent2022/common"
)

// Day 2: Rock Paper Scissors
// Part 1 answer: 13009
// Part 2 answer: 10398
func main() {
	fmt.Println("Advent of Code 2022, Day 2")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	score := map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}
	var sum int
	for _, s := range entries {
		sum += score[s]
	}
	return sum
}

func part2(entries []string) int {
	score := map[string]int{
		"A X": 3,
		"A Y": 4,
		"A Z": 8,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 2,
		"C Y": 6,
		"C Z": 7,
	}
	var sum int
	for _, s := range entries {
		sum += score[s]
	}
	return sum
}
