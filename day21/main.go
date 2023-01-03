// Advent of Code 2022, Day 21
package main

import (
	"fmt"
	"unicode"

	"github.com/ghonzo/advent2022/common"
)

// Day 21: Monkey Math
// Part 1 answer: 331319379445180
// Part 2 answer: 3715799488132
func main() {
	fmt.Println("Advent of Code 2022, Day 21")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

type monkey struct {
	name        string
	value       int
	hasValue    bool
	left, right string
	op          byte
}

func part1(entries []string) int {
	monkeyMap := make(map[string]*monkey)
	for _, m := range readMonkeys(entries) {
		monkeyMap[m.name] = m
	}
	return findValue(monkeyMap, "root")
}

func part2(entries []string) int {
	monkeyMap := make(map[string]*monkey)
	for _, m := range readMonkeys(entries) {
		monkeyMap[m.name] = m
	}
	rewriteMonkeyMap(monkeyMap, "humn")
	return findValue(monkeyMap, "humn")
}

func readMonkeys(entries []string) []*monkey {
	monkeys := make([]*monkey, len(entries))
	for i, s := range entries {
		m := monkey{name: s[:4]}
		if unicode.IsDigit(rune(s[6])) {
			m.value = common.Atoi(s[6:])
			m.hasValue = true
		} else {
			m.left = s[6:10]
			m.op = s[11]
			m.right = s[13:]
		}
		monkeys[i] = &m
	}
	return monkeys
}

func findValue(monkeyMap map[string]*monkey, name string) int {
	m := monkeyMap[name]
	if m.hasValue {
		return m.value
	}
	lv := findValue(monkeyMap, m.left)
	rv := findValue(monkeyMap, m.right)
	switch m.op {
	case '+':
		m.value = lv + rv
	case '-':
		m.value = lv - rv
	case '*':
		m.value = lv * rv
	case '/':
		m.value = lv / rv
	}
	m.hasValue = true
	return m.value
}

func rewriteMonkeyMap(monkeyMap map[string]*monkey, name string) {
	// Look for any monkeys that have "name" as left or right
	m := &monkey{name: name}
	for k, v := range monkeyMap {
		if v.left == name {
			if k == "root" {
				m.value = findValue(monkeyMap, v.right)
				m.hasValue = true
				break
			}
			m.left = v.name
			m.right = v.right
			switch v.op {
			case '+':
				m.op = '-'
			case '-':
				m.op = '+'
			case '*':
				m.op = '/'
			case '/':
				m.op = '*'
			}
			rewriteMonkeyMap(monkeyMap, k)
			break
		}
		if v.right == name {
			if k == "root" {
				m.value = findValue(monkeyMap, v.left)
				m.hasValue = true
				break
			}
			switch v.op {
			case '+':
				m.left = v.name
				m.right = v.left
				m.op = '-'
			case '-':
				m.left = v.left
				m.right = v.name
				m.op = '-'
			case '*':
				m.left = v.name
				m.right = v.left
				m.op = '/'
			case '/':
				m.left = v.left
				m.right = v.name
				m.op = '/'
			}
			rewriteMonkeyMap(monkeyMap, k)
			break
		}
	}
	monkeyMap[name] = m
}
