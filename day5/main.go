// Advent of Code 2022, Day 5
package main

import (
	"fmt"
	"regexp"
	"strings"

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

var leftSquareBracketRegex = regexp.MustCompile(`\[`)
var instructionRegex = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

func part1(entries []string) string {
	stack, blankLineNum := readStacks(entries)
	for _, str := range entries[blankLineNum+1:] {
		group := instructionRegex.FindStringSubmatch(str)
		n := common.Atoi(group[1])
		src := common.Atoi(group[2]) - 1
		dest := common.Atoi(group[3]) - 1
		crates := stack[src][:n]
		stack[src] = stack[src][n:]
		stack[dest] = common.Reverse(crates) + stack[dest]
	}
	var sb strings.Builder
	for _, s := range stack {
		sb.WriteByte(s[0])
	}
	return sb.String()
}

func readStacks(entries []string) ([]string, int) {
	stack := make([]string, 9)
	for lineNum, s := range entries {
		indexes := leftSquareBracketRegex.FindAllStringIndex(s, -1)
		if indexes == nil {
			s = strings.TrimSpace(s)
			numStacks := common.Atoi(s[len(s)-1:])
			return stack[:numStacks], lineNum + 1
		}
		for _, m := range indexes {
			// m is an []int with the first element being the index of the left square bracket
			pos := m[0]
			stack[pos/4] = stack[pos/4] + string(s[pos+1])
		}
	}
	panic("oops")
}

func part2(entries []string) string {
	stack, blankLineNum := readStacks(entries)
	for _, str := range entries[blankLineNum+1:] {
		group := instructionRegex.FindStringSubmatch(str)
		n := common.Atoi(group[1])
		src := common.Atoi(group[2]) - 1
		dest := common.Atoi(group[3]) - 1
		crates := stack[src][:n]
		stack[src] = stack[src][n:]
		stack[dest] = crates + stack[dest]
	}
	var sb strings.Builder
	for _, s := range stack {
		sb.WriteByte(s[0])
	}
	return sb.String()
}
