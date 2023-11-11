// Package common provides common data structures and utility functions for Advent of Code
package common

import (
	"bufio"
	"io"
	"strings"
)

// Grid represents a mutable 2D rectangle with a value at each integer coordinate
type Grid interface {
	Size() Point
	Get(coord Point) byte
	CheckedGet(coord Point) (v byte, ok bool)
	Set(coord Point, b byte)
	AllPoints() <-chan Point
	Clone() Grid
}

// ArraysGrid is a Grid that has an underlying representation of [][]byte
type ArraysGrid [][]byte

// Size returns a Point that represents the dimensions of the Grid
func (g *ArraysGrid) Size() Point {
	return Point{len((*g)[0]), len(*g)}
}

// Get returns the value at the given coordinate
//
// If the coordinate is outside the dimensions of the Grid, this will throw an error. You may
// want to use CheckedGet instead
func (g *ArraysGrid) Get(coord Point) byte {
	return (*g)[coord.y][coord.x]
}

// CheckedGet returns the value at the given coordinate (if present), as well as an ok value.
//
// If the coordinate is present, it will return the value and an ok value of true. If it is not,
// it will return 0 and an ok value of false.
func (g *ArraysGrid) CheckedGet(coord Point) (byte, bool) {
	size := g.Size()
	if coord.x < 0 || coord.x >= size.x || coord.y < 0 || coord.y >= size.y {
		return 0, false
	}
	return g.Get(coord), true
}

// Set sets a value at the given coordinate.
//
// If the coordinate is outside the bounds, this will throw an error.
func (g *ArraysGrid) Set(coord Point, b byte) {
	(*g)[coord.y][coord.x] = b
}

// Clone returns a copy of the Grid, leaving the original untouched.
func (g *ArraysGrid) Clone() Grid {
	size := g.Size()
	clone := make(ArraysGrid, size.y)
	for row := range *g {
		clone[row] = make([]byte, size.x)
		copy(clone[row], (*g)[row])
	}
	return &clone
}

// AllPoints returns a channel of all the points in the Grid.
func (g *ArraysGrid) AllPoints() <-chan Point {
	ch := make(chan Point)
	go func() {
		size := g.Size()
		for y := 0; y < size.y; y++ {
			for x := 0; x < size.x; x++ {
				ch <- Point{x, y}
			}
		}
		close(ch)
	}()
	return ch
}

// Count returns the number of instances of the given value in the grid
func Count(grid Grid, v byte) int {
	var count int
	for pt := range grid.AllPoints() {
		if grid.Get(pt) == v {
			count++
		}
	}
	return count
}

// MapGridValues applies a mapping function to each value in the grid
func MapGridValues(grid Grid, mapFunc func(v byte) byte) {
	for pt := range grid.AllPoints() {
		grid.Set(pt, mapFunc(grid.Get(pt)))
	}
}

// ReadArraysGrid parses the lines and created a Grid with the y-dimension given by the number of lines
// and the x-dimension given by the length of the first line.
func ReadArraysGrid(r io.Reader) *ArraysGrid {
	var grid ArraysGrid
	input := bufio.NewScanner(r)
	for input.Scan() {
		grid = append(grid, []byte(input.Text()))
	}
	return &grid
}

// ArraysGridFromLines parses the lines and created a Grid with the y-dimension given by the number of lines
// and the x-dimension given by the length of the first line.
func ArraysGridFromLines(lines []string) *ArraysGrid {
	grid := make(ArraysGrid, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}
	return &grid
}

// NewArraysGrid initializes an empty grid with the given size.
func NewArraysGrid(x, y int) *ArraysGrid {
	grid := make(ArraysGrid, y)
	for row := range grid {
		grid[row] = make([]byte, x)
	}
	return &grid
}

// SparseGrid is a Grid that has an underlying representation of map[Point]byte
type SparseGrid map[Point]byte

func NewSparseGrid() SparseGrid {
	return make(SparseGrid)
}

// This returns a point that has the maxX and maxY, or 0,0 if empty
func (g SparseGrid) Size() Point {
	mmX := new(MaxMin[int])
	mmY := new(MaxMin[int])
	for p := range g.AllPoints() {
		mmX.Accept(p.X())
		mmY.Accept(p.Y())
	}
	return Point{mmX.Max, mmY.Max}
}

// Get returns the value at the given coordinate
//
// If the coordinate is not present or is outside the dimensions of the Grid, this will return 0.
func (g SparseGrid) Get(coord Point) byte {
	return g[coord]
}

// CheckedGet returns the value at the given coordinate (if present), as well as an ok value.
//
// If the coordinate is present, it will return the value and an ok value of true. If it is not,
// it will return 0 and an ok value of false.
func (g SparseGrid) CheckedGet(coord Point) (byte, bool) {
	v, ok := g[coord]
	return v, ok
}

// Set sets a value at the given coordinate.
func (g SparseGrid) Set(coord Point, b byte) {
	g[coord] = b
}

// AllPoints returns a channel of all the points in the Grid.
func (g SparseGrid) AllPoints() <-chan Point {
	ch := make(chan Point)
	go func() {
		for k := range g {
			ch <- k
		}
		close(ch)
	}()
	return ch
}

// Clone returns a copy of the Grid, leaving the original untouched.
func (g SparseGrid) Clone() Grid {
	clone := NewSparseGrid()
	for k, v := range g {
		clone[k] = v
	}
	return clone
}

// RenderGrid will render the contents of the grid as a multiline string. If you would like a character for a "missing point"
// (which is only possible for a SparseGrid), then specify a missing character. Otherwise we will use a space.
func RenderGrid(g Grid, missing ...byte) string {
	xMM := new(MaxMin[int])
	yMM := new(MaxMin[int])
	for p := range g.AllPoints() {
		xMM.Accept(p.X())
		yMM.Accept(p.Y())
	}
	var sb strings.Builder
	blank := byte(' ')
	if len(missing) > 0 {
		blank = missing[0]
	}
	for y := yMM.Min; y <= yMM.Max; y++ {
		for x := xMM.Min; x <= xMM.Max; x++ {
			v, ok := g.CheckedGet(NewPoint(x, y))
			if !ok {
				v = blank
			}
			sb.WriteByte(v)
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}
