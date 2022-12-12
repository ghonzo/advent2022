// Advent of Code 2022, Day 12
package main

import (
	"fmt"

	"github.com/ghonzo/advent2022/common"
	"github.com/oleiade/lane/v2"
)

// Day 12: Hill Climbing Algorithm
// Part 1 answer: 484
// Part 2 answer: 478
func main() {
	fmt.Println("Advent of Code 2022, Day 12")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	g, start, end := readGrid(entries)
	return findShortestPath(start, g, func(pt common.Point) bool {
		return pt == end
	}, func(fromElevation, toElevation byte) bool {
		return toElevation <= fromElevation+1
	})
}

func part2(entries []string) int {
	g, _, end := readGrid(entries)
	return findShortestPath(end, g, func(pt common.Point) bool {
		return g.Get(pt) == 'a'
	}, func(fromElevation, toElevation byte) bool {
		return toElevation >= fromElevation-1
	})
}

func readGrid(entries []string) (g common.Grid, start common.Point, end common.Point) {
	g = common.ArraysGridFromLines(entries)
	for p := range g.AllPoints() {
		v := g.Get(p)
		if v == 'S' {
			start = p
			g.Set(p, 'a')
		} else if v == 'E' {
			end = p
			g.Set(p, 'z')
		}
	}
	return
}

func findShortestPath(start common.Point, grid common.Grid, completeFn func(common.Point) bool, validMoveFn func(fromElevation, toElevation byte) bool) int {
	// stores the minimum steps of all visited points
	minSteps := make(map[common.Point]int)
	// When you think shortest path finding, think Dijkstra's algorithm
	// (https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm). Use a generic priority queue
	// I found on the Internets. The priority is the number of steps.
	pq := lane.NewMinPriorityQueue[common.Point, int]()
	pq.Push(start, 0)
	for !pq.Empty() {
		pos, steps, _ := pq.Pop()
		// Did we complete?
		if completeFn(pos) {
			return steps
		}
		// Have we been here before?
		if min, ok := minSteps[pos]; ok && steps >= min {
			// Yep, so forget this path
			continue
		}
		// Remember we've been here
		minSteps[pos] = steps
		currentElevation := grid.Get(pos)
		// Test each direction
		for np := range pos.SurroundingCardinals() {
			if v, ok := grid.CheckedGet(np); ok && validMoveFn(currentElevation, v) {
				// Explore this path
				pq.Push(np, steps+1)
			}
		}
	}
	panic("Well crap")
}
