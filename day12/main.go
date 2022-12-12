// Advent of Code 2022, Day 12
package main

import (
	"fmt"

	"github.com/ghonzo/advent2022/common"
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

type globalState struct {
	grid           common.Grid
	minStepByPoint map[common.Point]int
	minSolution    int
	// Is the given point the end of the road?
	completeFn  func(common.Point) bool
	validMoveFn func(fromElevation, toElevation byte) bool
}

func part1(entries []string) int {
	g, start, end := readGrid(entries)
	gs := &globalState{
		grid:           g,
		minStepByPoint: make(map[common.Point]int),
		minSolution:    999999,
		completeFn: func(pt common.Point) bool {
			return pt == end
		},
		validMoveFn: func(fromElevation, toElevation byte) bool {
			return toElevation <= fromElevation+1
		}}
	return findShortestPath(start, 0, gs)
}

func readGrid(entries []string) (common.Grid, common.Point, common.Point) {
	g := common.ArraysGridFromLines(entries)
	var start, end common.Point
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
	return g, start, end
}

func findShortestPath(pos common.Point, step int, gs *globalState) int {
	if gs.completeFn(pos) {
		return step
	}
	// Bad path
	if step >= gs.minSolution {
		return gs.minSolution
	}
	if min, ok := gs.minStepByPoint[pos]; !ok || step < min {
		gs.minStepByPoint[pos] = step
	} else {
		return gs.minSolution
	}
	mm := new(common.MaxMin[int])
	mm.Accept(gs.minSolution)
	currentElevation := gs.grid.Get(pos)
	for np := range pos.SurroundingCardinals() {
		if v, ok := gs.grid.CheckedGet(np); ok && gs.validMoveFn(currentElevation, v) {
			// We can explore this path
			mm.Accept(findShortestPath(np, step+1, gs))
		}
	}
	return mm.Min
}

func part2(entries []string) int {
	g, _, end := readGrid(entries)
	gs := &globalState{
		grid:           g,
		minStepByPoint: make(map[common.Point]int),
		minSolution:    999999,
		completeFn: func(pt common.Point) bool {
			return g.Get(pt) == 'a'
		},
		validMoveFn: func(fromElevation, toElevation byte) bool {
			return toElevation >= fromElevation-1
		}}
	return findShortestPath(end, 0, gs)
}
