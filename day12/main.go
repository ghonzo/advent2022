// Advent of Code 2022, Day 12
package main

import (
	"fmt"

	"github.com/ghonzo/advent2022/common"
)

// Day 12: Hill Climbing Algorithm
// Part 1 answer: 484
// Part 2 answer: 478
func main() {
	fmt.Println("Advent of Code 2022, Day 12")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

type gamestate struct {
	pos   common.Point
	steps int
}

func part1(entries []string) int {
	g := common.ArraysGridFromLines(entries)
	var start, end common.Point
	for p := range g.AllPoints() {
		v := g.Get(p)
		if v == 'S' {
			start = p
			g.Set(p, 'a')
		} else if v == 'E' {
			end = p
			g.Set(p, 'z')
		}
	}
	state := &gamestate{pos: start}
	return findNextState(state, g, end, 999999)
}

var dirs = []common.Point{common.U, common.L, common.R, common.D}

var globalMinSteps = make(map[common.Point]int)

func findNextState(state *gamestate, g common.Grid, end common.Point, minSteps int) int {
	if state.pos == end {
		return state.steps
	}
	// Bad path
	if state.steps >= minSteps {
		return minSteps
	}
	if min, ok := globalMinSteps[state.pos]; !ok || state.steps < min {
		globalMinSteps[state.pos] = state.steps
	} else {
		return minSteps
	}
	mm := new(common.MaxMin[int])
	mm.Accept(minSteps)
	currentElevation := g.Get(state.pos)
	for _, dir := range dirs {
		np := state.pos.Add(dir)
		if v, ok := g.CheckedGet(np); ok && v <= currentElevation+1 {
			// We can move
			newState := &gamestate{pos: np, steps: state.steps + 1}
			mm.Accept(findNextState(newState, g, end, mm.Min))
		}
	}
	return mm.Min
}

var globalMinStepsDown = make(map[common.Point]int)

func part2(entries []string) int {
	g := common.ArraysGridFromLines(entries)
	var end common.Point
	for p := range g.AllPoints() {
		v := g.Get(p)
		if v == 'S' {
			g.Set(p, 'a')
		} else if v == 'E' {
			end = p
			g.Set(p, 'z')
		}
	}
	state := &gamestate{pos: end}
	return findNextStateDown(state, g, 999999)
}

func findNextStateDown(state *gamestate, g common.Grid, minSteps int) int {
	currentElevation := g.Get(state.pos)
	if currentElevation == 'a' {
		return state.steps
	}
	// Bad path
	if state.steps >= minSteps {
		return minSteps
	}
	if min, ok := globalMinStepsDown[state.pos]; !ok || state.steps < min {
		globalMinStepsDown[state.pos] = state.steps
	} else {
		return minSteps
	}
	mm := new(common.MaxMin[int])
	mm.Accept(minSteps)
	for _, dir := range dirs {
		np := state.pos.Add(dir)
		if v, ok := g.CheckedGet(np); ok && v >= currentElevation-1 {
			// We can move
			newState := &gamestate{pos: np, steps: state.steps + 1}
			mm.Accept(findNextStateDown(newState, g, mm.Min))
		}
	}
	return mm.Min
}
