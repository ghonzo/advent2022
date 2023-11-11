// Advent of Code 2022, Day 17
package main

import (
	"fmt"

	"github.com/ghonzo/advent2022/common"
)

// Day 17: Pyroclastic Flow
// Part 1 answer: 3111
// Part 2 answer: 1526744186042
func main() {
	fmt.Println("Advent of Code 2022, Day 17")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

type movesIter struct {
	str string
	i   int
}

type rock []common.Point

var rocks = []rock{
	{common.NewPoint(0, 0), common.NewPoint(1, 0), common.NewPoint(2, 0), common.NewPoint(3, 0)},
	{common.NewPoint(1, 0), common.NewPoint(0, 1), common.NewPoint(1, 1), common.NewPoint(2, 1), common.NewPoint(1, 2)},
	{common.NewPoint(0, 0), common.NewPoint(1, 0), common.NewPoint(2, 0), common.NewPoint(2, 1), common.NewPoint(2, 2)},
	{common.NewPoint(0, 0), common.NewPoint(0, 1), common.NewPoint(0, 2), common.NewPoint(0, 3)},
	{common.NewPoint(0, 0), common.NewPoint(1, 0), common.NewPoint(0, 1), common.NewPoint(1, 1)},
}

func (m *movesIter) next() byte {
	b := m.str[m.i]
	m.i++
	if m.i == len(m.str) {
		m.i = 0
	}
	return b
}

func part1(entries []string) int {
	moves := &movesIter{entries[0], 0}
	g := common.NewSparseGrid()
	var ymax int
	for n := 0; n < 2022; n++ {
		r := rocks[n%len(rocks)]
		pos := common.NewPoint(2, 4+ymax)
		for {
			var npos common.Point
			// Get moved
			if moves.next() == '>' {
				npos = pos.Add(common.R)
			} else {
				npos = pos.Add(common.L)
			}
			if valid(r, npos, g) {
				pos = npos
			} // else don't move
			// Now move down one
			npos = pos.Add(common.U)
			if !valid(r, npos, g) {
				ymax = max(ymax, addToGrid(r, pos, g))
				break
			}
			pos = npos
		}
		//fmt.Println(common.RenderGrid(g, '.'))
	}
	return ymax
}

func valid(r rock, pos common.Point, g common.Grid) bool {
	for _, p := range r {
		partPoint := pos.Add(p)
		if _, ok := g.CheckedGet(partPoint); ok || partPoint.Y() <= 0 || partPoint.X() < 0 || partPoint.X() > 6 {
			return false
		}
	}
	return true
}

// returns ymax
func addToGrid(r rock, pos common.Point, g common.Grid) int {
	mmy := new(common.MaxMin[int])
	for _, p := range r {
		rp := pos.Add(p)
		g.Set(rp, '#')
		mmy.Accept(rp.Y())
	}
	return mmy.Max
}

type state struct {
	rockNum int
	moveNum int
	depth   [7]int
}

type seen struct {
	move   int
	height int
}

func part2(entries []string) int {
	moves := &movesIter{entries[0], 0}
	g := common.NewSparseGrid()
	var ymax int
	stateMap := make(map[state]seen)
	var addedHeight int
	rocksLimit := 1000000000000
	for n := 0; n < rocksLimit; n++ {
		rockNum := n % len(rocks)
		r := rocks[rockNum]
		pos := common.NewPoint(2, 4+ymax)
		for {
			var npos common.Point
			// Get moved
			if moves.next() == '>' {
				npos = pos.Add(common.R)
			} else {
				npos = pos.Add(common.L)
			}
			if valid(r, npos, g) {
				pos = npos
			} // else don't move
			// Now move down one
			npos = pos.Add(common.U)
			if !valid(r, npos, g) {
				ymax = max(ymax, addToGrid(r, pos, g))
				break
			}
			pos = npos
		}
		// Have we ever seen this exact state before?
		s := state{rockNum, moves.i, calcDepth(g, ymax)}
		if prev, ok := stateMap[s]; ok {
			// how many cycles can we skip?
			deltaHeight := ymax - prev.height
			cycleLength := n - prev.move
			cyclesLeft := (rocksLimit - n) / cycleLength
			addedHeight = deltaHeight * cyclesLeft
			n += cycleLength * cyclesLeft
			// Clear out the state map, we don't need it anymore
			stateMap = make(map[state]seen)
		} else {
			stateMap[s] = seen{move: n, height: ymax}
		}
	}
	return ymax + addedHeight
}

func calcDepth(g common.Grid, top int) [7]int {
	// For each of the x-coordinate, how far down until we hit a rock?
	depths := [7]int{}
	for x := 0; x < 7; x++ {
		for y := top; y >= 0; y-- {
			if _, ok := g.CheckedGet(common.NewPoint(x, y)); ok {
				depths[x] = top - y
				break
			}
		}
	}
	return depths
}
