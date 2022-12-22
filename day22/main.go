// Advent of Code 2022, Day 22
package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/ghonzo/advent2022/common"
)

// Day 22: Monkey Map
// Part 1 answer: 162186
// Part 2 answer: 55267
func main() {
	fmt.Println("Advent of Code 2022, Day 22")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

var pwRegexp = regexp.MustCompile(`\d+[RLX]`)

func part1(entries []string) int {
	g := common.ArraysGridFromLines(padToLongest(entries[:len(entries)-2]))
	pw := entries[len(entries)-1]
	if unicode.IsDigit(rune(pw[len(pw)-1])) {
		// Append an "X" which means don't change facing
		pw = pw + "X"
	}
	facing := common.R
	p := findFirst(g, common.NewPoint(0, 0), facing)
	for _, movement := range pwRegexp.FindAllString(pw, -1) {
		p = move(g, p, facing, common.Atoi(movement[:len(movement)-1]))
		switch movement[len(movement)-1] {
		case 'R':
			facing = facing.Right()
		case 'L':
			facing = facing.Left()
		case 'X':
			// last instruction might not have a direction
		}
	}
	score := 1000*(p.Y()+1) + 4*(p.X()+1)
	switch facing {
	case common.R:
		score += 0
	case common.D:
		score += 1
	case common.L:
		score += 2
	case common.U:
		score += 3
	}
	return score
}

func padToLongest(lines []string) []string {
	longest := 0
	for _, s := range lines {
		longest = common.Max(longest, len(s))
	}
	for i, s := range lines {
		ls := len(s)
		if ls < longest {
			lines[i] = lines[i] + strings.Repeat(" ", longest-ls)
		}
	}
	return lines
}

func findFirst(g common.Grid, startingPoint common.Point, dir common.Point) common.Point {
	for p := startingPoint; ; p = p.Add(dir) {
		if g.Get(p) != ' ' {
			return p
		}
	}
}

func move(g common.Grid, startingPoint common.Point, dir common.Point, dist int) common.Point {
	p := startingPoint
	for i := 0; i < dist; i++ {
		// try to move a step
		np := p.Add(dir)
		// let's see what's there
		v, ok := g.CheckedGet(np)
		// blank space or off the grid?
		if !ok || v == ' ' {
			// wrap around
			switch dir {
			case common.R:
				np = findFirst(g, common.NewPoint(0, np.Y()), dir)
			case common.L:
				np = findFirst(g, common.NewPoint(g.Size().X()-1, np.Y()), dir)
			case common.D:
				// Remember, down is *positive*
				np = findFirst(g, common.NewPoint(np.X(), 0), dir)
			case common.U:
				np = findFirst(g, common.NewPoint(np.X(), g.Size().Y()-1), dir)
			}
			v = g.Get(np)
		}
		// If it's a wall, we're done, return the old point
		if v == '#' {
			return p
		}
		// Successful move
		p = np
	}
	return p
}

func part2(entries []string) int {
	g := common.ArraysGridFromLines(padToLongest(entries[:len(entries)-2]))
	pw := entries[len(entries)-1]
	if unicode.IsDigit(rune(pw[len(pw)-1])) {
		// Append an "X" which means don't change facing
		pw = pw + "X"
	}
	facing := common.R
	p := findFirst(g, common.NewPoint(0, 0), facing)
	for _, movement := range pwRegexp.FindAllString(pw, -1) {
		p, facing = move3d(g, p, facing, common.Atoi(movement[:len(movement)-1]))
		switch movement[len(movement)-1] {
		case 'R':
			facing = facing.Right()
		case 'L':
			facing = facing.Left()
		case 'X':
			// last instruction might not have a direction
		}
	}
	score := 1000*(p.Y()+1) + 4*(p.X()+1)
	switch facing {
	case common.R:
		score += 0
	case common.D:
		score += 1
	case common.L:
		score += 2
	case common.U:
		score += 3
	}
	return score
}

func face(p common.Point) int {
	switch {
	case p.X() >= 50 && p.X() < 100 && p.Y() < 50:
		return 1
	case p.X() >= 100 && p.Y() < 50:
		return 2
	case p.X() >= 50 && p.X() < 100 && p.Y() >= 50 && p.Y() < 100:
		return 3
	case p.X() >= 50 && p.X() < 100 && p.Y() >= 100 && p.Y() < 150:
		return 4
	case p.X() < 50 && p.Y() >= 100 && p.Y() < 150:
		return 5
	case p.X() < 50 && p.Y() >= 150:
		return 6
	}
	panic("oops")
}

// New Point, New Facing
func move3d(g common.Grid, startingPoint common.Point, dir common.Point, dist int) (common.Point, common.Point) {
	p := startingPoint
	facing := dir
	for i := 0; i < dist; i++ {
		// try to move a step
		np := p.Add(facing)
		nf := facing
		// let's see what's there
		v, ok := g.CheckedGet(np)
		// blank space or off the grid?
		if !ok || v == ' ' {
			// wrap around
			switch facing {
			case common.R:
				switch face(p) {
				case 2:
					np = common.NewPoint(99, 149-p.Y())
					nf = common.L
				case 3:
					np = common.NewPoint(p.Y()-50+100, 49)
					nf = common.U
				case 4:
					np = common.NewPoint(149, 149-p.Y())
					nf = common.L
				case 6:
					np = common.NewPoint(p.Y()-150+50, 149)
					nf = common.U
				}
			case common.L:
				switch face(p) {
				case 1:
					np = common.NewPoint(0, 49-p.Y()+100)
					nf = common.R
				case 3:
					np = common.NewPoint(p.Y()-50, 100)
					nf = common.D
				case 5:
					np = common.NewPoint(50, 149-p.Y())
					nf = common.R
				case 6:
					np = common.NewPoint(p.Y()-150+50, 0)
					nf = common.D
				}
			case common.D:
				switch face(p) {
				case 2:
					np = common.NewPoint(99, p.X()-100+50)
					nf = common.L
				case 4:
					np = common.NewPoint(49, p.X()-50+150)
					nf = common.L
				case 6:
					np = common.NewPoint(p.X()+100, 0)
					nf = common.D
				}
			case common.U:
				switch face(p) {
				case 1:
					np = common.NewPoint(0, p.X()-50+150)
					nf = common.R
				case 2:
					np = common.NewPoint(p.X()-100, 199)
					nf = common.U
				case 5:
					np = common.NewPoint(50, p.X()+50)
					nf = common.R
				}
			}
			v = g.Get(np)
		}
		// If it's a wall, we're done, return the old point and facing
		if v == '#' {
			return p, facing
		}
		// Successful move
		p = np
		facing = nf
	}
	return p, facing
}
