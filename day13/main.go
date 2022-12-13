// Advent of Code 2022, Day 13
package main

import (
	"fmt"
	"sort"

	"github.com/ghonzo/advent2022/common"
	"github.com/oleiade/lane/v2"
)

// Day 13: Distress Signal
// Part 1 answer: 5808
// Part 2 answer: 22713
func main() {
	fmt.Println("Advent of Code 2022, Day 13")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	sum := 0
	for i := 1; i <= (len(entries)+1)/3; i++ {
		if inOrderStr(entries[(i-1)*3], entries[(i-1)*3+1]) {
			sum += i
		}
	}
	return sum
}

func part2(entries []string) int {
	var allPackets []*packetData
	for _, s := range entries {
		if len(s) > 0 {
			allPackets = append(allPackets, toPacketData(s))
		}
	}
	divPacket1 := toPacketData("[[2]]")
	divPacket2 := toPacketData("[[6]]")
	allPackets = append(allPackets, divPacket1, divPacket2)
	sort.Slice(allPackets, func(i, j int) bool {
		return inOrder(allPackets[i], allPackets[j]) < 0
	})
	key := 1
	for i, pd := range allPackets {
		if pd == divPacket1 || pd == divPacket2 {
			key *= (i + 1)
		}
	}
	return key
}

type packetData struct {
	isVal bool
	val   int
	list  []*packetData
}

func inOrderStr(left, right string) bool {
	lpd := toPacketData(left)
	rpd := toPacketData(right)
	return inOrder(lpd, rpd) < 0
}

func toPacketData(s string) *packetData {
	stack := lane.NewStack(new(packetData))
	var inNum bool
	for _, r := range s {
		if r == '[' {
			stack.Push(new(packetData))
			inNum = false
		} else if r >= '0' && r <= '9' {
			v := int(r - '0')
			h, _ := stack.Head()
			if inNum {
				// two digit number
				lastNumPd := h.list[len(h.list)-1]
				lastNumPd.val *= 10
				lastNumPd.val += v
			} else {
				numPd := &packetData{isVal: true, val: v}
				h.list = append(h.list, numPd)
				inNum = true
			}
		} else if r == ',' {
			inNum = false
		} else if r == ']' {
			pd, _ := stack.Pop()
			h, _ := stack.Head()
			h.list = append(h.list, pd)
			inNum = false
		}
	}
	top, _ := stack.Pop()
	return top.list[0]
}

// negative if in order, postive if not in order, 0 if equal
func inOrder(left, right *packetData) int {
	if left.isVal && right.isVal {
		return left.val - right.val
	}
	if left.isVal {
		left.list = []*packetData{{isVal: true, val: left.val}}
	}
	if right.isVal {
		right.list = []*packetData{{isVal: true, val: right.val}}
	}
	for i, lv := range left.list {
		if i >= len(right.list) {
			// Right ran out first
			return 1
		}
		if ord := inOrder(lv, right.list[i]); ord != 0 {
			return ord
		}
	}
	// Did left run out?
	if len(left.list) < len(right.list) {
		// Yep
		return -1
	}
	// Nope keep going
	return 0
}
