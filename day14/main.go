// Advent of Code 2022, Day 14
package main

import (
	"fmt"
	"strings"

	"github.com/ghonzo/advent2022/common"
)

// Day 14: Regolith Reservoir
// Part 1 answer: 1061
// Part 2 answer: 25055
func main() {
	fmt.Println("Advent of Code 2022, Day 14")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	g := readGrid(entries)
	var units int
	for produceSand(g) {
		units++
	}
	return units
}

func part2(entries []string) int {
	g := readGrid(entries)
	var units int
	floorY := g.Size().Y() + 2
	for produceSand2(g, floorY) {
		units++
	}
	return units
}

func readGrid(entries []string) common.Grid {
	g := common.NewSparseGrid()
	for _, line := range entries {
		coords := strings.Split(line, " -> ")
		points := make([]common.Point, len(coords))
		for i, c := range coords {
			xy := strings.Split(c, ",")
			points[i] = common.NewPoint(common.Atoi(xy[0]), common.Atoi(xy[1]))
		}
		for i := 0; i < len(points)-1; i++ {
			addRocks(g, points[i], points[i+1])
		}
	}
	return g
}

func addRocks(g common.Grid, p1, p2 common.Point) {
	dx := common.Sgn(p2.X() - p1.X())
	dy := common.Sgn(p2.Y() - p1.Y())
	add := common.NewPoint(dx, dy)
	for p := p1; p != p2; p = p.Add(add) {
		g.Set(p, '#')
	}
	g.Set(p2, '#')
}

// Order of points to try
var dirs = []common.Point{common.D, common.DL, common.DR}

var seedPoint = common.NewPoint(500, 0)

func produceSand(g common.Grid) bool {
	maxY := g.Size().Y()
Down:
	for p := seedPoint; p.Y() < maxY; {
		for _, d := range dirs {
			if g.Get(p.Add(d)) == 0 {
				p = p.Add(d)
				continue Down
			}
		}
		// Nope come to rest here
		g.Set(p, 'o')
		return true
	}
	// Fell off the edge
	return false
}

func produceSand2(g common.Grid, floorY int) bool {
	if g.Get(seedPoint) != 0 {
		return false
	}
	p := seedPoint
Down:
	for p.Y() < floorY-1 {
		for _, d := range dirs {
			if g.Get(p.Add(d)) == 0 {
				p = p.Add(d)
				continue Down
			}
		}
		break
	}
	g.Set(p, 'o')
	return true
}
