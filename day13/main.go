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
		if inOrder(entries[(i-1)*3], entries[(i-1)*3+1]) {
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
		return compare(allPackets[i], allPackets[j]) < 0
	})
	key := 1
	for i, pd := range allPackets {
		if pd == divPacket1 || pd == divPacket2 {
			key *= (i + 1)
		}
	}
	return key
}

// Kind of a sloppy data structure. If isVal is true, then use val. Otherwise, use list.
type packetData struct {
	isVal bool
	val   int
	list  []*packetData
}

// Create a packetData that represents a value
func number(v int) *packetData {
	return &packetData{isVal: true, val: v}
}

func inOrder(left, right string) bool {
	return compare(toPacketData(left), toPacketData(right)) < 0
}

func toPacketData(s string) *packetData {
	stack := lane.NewStack(new(packetData))
	// Set to true if we're in the middle of parsing a number
	var inNum bool
	for _, r := range s {
		if r == '[' {
			stack.Push(new(packetData))
			inNum = false
		} else if r >= '0' && r <= '9' {
			v := int(r - '0')
			h, _ := stack.Head()
			if inNum {
				// mulit-digit number
				lastNumPd := h.list[len(h.list)-1]
				lastNumPd.val *= 10
				lastNumPd.val += v
			} else {
				h.list = append(h.list, number(v))
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
func compare(left, right *packetData) int {
	if left.isVal && right.isVal {
		return left.val - right.val
	}
	leftList := left.list
	if left.isVal {
		leftList = []*packetData{number(left.val)}
	}
	rightList := right.list
	if right.isVal {
		rightList = []*packetData{number(right.val)}
	}
	for i, lv := range leftList {
		if i >= len(rightList) {
			// Right ran out first
			return 1
		}
		if ord := compare(lv, rightList[i]); ord != 0 {
			return ord
		}
	}
	// Did left run out?
	if len(leftList) < len(rightList) {
		// Yep
		return -1
	}
	// Nope keep going
	return 0
}
