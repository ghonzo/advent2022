// Advent of Code 2022, Day 25
package main

import (
	"fmt"
	"strings"

	"github.com/ghonzo/advent2022/common"
)

// Day 25: Full of Hot Air
// Part 1 answer: 2-0-020-1==1021=--01
// Part 2 answer: None!
func main() {
	fmt.Println("Advent of Code 2022, Day 25")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %s\n", part1(entries))
}

func part1(entries []string) string {
	var sum int
	for _, s := range entries {
		sum += fromSnafu(s)
	}
	return toSnafu(sum)
}

func fromSnafu(s string) int {
	var num int
	placeValue := 1
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case '1':
			num += placeValue
		case '2':
			num += 2 * placeValue
		case '-':
			num -= placeValue
		case '=':
			num -= 2 * placeValue
		}
		placeValue *= 5
	}
	return num
}

func toSnafu(n int) string {
	var sb strings.Builder
	for n > 0 {
		switch n % 5 {
		case 0:
			sb.WriteByte('0')
		case 1:
			sb.WriteByte('1')
		case 2:
			sb.WriteByte('2')
		case 3:
			sb.WriteByte('=')
			n += 5
		case 4:
			sb.WriteByte('-')
			n += 5
		}
		n /= 5
	}
	return common.Reverse(sb.String())
}
