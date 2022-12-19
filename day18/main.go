// Advent of Code 2022, Day 18
package main

import (
	"fmt"
	"strings"

	"github.com/ghonzo/advent2022/common"
	"github.com/oleiade/lane/v2"
)

// Day 18: Boiling Boulders
// Part 1 answer: 3550
// Part 2 answer: 2028
func main() {
	fmt.Println("Advent of Code 2022, Day 18")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

type point3 struct {
	x, y, z int
}

func part1(entries []string) int {
	point := readPoints(entries)
	var sides int
	for _, p := range point {
		sides += 6
		for _, p2 := range point {
			if p == p2 {
				continue
			}
			if adjacent(p, p2) {
				sides--
			}
		}
	}
	return sides
}

func readPoints(entries []string) []point3 {
	p := make([]point3, len(entries))
	for i, s := range entries {
		pstr := strings.Split(s, ",")
		p[i] = point3{common.Atoi(pstr[0]), common.Atoi(pstr[1]), common.Atoi(pstr[2])}
	}
	return p
}

func adjacent(p1, p2 point3) bool {
	return common.Abs(p2.x-p1.x)+common.Abs(p2.y-p1.y)+common.Abs(p2.z-p1.z) == 1
}

type volume struct {
	x, y, z common.MaxMin[int]
}

func (v *volume) contains(p point3) bool {
	return p.x >= v.x.Min && p.x <= v.x.Max &&
		p.y >= v.y.Min && p.y <= v.y.Max &&
		p.z >= v.z.Min && p.z <= v.z.Max
}

func part2(entries []string) int {
	point := readPoints(entries)
	// Figure out the bounds of our space
	bounds := new(volume)
	space := make(map[point3]bool)
	for _, p := range point {
		space[p] = true
		bounds.x.Accept(p.x)
		bounds.y.Accept(p.y)
		bounds.z.Accept(p.z)
	}
	// Now we walk the entire space and fill any holes we find
	for x := bounds.x.Min + 1; x < bounds.x.Max; x++ {
		for y := bounds.y.Min + 1; y < bounds.y.Max; y++ {
			for z := bounds.z.Min + 1; z < bounds.z.Max; z++ {
				p := point3{x, y, z}
				if space[p] {
					continue
				}
				for p2 := range tryFill(p, space, bounds) {
					// Returns a set of spots to fill
					space[p2] = true
				}
			}
		}
	}
	// We can do this more efficiently later
	var sides int
	for p := range space {
		sides += 6
		for p2 := range space {
			if p == p2 {
				continue
			}
			if adjacent(p, p2) {
				sides--
			}
		}
	}
	return sides
}

func tryFill(start point3, space map[point3]bool, bounds *volume) map[point3]bool {
	// Start at the point
	addedSpace := map[point3]bool{start: true}
	queue := lane.NewQueue[point3]()
	queue.Enqueue(start)
	for queue.Size() > 0 {
		p, _ := queue.Dequeue()
		// Look for adjacent spaces
		for _, adj := range []point3{{p.x - 1, p.y, p.z}, {p.x + 1, p.y, p.z},
			{p.x, p.y - 1, p.z}, {p.x, p.y + 1, p.z},
			{p.x, p.y, p.z - 1}, {p.x, p.y, p.z + 1}} {
			if !space[adj] && !addedSpace[adj] {
				// open space ...  are we out of bounds though?
				if !bounds.contains(adj) {
					// We escaped so none of this matters
					return make(map[point3]bool)
				}
				addedSpace[adj] = true
				queue.Enqueue(adj)
			}
		}
	}
	// Must be a hole. Return all the spaces we found
	return addedSpace
}
