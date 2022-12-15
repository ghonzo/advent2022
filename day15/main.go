// Advent of Code 2022, Day 15
package main

import (
	"fmt"
	"regexp"

	"github.com/ghonzo/advent2022/common"
)

// Day 15: Beacon Exclusion Zone
// Part 1 answer: 4886370
// Part 2 answer: 11374534948438
func main() {
	fmt.Println("Advent of Code 2022, Day 15")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries, 2000000))
	fmt.Printf("Part 2: %d\n", part2(entries, 4000000))
}

var lineRegexp = regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)

type sensor struct {
	p common.Point
	// closest beacon distance
	cbd int
}

func part1(entries []string, row int) int {
	sensors := make([]*sensor, len(entries))
	beaconSet := make(map[common.Point]bool)
	mmx := new(common.MaxMin[int])
	mmcdb := new(common.MaxMin[int])
	for i, line := range entries {
		group := lineRegexp.FindStringSubmatch(line)
		sp := common.NewPoint(common.Atoi(group[1]), common.Atoi(group[2]))
		bp := common.NewPoint(common.Atoi(group[3]), common.Atoi(group[4]))
		sensors[i] = &sensor{sp, sp.Sub(bp).ManhattanDistance()}
		mmx.Accept(sp.X())
		mmcdb.Accept(sensors[i].cbd)
		beaconSet[bp] = true
	}
	fmt.Println(mmx, mmcdb)
	var count int
	for x := mmx.Min - mmcdb.Max; x < mmx.Max+mmcdb.Max; x++ {
		p := common.NewPoint(x, row)
		if beaconSet[p] {
			continue
		}
		for _, s := range sensors {
			// Check to see if it's closer than the closest
			if s.p.Sub(p).ManhattanDistance() <= s.cbd {
				count++
				break
			}
		}
	}
	return count
}

func part2(entries []string, max int) int {
	sensors := make([]*sensor, len(entries))
	beaconSet := make(map[common.Point]bool)
	// As we are parsing, remember the sensors with the smallest cdb
	for i, line := range entries {
		group := lineRegexp.FindStringSubmatch(line)
		sp := common.NewPoint(common.Atoi(group[1]), common.Atoi(group[2]))
		bp := common.NewPoint(common.Atoi(group[3]), common.Atoi(group[4]))
		cdb := sp.Sub(bp).ManhattanDistance()
		sensors[i] = &sensor{sp, cdb}
		beaconSet[bp] = true
	}
	for _, s := range sensors {
		// search distance is cdb+1
		sd := s.cbd + 1

	Outer1:
		for i := 0; i <= sd; i++ {
			p := s.p.Add(common.NewPoint(i, -(sd - i)))
			if beaconSet[p] || p.X() < 0 || p.X() > max || p.Y() < 0 || p.Y() > max {
				continue
			}
			for _, s2 := range sensors {
				// Check to see if it's closer than the closets
				if s2.p.Sub(p).ManhattanDistance() <= s2.cbd {
					continue Outer1
				}
			}
			// must be it?
			return p.X()*4000000 + p.Y()
		}

	Outer2:
		for i := 0; i <= sd; i++ {
			p := s.p.Add(common.NewPoint(sd-i, -i))
			if beaconSet[p] || p.X() < 0 || p.X() > max || p.Y() < 0 || p.Y() > max {
				continue
			}
			for _, s2 := range sensors {
				// Check to see if it's closer than the closets
				if s2.p.Sub(p).ManhattanDistance() <= s2.cbd {
					continue Outer2
				}
			}
			// must be it?
			return p.X()*4000000 + p.Y()
		}

	Outer3:
		for i := 0; i <= sd; i++ {
			p := s.p.Add(common.NewPoint(-i, sd-i))
			if beaconSet[p] || p.X() < 0 || p.X() > max || p.Y() < 0 || p.Y() > max {
				continue
			}
			for _, s2 := range sensors {
				// Check to see if it's closer than the closets
				if s2.p.Sub(p).ManhattanDistance() <= s2.cbd {
					continue Outer3
				}
			}
			// must be it?
			return p.X()*4000000 + p.Y()
		}

	Outer4:
		for i := 0; i <= sd; i++ {
			p := s.p.Add(common.NewPoint(-(sd - i), i))
			if beaconSet[p] || p.X() < 0 || p.X() > max || p.Y() < 0 || p.Y() > max {
				continue
			}
			for _, s2 := range sensors {
				// Check to see if it's closer than the closets
				if s2.p.Sub(p).ManhattanDistance() <= s2.cbd {
					continue Outer4
				}
			}
			// must be it?
			return p.X()*4000000 + p.Y()
		}
	}
	return 0

}
