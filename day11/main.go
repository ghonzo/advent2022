// Advent of Code 2022, Day 11
package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ghonzo/advent2022/common"
)

// Day 11:
// Part 1 answer:
// Part 2 answer:
func main() {
	fmt.Println("Advent of Code 2022, Day 11")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

type monkey struct {
	items            []int
	op               func(int) int
	mod              int
	tMonkey, fMonkey int
	inpects          int
}

func part1(entries []string) int {
	monkeys := readMonkeys(entries)
	monkeys[0].op = func(old int) int { return old * 11 }
	monkeys[1].op = func(old int) int { return old + 1 }
	monkeys[2].op = func(old int) int { return old + 6 }
	monkeys[3].op = func(old int) int { return old * old }
	monkeys[4].op = func(old int) int { return old * 7 }
	monkeys[5].op = func(old int) int { return old + 8 }
	monkeys[6].op = func(old int) int { return old + 5 }
	monkeys[7].op = func(old int) int { return old + 7 }
	for i := 0; i < 20; i++ {
		for _, m := range monkeys {
			for len(m.items) > 0 {
				m.inpects++
				worry := m.items[0]
				m.items = m.items[1:]
				worry = m.op(worry)
				worry /= 3
				if worry%m.mod == 0 {
					monkeys[m.tMonkey].items = append(monkeys[m.tMonkey].items, worry)
				} else {
					monkeys[m.fMonkey].items = append(monkeys[m.fMonkey].items, worry)
				}
			}
		}
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inpects > monkeys[j].inpects
	})
	return monkeys[0].inpects * monkeys[1].inpects
}

func readMonkeys(entries []string) []*monkey {
	var monkeys []*monkey
	for i := 0; i < len(entries); {
		m := monkey{}
		i++
		itemsStr := strings.Split(strings.TrimPrefix(entries[i], "  Starting items: "), ", ")
		for _, is := range itemsStr {
			m.items = append(m.items, common.Atoi(is))
		}
		i++
		// This is Operation
		i++
		m.mod = common.Atoi(strings.TrimPrefix(entries[i], "  Test: divisible by "))
		i++
		m.tMonkey = common.Atoi(strings.TrimPrefix(entries[i], "    If true: throw to monkey "))
		i++
		m.fMonkey = common.Atoi(strings.TrimPrefix(entries[i], "    If false: throw to monkey "))
		i++
		i++
		monkeys = append(monkeys, &m)
	}
	return monkeys
}

func part2(entries []string) int {
	monkeys := readMonkeys(entries)
	monkeys[0].op = func(old int) int { return old * 11 }
	monkeys[1].op = func(old int) int { return old + 1 }
	monkeys[2].op = func(old int) int { return old + 6 }
	monkeys[3].op = func(old int) int { return old * old }
	monkeys[4].op = func(old int) int { return old * 7 }
	monkeys[5].op = func(old int) int { return old + 8 }
	monkeys[6].op = func(old int) int { return old + 5 }
	monkeys[7].op = func(old int) int { return old + 7 }
	mod := 1
	for _, m := range monkeys {
		mod *= m.mod
	}
	for i := 0; i < 10000; i++ {
		for _, m := range monkeys {
			for len(m.items) > 0 {
				m.inpects++
				worry := m.items[0]
				m.items = m.items[1:]
				worry = m.op(worry)
				worry %= mod
				if worry%m.mod == 0 {
					monkeys[m.tMonkey].items = append(monkeys[m.tMonkey].items, worry)
				} else {
					monkeys[m.fMonkey].items = append(monkeys[m.fMonkey].items, worry)
				}
			}
		}
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inpects > monkeys[j].inpects
	})
	return monkeys[0].inpects * monkeys[1].inpects
}
