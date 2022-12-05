// Advent of Code 2022, Day 5
package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gammazero/deque"
	"github.com/ghonzo/advent2022/common"
)

// Day 5: Supply Stacks
// Part 1 answer: LBLVVTVLP
// Part 2 answer: TPFFBDRJD
func main() {
	fmt.Println("Advent of Code 2022, Day 5")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %s\n", part1(entries))
	fmt.Printf("Part 2: %s\n", part2(entries))
}

var instructionRegex = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

func part1(entries []string) string {
	stacks := make([]*deque.Deque[byte], 9)
	stacks[0] = makeStack("WBDNCFJ")
	stacks[1] = makeStack("PZVQLST")
	stacks[2] = makeStack("PZBGJT")
	stacks[3] = makeStack("DTLJZBHC")
	stacks[4] = makeStack("GVBJS")
	stacks[5] = makeStack("PSQ")
	stacks[6] = makeStack("BVDFLMPN")
	stacks[7] = makeStack("PSMFBDLR")
	stacks[8] = makeStack("VDTR")
	for _, s := range entries {
		if len(s) == 0 || s[0] != 'm' {
			continue
		}
		group := instructionRegex.FindStringSubmatch(s)
		n := common.Atoi(group[1])
		src := common.Atoi(group[2]) - 1
		dest := common.Atoi(group[3]) - 1
		for i := 0; i < n; i++ {
			b := stacks[src].PopBack()
			stacks[dest].PushBack(b)
		}
	}
	var sb strings.Builder
	for _, stack := range stacks {
		sb.WriteByte(stack.Back())
	}
	return sb.String()
}

func makeStack(s string) *deque.Deque[byte] {
	stack := deque.New[byte]()
	for _, r := range s {
		stack.PushBack(byte(r))
	}
	return stack
}

func part2(entries []string) string {
	stacks := make([]*deque.Deque[byte], 9)
	stacks[0] = makeStack("WBDNCFJ")
	stacks[1] = makeStack("PZVQLST")
	stacks[2] = makeStack("PZBGJT")
	stacks[3] = makeStack("DTLJZBHC")
	stacks[4] = makeStack("GVBJS")
	stacks[5] = makeStack("PSQ")
	stacks[6] = makeStack("BVDFLMPN")
	stacks[7] = makeStack("PSMFBDLR")
	stacks[8] = makeStack("VDTR")
	for _, s := range entries {
		if len(s) == 0 || s[0] != 'm' {
			continue
		}
		group := instructionRegex.FindStringSubmatch(s)
		n := common.Atoi(group[1])
		src := common.Atoi(group[2]) - 1
		dest := common.Atoi(group[3]) - 1
		tempStack := deque.New[byte]()
		for i := 0; i < n; i++ {
			b := stacks[src].PopBack()
			tempStack.PushBack(b)
		}
		for i := 0; i < n; i++ {
			b := tempStack.PopBack()
			stacks[dest].PushBack(b)
		}
	}
	var sb strings.Builder
	for _, stack := range stacks {
		sb.WriteByte(stack.Back())
	}
	return sb.String()
}
