// Advent of Code 2022, Day 16
package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ghonzo/advent2022/common"
	"github.com/oleiade/lane/v2"
)

// Day 16: Proboscidea Volcanium
// Part 1 answer: 1986
// Part 2 answer: 2464
func main() {
	fmt.Println("Advent of Code 2022, Day 16")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

type valve struct {
	name      string
	flowRate  int
	connected []*valve
}

type gamestate struct {
	cur *valve
	// Only used in part 2
	cur2   *valve
	minute int
	score  int
	openAt map[*valve]int
}

func (gs *gamestate) totalFlow() int {
	var sum int
	for k := range gs.openAt {
		sum += k.flowRate
	}
	return sum
}

// Signature is the time and the position(s)
func (gs *gamestate) signature() string {
	var sb strings.Builder
	sb.WriteByte(byte('A' + gs.minute))
	if gs.cur2 != nil {
		if gs.cur2.name < gs.cur.name {
			sb.WriteString(gs.cur2.name)
			sb.WriteString(gs.cur.name)
		} else {
			sb.WriteString(gs.cur.name)
			sb.WriteString(gs.cur2.name)
		}
	} else {
		sb.WriteString(gs.cur.name)
	}
	return sb.String()
}

func (gs *gamestate) open(v *valve) {
	gs.openAt[v] = gs.minute
}

func (gs *gamestate) copy() *gamestate {
	return &gamestate{gs.cur, gs.cur2, gs.minute, gs.score, gs.openAt}
}

// Also does a copy of the openAt map
func (gs *gamestate) fullCopy() *gamestate {
	ngs := gs.copy()
	ngs.openAt = make(map[*valve]int)
	for k, v := range gs.openAt {
		ngs.openAt[k] = v
	}
	return ngs
}

func part1(entries []string) int {
	valves := readInput(entries)
	var numNonZeroValves int
	for _, v := range valves {
		if v.flowRate > 0 {
			numNonZeroValves++
		}
	}
	// Map of state signature to max score
	maxState := make(map[string]int)
	pq := lane.NewMaxPriorityQueue[*gamestate, int]()
	pq.Push(&gamestate{valves["AA"], nil, 1, 0, make(map[*valve]int)}, 0)
	var maxScore int
	for !pq.Empty() {
		gs, _, _ := pq.Pop()
		if gs.minute == 30 {
			if gs.score > maxScore {
				maxScore = gs.score
			}
			continue
		}
		// Check to see if have already done better with this state
		sig := gs.signature()
		if ms, ok := maxState[sig]; ok && ms >= gs.score {
			// Yep, this is a dead path
			continue
		}
		maxState[sig] = gs.score
		// Find all the possible moves
		// Degenerate case: all valves are open, so we just wait around
		if len(gs.openAt) == numNonZeroValves {
			gs.score += gs.totalFlow() * (30 - gs.minute)
			gs.minute = 30
			pq.Push(gs, gs.score)
		}
		// Now let's find new paths
		// Do we have a valve we can open here?
		if gs.cur.flowRate > 0 && gs.openAt[gs.cur] == 0 {
			// Let's open it
			ngs := gs.fullCopy()
			ngs.open(ngs.cur)
			// And update the other state
			ngs.minute++
			ngs.score += ngs.totalFlow()
			pq.Push(ngs, ngs.score)
		}
		// Let's go to all the connected
		tf := gs.totalFlow()
		for _, v := range gs.cur.connected {
			ngs := gs.copy()
			ngs.cur = v
			ngs.minute++
			ngs.score += tf
			pq.Push(ngs, ngs.score)
		}
	}
	return maxScore
}

func part2(entries []string) int {
	valves := readInput(entries)
	var numNonZeroValves int
	for _, v := range valves {
		if v.flowRate > 0 {
			numNonZeroValves++
		}
	}
	// Map of state signature to max score
	maxState := make(map[string]int)
	pq := lane.NewMaxPriorityQueue[*gamestate, int]()
	pq.Push(&gamestate{valves["AA"], valves["AA"], 1, 0, make(map[*valve]int)}, 0)
	var maxScore int
	for !pq.Empty() {
		gs, _, _ := pq.Pop()
		if gs.minute == 26 {
			if gs.score > maxScore {
				maxScore = gs.score
			}
			continue
		}
		// Check to see if have already done better with this state
		sig := gs.signature()
		if ms, ok := maxState[sig]; ok && ms >= gs.score {
			// Yep, this is a dead path
			continue
		}
		maxState[sig] = gs.score
		// Find all the possible moves
		// Degenerate case: all valves are open, so we just wait around
		if len(gs.openAt) == numNonZeroValves {
			gs.score += gs.totalFlow() * (26 - gs.minute)
			gs.minute = 26
			pq.Push(gs, gs.score)
			continue
		}
		// Now let's find new paths
		// Do we have a valve we can open here?
		if gs.cur.flowRate > 0 && gs.openAt[gs.cur] == 0 {
			// Let's open it
			ngs := gs.fullCopy()
			ngs.open(ngs.cur)
			// What about the elephant? Can they open a valve too?
			if ngs.cur2.flowRate > 0 && ngs.openAt[ngs.cur2] == 0 {
				// Yep! Need to copy the copy
				nngs := ngs.fullCopy()
				nngs.open(nngs.cur2)
				// And update the other state
				nngs.minute++
				nngs.score += nngs.totalFlow()
				pq.Push(nngs, nngs.score)
			}
			// Or maybe the elephant travels somewhere else
			tf := ngs.totalFlow()
			for _, v := range ngs.cur2.connected {
				nngs := ngs.copy()
				nngs.cur2 = v
				nngs.minute++
				nngs.score += tf
				pq.Push(nngs, nngs.score)
			}
		}
		// What if we travel to a different room instead
		for _, v := range gs.cur.connected {
			// Can the elephant open a valve?
			if gs.cur2.flowRate > 0 && gs.openAt[gs.cur2] == 0 {
				// Yes, need a fully copy here
				ngs := gs.fullCopy()
				ngs.cur = v
				ngs.open(ngs.cur2)
				ngs.minute++
				ngs.score += ngs.totalFlow()
				pq.Push(ngs, ngs.score)
			}
			// And the elephant also moves
			tf := gs.totalFlow()
			for _, v2 := range gs.cur2.connected {
				ngs := gs.copy()
				ngs.cur = v
				ngs.cur2 = v2
				ngs.minute++
				ngs.score += tf
				pq.Push(ngs, ngs.score)
			}
		}
	}
	return maxScore
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
