// Advent of Code 2022, Day 16
package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/ghonzo/advent2022/common"
	"github.com/oleiade/lane/v2"
)

// Day 16:
// Part 1 answer:
// Part 2 answer:
func main() {
	fmt.Println("Advent of Code 2022, Day 16")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	//fmt.Printf("Part 2: %d\n", part2(entries))
}

type valve struct {
	name      string
	flowRate  int
	connected []*valve
}

type gamestate struct {
	cur    *valve
	minute int
	openAt map[*valve]int
}

func (gs *gamestate) score() int {
	var score int
	for k, v := range gs.openAt {
		score += k.flowRate * (30 - v)
	}
	return score
}

func (gs *gamestate) isOpen() bool {
	return gs.openAt[gs.cur] != 0
}

func (gs *gamestate) signature() string {
	var sb strings.Builder
	sb.WriteString(gs.cur.name)
	var opened []string
	for k := range gs.openAt {
		opened = append(opened, k.name)
	}
	sort.Strings(opened)
	sb.WriteString(strings.Join(opened, ""))
	return sb.String()
}

func part1(entries []string) int {
	valves := readInput(entries)
	// Map to state signature to map of minute pointing to max score
	maxState := make(map[string]map[int]int)
	pq := lane.NewMaxPriorityQueue[*gamestate, int]()
	pq.Push(&gamestate{valves["AA"], 0, make(map[*valve]int)}, 0)
	var maxScore int
Outer:
	for !pq.Empty() {
		gs, _, _ := pq.Pop()
		if gs.minute == 30 {
			score := gs.score()
			if score > maxScore {
				maxScore = score
				fmt.Println(maxScore)
			}
			continue
		}
		// Check to see if we are moving backwards
		score := gs.score()
		sig := gs.signature()
		if scoreMap, ok := maxState[sig]; !ok {
			maxState[sig] = make(map[int]int)
			maxState[sig][gs.minute] = score
		} else {
			// see if there is a minute less than or equal that has a bigger score
			for k, v := range scoreMap {
				if k <= gs.minute && v >= score {
					// dead path
					continue Outer
				}
			}
			scoreMap[gs.minute] = score
		}
		// Find all the possible moves
		m := gs.minute + 1
		// Stay here
		pq.Push(gs.copyMoveTo(gs.cur, m), score)
		// Can we open the current valve?
		if !gs.isOpen() && gs.cur.flowRate > 0 {
			pq.Push(gs.copyOpenAt(m), score)
		}
		// Otherwise go to all the connected
		for _, v := range gs.cur.connected {
			pq.Push(gs.copyMoveTo(v, m), score)
		}
	}
	return maxScore
}

func (gs *gamestate) copyOpenAt(m int) *gamestate {
	gs2 := &gamestate{cur: gs.cur, minute: m, openAt: make(map[*valve]int)}
	for k, v := range gs.openAt {
		gs2.openAt[k] = v
	}
	gs2.openAt[gs.cur] = m
	return gs2
}

func (gs *gamestate) copyMoveTo(v *valve, m int) *gamestate {
	return &gamestate{cur: v, minute: m, openAt: gs.openAt}
}

var lineRegexp = regexp.MustCompile(`Valve (..) has flow rate=(\d+); tunnels? leads? to valves? (.+)`)

func readInput(entries []string) map[string]*valve {
	valves := make(map[string]*valve)
	for _, line := range entries {
		group := lineRegexp.FindStringSubmatch(line)
		name := group[1]
		valves[name] = &valve{name: name, flowRate: common.Atoi(group[2])}
	}
	for _, line := range entries {
		group := lineRegexp.FindStringSubmatch(line)
		v := valves[group[1]]
		for _, name := range strings.Split(group[3], ", ") {
			v.connected = append(v.connected, valves[name])
		}
	}
	return valves
}
