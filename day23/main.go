// Advent of Code 2022, Day 23
package main

import (
	"fmt"

	"github.com/ghonzo/advent2022/common"
)

// Day 23: Unstable Diffusion
// Part 1 answer: 3970
// Part 2 answer: 923
func main() {
	fmt.Println("Advent of Code 2022, Day 23")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

var directions = [4][3]common.Point{
	{common.N, common.NE, common.NW},
	{common.S, common.SE, common.SW},
	{common.W, common.NW, common.SW},
	{common.E, common.NE, common.SE},
}

func part1(entries []string) int {
	ag := common.ArraysGridFromLines(entries)
	g := common.NewSparseGrid()
	for p := range ag.AllPoints() {
		if ag.Get(p) == '#' {
			g.Set(p, '#')
		}
	}
	for round := 0; round < 10; round++ {
		// New location, pointing to old location(s)
		proposedMoves := make(map[common.Point][]common.Point)
		// First half, consider positions
	ElfLoop:
		for elf := range g.AllPoints() {
			// If none around, do nothing
			if !hasSurrounding(g, elf) {
				proposedMoves[elf] = []common.Point{elf}
				continue
			}
			// Now consider the directions
		DirLoop:
			for i := 0; i < 4; i++ {
				for _, dir := range directions[(round+i)%4] {
					if g.Get(elf.Add(dir)) == '#' {
						// Nope
						continue DirLoop
					}
				}
				// Okay, empty, so propose this move
				newPt := elf.Add(directions[(round+i)%4][0])
				proposedMoves[newPt] = append(proposedMoves[newPt], elf)
				continue ElfLoop
			}
			// Stay where you are
			proposedMoves[elf] = []common.Point{elf}
		}
		// Second half ...
		g = common.NewSparseGrid()
		for k, v := range proposedMoves {
			if len(v) == 1 {
				g.Set(k, '#')
			} else {
				for _, p := range v {
					g.Set(p, '#')
				}
			}
		}
		//fmt.Println(common.RenderGrid(g, '.'))
	}
	// Find bounds
	mmX := new(common.MaxMin[int])
	mmY := new(common.MaxMin[int])
	for p := range g.AllPoints() {
		mmX.Accept(p.X())
		mmY.Accept(p.Y())
	}
	area := (mmX.Max - mmX.Min + 1) * (mmY.Max - mmY.Min + 1)
	return area - len(g)
}

func hasSurrounding(g common.Grid, p common.Point) bool {
	for surr := range p.SurroundingPoints() {
		if g.Get(surr) == '#' {
			// Yes an elf
			return true
		}
	}
	// Nope
	return false
}

func part2(entries []string) int {
	ag := common.ArraysGridFromLines(entries)
	g := common.NewSparseGrid()
	for p := range ag.AllPoints() {
		if ag.Get(p) == '#' {
			g.Set(p, '#')
		}
	}
	for round := 0; ; round++ {
		var stayed []common.Point
		// New location, pointing to old location(s)
		proposedMoves := make(map[common.Point][]common.Point)
		// First half, consider positions
	ElfLoop:
		for elf := range g.AllPoints() {
			// If none around, do nothing
			if !hasSurrounding(g, elf) {
				stayed = append(stayed, elf)
				continue
			}
			// Now consider the directions
		DirLoop:
			for i := 0; i < 4; i++ {
				for _, dir := range directions[(round+i)%4] {
					if g.Get(elf.Add(dir)) == '#' {
						// Nope
						continue DirLoop
					}
				}
				// Okay, empty, so propose this move
				newPt := elf.Add(directions[(round+i)%4][0])
				proposedMoves[newPt] = append(proposedMoves[newPt], elf)
				continue ElfLoop
			}
			// Stay where you are
			stayed = append(stayed, elf)
		}
		// Second half ...
		anyMoved := false
		g = common.NewSparseGrid()
		for k, v := range proposedMoves {
			if len(v) == 1 {
				g.Set(k, '#')
				anyMoved = true
			} else {
				for _, p := range v {
					g.Set(p, '#')
				}
			}
		}
		if !anyMoved {
			return round + 1
		}
		// Add in the ones that stayed
		for _, p := range stayed {
			g.Set(p, '#')
		}
		//fmt.Println(common.RenderGrid(g, '.'))
	}
}
