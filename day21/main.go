// Advent of Code 2022, Day 21
package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/ghonzo/advent2022/common"
)

// Day 21:
// Part 1 answer: 331319379445180
// Part 2 answer: 199628
func main() {
	fmt.Println("Advent of Code 2022, Day 21")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %s\n", part2(entries))
}

type monkey struct {
	name string
	// use in part 1
	value int
	// use in part 2
	valueExpr   string
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

func part2(entries []string) string {
	monkeyMap := make(map[string]*monkey)
	for _, m := range readMonkeys(entries) {
		if m.hasValue {
			m.valueExpr = strconv.Itoa(m.value)
		}
		monkeyMap[m.name] = m
	}
	monkeyMap["root"].op = '='
	monkeyMap["humn"].valueExpr = "x"
	return findValue2(monkeyMap, "root")
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

func findValue2(monkeyMap map[string]*monkey, name string) string {
	m := monkeyMap[name]
	if m.hasValue {
		return m.valueExpr
	}
	lExpr := findValue2(monkeyMap, m.left)
	lv, err1 := strconv.Atoi(lExpr)
	rExpr := findValue2(monkeyMap, m.right)
	rv, err2 := strconv.Atoi(rExpr)
	if err1 == nil && err2 == nil {
		switch m.op {
		case '+':
			m.valueExpr = strconv.Itoa(lv + rv)
		case '-':
			m.valueExpr = strconv.Itoa(lv - rv)
		case '*':
			m.valueExpr = strconv.Itoa(lv * rv)
		case '/':
			m.valueExpr = strconv.Itoa(lv / rv)
		}
	} else {
		var sb strings.Builder
		if err1 != nil {
			sb.WriteByte('(')
			sb.WriteString(lExpr)
			sb.WriteByte(')')
		} else {
			sb.WriteString(lExpr)
		}
		sb.WriteByte(m.op)
		if err2 != nil {
			sb.WriteByte('(')
			sb.WriteString(rExpr)
			sb.WriteByte(')')
		} else {
			sb.WriteString(rExpr)
		}
		m.valueExpr = sb.String()
	}
	m.hasValue = true
	return m.valueExpr
}
