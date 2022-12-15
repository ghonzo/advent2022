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

type sensor struct {
	p common.Point
	// closest beacon distance
	cbd int
}

// If the point cannot be a beacon because it is closer than the cbd, then return true
func (s *sensor) invalidPoint(p common.Point) bool {
	return s.p.Sub(p).ManhattanDistance() <= s.cbd
}

func part1(entries []string, row int) int {
	sensors, beaconSet := parseInput(entries)
	// We need to find the min and max x-coord and the max cdb so we know what x-range to search
	mmx := new(common.MaxMin[int])
	mmcdb := new(common.MaxMin[int])
	for _, s := range sensors {
		mmx.Accept(s.p.X())
		mmcdb.Accept(s.cbd)
	}
	// So how to we restrict our search? Using the min and max x-coords and the max beacon distance.
	var count int
	for x := mmx.Min - mmcdb.Max; x < mmx.Max+mmcdb.Max; x++ {
		p := common.NewPoint(x, row)
		// Skip if there's a beacon there already
		if beaconSet[p] {
			continue
		}
		for _, s := range sensors {
			if s.invalidPoint(p) {
				count++
				break
			}
		}
	}
	return count
}

var lineRegexp = regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)

// Parse the input and return all the sensor locations, as well as a set that represents the beacon locations
func parseInput(entries []string) (sensors []*sensor, beaconSet map[common.Point]bool) {
	sensors = make([]*sensor, len(entries))
	beaconSet = make(map[common.Point]bool)
	for i, line := range entries {
		group := lineRegexp.FindStringSubmatch(line)
		sp := common.NewPoint(common.Atoi(group[1]), common.Atoi(group[2]))
		bp := common.NewPoint(common.Atoi(group[3]), common.Atoi(group[4]))
		sensors[i] = &sensor{sp, sp.Sub(bp).ManhattanDistance()}
		beaconSet[bp] = true
	}
	return
}

func part2(entries []string, maxCoord int) int {
	sensors, beaconSet := parseInput(entries)
	// How do we restrict our search space? Since there's only going to be one point
	// in the entire space that is valid, it stands to reason that it will be one space
	// further away from some sensor's minimum beacon distance. So search the diamond
	// around each sensor for a point that is not "invalid"
	for _, s := range sensors {
		// search distance is cdb+1
		sd := s.cbd + 1
		for i := 0; i <= sd; i++ {
		Offset:
			for _, offset := range [4]common.Point{
				common.NewPoint(i, -(sd - i)), // NE edge
				common.NewPoint(sd-i, -i),     // SE edge
				common.NewPoint(-i, sd-i),     // SW edge
				common.NewPoint(-(sd - i), i), // WW edge
			} {
				p := s.p.Add(offset)
				if beaconSet[p] || p.X() < 0 || p.X() > maxCoord || p.Y() < 0 || p.Y() > maxCoord {
					continue
				}
				for _, s2 := range sensors {
					if s2.invalidPoint(p) {
						continue Offset
					}
				}
				// This must be it
				return p.X()*4000000 + p.Y()
			}
		}
	}
	panic("no dice")
}
