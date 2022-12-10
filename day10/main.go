// Advent of Code 2022, Day 10
package main

import (
	"fmt"
	"strings"

	"github.com/ghonzo/advent2022/common"
)

// Day 10: Cathode-Ray Tube
// Part 1 answer: 14220
// Part 2 answer: ZRARLFZU
func main() {
	fmt.Println("Advent of Code 2022, Day 10")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	part2(entries)
}

func part1(entries []string) int {
	x := 1
	cycle := 1
	sum := 0
	for _, s := range entries {
		if strings.HasPrefix(s, "addx") {
			cycle++
			if isInteresting(cycle) {
				sum += x * cycle
			}
			x += common.Atoi(s[5:])
		}
		cycle++
		if isInteresting(cycle) {
			sum += x * cycle
		}
	}
	return sum
}

func isInteresting(n int) bool {
	return n == 20 || n == 60 || n == 100 || n == 140 || n == 180 || n == 220
}

func part2(entries []string) {
	x := 1
	cycle := 1
	for _, s := range entries {
		printPixel(cycle, x)
		if strings.HasPrefix(s, "addx") {
			cycle++
			printPixel(cycle, x)
			x += common.Atoi(s[5:])
		}
		cycle++
	}
}

func printPixel(cycle, x int) {
	pos := (cycle - 1) % 40
	if pos == 0 {
		fmt.Println()
	}
	if pos >= x-1 && pos <= x+1 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
}
