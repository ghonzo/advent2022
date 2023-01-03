// Advent of Code 2022, Day 19
package main

import (
	"fmt"
	"math"
	"regexp"

	"github.com/ghonzo/advent2022/common"
	"github.com/oleiade/lane/v2"
)

// Day 19: Not Enough Minerals
// Part 1 answer: 817
// Part 2 answer: 4216
func main() {
	fmt.Println("Advent of Code 2022, Day 19")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

// These are the types of elements, which will be used as indexes into resources, robots, and costs
const (
	ore = iota
	clay
	obsidian
	geode
)

type cost [4]int

type blueprint struct {
	robotCosts [4]cost
	maxRobots  [4]int
}

type state struct {
	time      int
	inventory [4]int
	robots    [4]int
}

func part1(entries []string) int {
	blueprints := readBlueprints(entries)
	var totalQualityLevel int
	for i, bp := range blueprints {
		m := calcMaxGeodes(bp, 24)
		totalQualityLevel += (i + 1) * m
	}
	return totalQualityLevel
}

func part2(entries []string) int {
	blueprints := readBlueprints(entries)
	var totalQualityLevel int = 1
	for i, bp := range blueprints {
		if i == 3 {
			break
		}
		m := calcMaxGeodes(bp, 32)
		totalQualityLevel *= m
	}
	return totalQualityLevel
}

var lineRegexp = regexp.MustCompile(`Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.`)

func readBlueprints(entries []string) []*blueprint {
	blueprints := make([]*blueprint, len(entries))
	for i, line := range entries {
		group := lineRegexp.FindStringSubmatch(line)
		bp := new(blueprint)
		bp.robotCosts[ore][ore] = common.Atoi(group[2])
		bp.robotCosts[clay][ore] = common.Atoi(group[3])
		bp.maxRobots[ore] = common.Max(bp.robotCosts[ore][ore], bp.robotCosts[clay][ore])
		bp.robotCosts[obsidian][ore] = common.Atoi(group[4])
		bp.maxRobots[ore] = common.Max(bp.maxRobots[ore], bp.robotCosts[obsidian][ore])
		bp.robotCosts[obsidian][clay] = common.Atoi(group[5])
		bp.maxRobots[clay] = bp.robotCosts[obsidian][clay]
		bp.robotCosts[geode][ore] = common.Atoi(group[6])
		bp.maxRobots[ore] = common.Max(bp.maxRobots[ore], bp.robotCosts[geode][ore])
		bp.robotCosts[geode][obsidian] = common.Atoi(group[7])
		bp.maxRobots[obsidian] = bp.robotCosts[geode][obsidian]
		bp.maxRobots[geode] = math.MaxInt
		blueprints[i] = bp
	}
	return blueprints
}

func calcMaxGeodes(bp *blueprint, timeLimit int) int {
	var maxGeodes int
	pq := lane.NewMaxPriorityQueue[state, int]()
	pq.Push(state{robots: [4]int{1, 0, 0, 0}}, 1)
	for !pq.Empty() {
		curState, _, _ := pq.Pop()
		// Try to build each kind of robot
		for robotType := ore; robotType <= geode; robotType++ {
			// Do we already have enough? If so, move along
			if curState.robots[robotType] >= bp.maxRobots[robotType] {
				continue
			}
			costToBuild := bp.robotCosts[robotType]
			timeToBuild := 0
			for resourceType := ore; resourceType <= geode; resourceType++ {
				if costToBuild[resourceType] > curState.inventory[resourceType] {
					// If we don't have any bots then it will never happen
					if curState.robots[resourceType] == 0 {
						timeToBuild = timeLimit + 1
						break
					}
					// Otherwise, this is how long it will take
					timeToBuild = common.Max(timeToBuild, (costToBuild[resourceType]-curState.inventory[resourceType]+curState.robots[resourceType]-1)/curState.robots[resourceType])
				}
			}
			nextState := curState
			nextState.time += timeToBuild + 1
			// If it'll take too long, then we can't build it
			if nextState.time >= timeLimit {
				continue
			}

			// Update the amount of resources
			for resourceType := ore; resourceType <= geode; resourceType++ {
				nextState.inventory[resourceType] += curState.robots[resourceType]*(timeToBuild+1) - costToBuild[resourceType]
			}
			// And update the selected bot
			nextState.robots[robotType]++
			// If we only built geode bots from here until the end of time and still can't catch up, skip
			remainingTime := timeLimit - nextState.time
			if ((remainingTime-1)*remainingTime)/2+nextState.inventory[geode]+remainingTime*nextState.robots[geode] <= maxGeodes {
				continue
			}
			pq.Push(nextState, nextState.inventory[geode]+1)
		}
		maxGeodes = common.Max(maxGeodes, curState.inventory[geode]+curState.robots[geode]*(timeLimit-curState.time))
	}
	return maxGeodes
}
