// Advent of Code 2022, Day 24
package main

import (
	"fmt"

	"github.com/ghonzo/advent2022/common"
	"github.com/oleiade/lane/v2"
)

// Day 24: Blizzard Basin
// Part 1 answer: 326
// Part 2 answer: 976
func main() {
	fmt.Println("Advent of Code 2022, Day 24")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

type expState struct {
	pos    common.Point
	minute int
}

type blizzard struct {
	pos common.Point
	dir common.Point
}

func part1(entries []string) int {
	g := common.ArraysGridFromLines(entries)
	dim := g.Bounds()
	start := common.NewPoint(1, 0)
	end := common.NewPoint(dim.Width()-2, dim.Height()-1)
	blizzards := extractBlizzards(g)
	return findFastestPath(start, end, 0, blizzards, dim)
}

func score(start, end common.Point) int {
	return end.Sub(start).ManhattanDistance()
}

func extractBlizzards(g common.Grid) []blizzard {
	var blizzards []blizzard
	for p := range g.AllPoints() {
		v := g.Get(p)
		switch v {
		case '>':
			blizzards = append(blizzards, blizzard{p, common.E})
		case 'v':
			blizzards = append(blizzards, blizzard{p, common.S})
		case '<':
			blizzards = append(blizzards, blizzard{p, common.W})
		case '^':
			blizzards = append(blizzards, blizzard{p, common.N})
		}
	}
	return blizzards
}

func part2(entries []string) int {
	g := common.ArraysGridFromLines(entries)
	dim := g.Bounds()
	start := common.NewPoint(1, 0)
	end := common.NewPoint(dim.Width()-2, dim.Height()-1)
	blizzards := extractBlizzards(g)
	firstPartTime := findFastestPath(start, end, 0, blizzards, dim)
	secondPartTime := findFastestPath(end, start, firstPartTime, blizzards, dim)
	thirdPartTime := findFastestPath(start, end, secondPartTime, blizzards, dim)
	return thirdPartTime
}

func findFastestPath(start, end common.Point, startingMinute int, blizzards []blizzard, dim common.Rect) int {
	pq := lane.NewMinPriorityQueue[expState, int]()
	alreadyVisited := make(map[int]common.Grid)
	pq.Push(expState{start, startingMinute}, score(start, end))
	var minMinute = 9999999999
	for !pq.Empty() {
		state, prior, _ := pq.Pop()
		// Record if we have been in this state before and can prune
		if state.minute >= minMinute {
			continue
		}
		modMinute := state.minute % ((dim.Width() - 2) * (dim.Height() - 2))
		var visitedGrid common.Grid
		var ok bool
		if visitedGrid, ok = alreadyVisited[modMinute]; !ok {
			visitedGrid = common.NewSparseGrid()
			alreadyVisited[modMinute] = visitedGrid
		}
		if visitedGrid.Get(state.pos) != 0 {
			continue
		}
		visitedGrid.Set(state.pos, 1)
		minute := state.minute + 1
		blizzardGrid := findBlizzards(blizzards, minute, dim)
		for p := range state.pos.SurroundingCardinals() {
			if p == end {
				if minute < minMinute {
					minMinute = minute
				}
			}
			if p.X() >= 1 && p.X() < dim.Width()-1 && p.Y() >= 1 && p.Y() < dim.Height()-1 && blizzardGrid.Get(p) == 0 {
				// Possible move
				pq.Push(expState{p, minute}, score(p, end))
			}
		}
		if blizzardGrid.Get(state.pos) == 0 {
			// stay
			state.minute = minute
			pq.Push(state, prior)
		}
	}
	return minMinute
}

var blizzardsMemo = make(map[int]common.SparseGrid)

func findBlizzards(bList []blizzard, minute int, dim common.Rect) common.SparseGrid {
	var modMinute = minute % ((dim.Width() - 2) * (dim.Height() - 2))
	var g common.SparseGrid
	var ok bool
	if g, ok = blizzardsMemo[modMinute]; !ok {
		g = common.NewSparseGrid()
		for _, b := range bList {
			g.Set(computeBlizzardPos(b, modMinute, dim), 1)
		}
		blizzardsMemo[modMinute] = g
	}
	return g
}

func computeBlizzardPos(b blizzard, minute int, dim common.Rect) common.Point {
	xd := dim.Width() - 2
	yd := dim.Height() - 2
	// First the translate to get to a (0,0) origin
	p := b.pos.Add(common.NW).Add(b.dir.Times(minute))
	px := p.X() % xd
	if px < 0 {
		px += xd
	}
	py := p.Y() % yd
	if py < 0 {
		py += yd
	}
	// And translate back
	return common.NewPoint(px+1, py+1)
}
