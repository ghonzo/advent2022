// Advent of Code 2022, Day 20
package main

import (
	"fmt"

	"github.com/gammazero/deque"
	"github.com/ghonzo/advent2022/common"
)

// Day 20: Grove Positioning System
// Part 1 answer: 10707
// Part 2 answer: 2488332343098
func main() {
	fmt.Println("Advent of Code 2022, Day 20")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

func part1(entries []string) int {
	// Maintain two deques -- q has the input, while iq contains the indexes
	var q, iq deque.Deque[int]
	for i, s := range entries {
		q.PushBack(common.Atoi(s))
		iq.PushBack(i)
	}
	for pos := 0; pos < q.Len(); pos++ {
		index := iq.Index(func(a int) bool { return a == pos })
		q.Rotate(index)
		iq.Rotate(index)
		a := q.PopFront()
		b := iq.PopFront()
		q.Rotate(a)
		iq.Rotate(a)
		q.PushFront(a)
		iq.PushFront(b)
		// If we wanted to maintain the head, do this, but we don't care
		// q.Rotate(-a - index)
		// iq.Rotate(-a - index)
	}
	// Find the zero
	index := q.Index(func(a int) bool { return a == 0 })
	q.Rotate(index + 1000)
	sum := q.Front()
	q.Rotate(1000)
	sum += q.Front()
	q.Rotate(1000)
	sum += q.Front()
	return sum
}

func part2(entries []string) int {
	var q, iq deque.Deque[int]
	for i, s := range entries {
		q.PushBack(common.Atoi(s) * 811589153)
		iq.PushBack(i)
	}
	for mix := 0; mix < 10; mix++ {
		for pos := 0; pos < q.Len(); pos++ {
			index := iq.Index(func(a int) bool { return a == pos })
			q.Rotate(index)
			iq.Rotate(index)
			a := q.PopFront()
			b := iq.PopFront()
			q.Rotate(a)
			iq.Rotate(a)
			q.PushFront(a)
			iq.PushFront(b)
		}
	}
	// Find the zero
	index := q.Index(func(a int) bool { return a == 0 })
	q.Rotate(index + 1000)
	sum := q.Front()
	q.Rotate(1000)
	sum += q.Front()
	q.Rotate(1000)
	sum += q.Front()
	return sum
}
