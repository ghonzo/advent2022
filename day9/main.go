// Advent of Code 2022, Day 9
package main

import (
	"fmt"

	"github.com/ghonzo/advent2022/common"
)

// Day 9: Rope Bridge
// Part 1 answer: 6503
// Part 2 answer: 2724
func main() {
	fmt.Println("Advent of Code 2022, Day 9")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	var head, tail common.Point
	visited := make(map[common.Point]bool)
	for _, s := range entries {
		dir := resolveDir(s[0])
		n := common.Atoi(s[2:])
		for i := 0; i < n; i++ {
			head = head.Add(dir)
			tail = updateTail(head, tail)
			visited[tail] = true
		}
	}
	return len(visited)
}

func resolveDir(b byte) common.Point {
	switch b {
	case 'U':
		return common.U
	case 'R':
		return common.R
	case 'D':
		return common.D
	case 'L':
		return common.L
	}
	panic("Bad dir")
}

func updateTail(head, tail common.Point) common.Point {
	dx := head.X() - tail.X()
	dy := head.Y() - tail.Y()
	if common.Abs(dx) == 2 || common.Abs(dy) == 2 {
		return tail.Add(common.NewPoint(common.Sgn(dx), common.Sgn(dy)))
	}
	return tail
}

func part2(entries []string) int {
	var knot [10]common.Point
	visited := make(map[common.Point]bool)
	for _, s := range entries {
		dir := resolveDir(s[0])
		n := common.Atoi(s[2:])
		for i := 0; i < n; i++ {
			// move knot[0]
			knot[0] = knot[0].Add(dir)
			// Now move all the others
			for j := 1; j < 10; j++ {
				knot[j] = updateTail(knot[j-1], knot[j])
			}
			visited[knot[9]] = true
		}
	}
	return len(visited)
}
