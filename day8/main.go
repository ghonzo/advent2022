// Advent of Code 2022, Day 8
package main

import (
	"fmt"

	"github.com/ghonzo/advent2022/common"
)

// Day 8: Treetop Tree House
// Part 1 answer: 1695
// Part 2 answer: 287040
func main() {
	fmt.Println("Advent of Code 2022, Day 8")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	g := common.ArraysGridFromLines(entries)
	var visible int
	for p := range g.AllPoints() {
		if isVisible(p, g) {
			visible++
		}
	}
	return visible
}

var cardinal = []common.Point{common.U, common.L, common.R, common.D}

func isVisible(p common.Point, g common.Grid) bool {
	val := g.Get(p)
	for _, d := range cardinal {
		np := p
		for {
			np = np.Add(d)
			v, ok := g.CheckedGet(np)
			if !ok {
				// made it to an edge, so visible
				return true
			}
			if v >= val {
				// as tall as the tree, so this direction is a no-go
				break
			}
		}
	}
	// Nope no directions
	return false
}

func part2(entries []string) int {
	g := common.ArraysGridFromLines(entries)
	mm := new(common.MaxMin[int])
	for p := range g.AllPoints() {
		mm.Accept(scenicScore(p, g))
	}
	return mm.Max
}

func scenicScore(p common.Point, g common.Grid) int {
	score := 1
	val := g.Get(p)
	for _, d := range cardinal {
		np := p
		dirScore := 0
		for {
			np = np.Add(d)
			v, ok := g.CheckedGet(np)
			if !ok {
				// made it to an edge
				if dirScore == 0 {
					return 0
				}
				break
			}
			dirScore++
			if v >= val {
				break
			}
		}
		score *= dirScore
	}
	return score
}
