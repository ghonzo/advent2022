// Advent of Code 2022, Day 7
package main

import (
	"fmt"
	"strings"

	"github.com/ghonzo/advent2022/common"
)

// Day 7: No Space Left On Device
// Part 1 answer: 1989474
// Part 2 answer: 1111607
func main() {
	fmt.Println("Advent of Code 2022, Day 7")
	entries := common.ReadStringsFromFile("input.txt")
	fmt.Printf("Part 1: %d\n", part1(entries))
	fmt.Printf("Part 2: %d\n", part2(entries))
}

type fd interface {
	Name() string
	Size() int
	Children() []fd
	Parent() *dir
}

type dir struct {
	name     string
	children []fd
	parent   *dir
}

type file struct {
	name   string
	size   int
	parent *dir
}

func (d *dir) Name() string {
	return d.name
}

func (d *dir) Size() int {
	sum := 0
	for _, c := range d.children {
		sum += c.Size()
	}
	return sum
}

func (d *dir) Children() []fd {
	return d.children
}

func (d *dir) Parent() *dir {
	return d.parent
}

func (d *dir) Cd(n string) *dir {
	for _, c := range d.children {
		if c.Name() == n {
			return c.(*dir)
		}
	}
	return nil
}

func (f *file) Name() string {
	return f.name
}

func (f *file) Size() int {
	return f.size
}

func (f *file) Children() []fd {
	return []fd{}
}

func (f *file) Parent() *dir {
	return f.parent
}

func part1(entries []string) int {
	root := new(dir)
	cwd := root
	for _, s := range entries {
		if strings.HasPrefix(s, "$ cd ") {
			childDir := strings.TrimPrefix(s, "$ cd ")
			if childDir == "/" {
				cwd = root
			} else if childDir == ".." {
				cwd = cwd.Parent()
			} else {
				cwd = cwd.Cd(childDir)
			}
		} else if s != "$ ls" {
			// Must be in the middle of a directory listing
			if strings.HasPrefix(s, "dir ") {
				cwd.children = append(cwd.children, &dir{name: strings.TrimPrefix(s, "dir "), parent: cwd})
			} else {
				size, n, ok := strings.Cut(s, " ")
				if !ok {
					panic("oops")
				}
				cwd.children = append(cwd.children, &file{name: n, size: common.Atoi(size), parent: cwd})
			}
		} else {
			// Must be ls ... panic if already have children
			if len(cwd.children) > 0 {
				panic("already have children")
			}
		}
	}
	sum := 0
	sumOfSmall(root, &sum)
	return sum
}

func sumOfSmall(x fd, sum *int) {
	if _, ok := x.(*dir); ok {
		if x.Size() <= 100000 {
			*sum += x.Size()
		}
		for _, sub := range x.Children() {
			sumOfSmall(sub, sum)
		}
	}
}

func part2(entries []string) int {
	root := new(dir)
	cwd := root
	for _, s := range entries {
		if strings.HasPrefix(s, "$ cd ") {
			childDir := strings.TrimPrefix(s, "$ cd ")
			if childDir == "/" {
				cwd = root
			} else if childDir == ".." {
				cwd = cwd.Parent()
			} else {
				cwd = cwd.Cd(childDir)
			}
		} else if s != "$ ls" {
			// Must be in the middle of a directory listing
			if strings.HasPrefix(s, "dir ") {
				cwd.children = append(cwd.children, &dir{name: strings.TrimPrefix(s, "dir "), parent: cwd})
			} else {
				size, n, ok := strings.Cut(s, " ")
				if !ok {
					panic("oops")
				}
				cwd.children = append(cwd.children, &file{name: n, size: common.Atoi(size), parent: cwd})
			}
		} else {
			// Must be ls ... panic if already have children
			if len(cwd.children) > 0 {
				panic("already have children")
			}
		}
	}
	unused := 70000000 - root.Size()
	need := 30000000 - unused
	dirSizes := make(map[*dir]int)
	addDirSizes(root, dirSizes)
	smallest := unused
	for _, v := range dirSizes {
		if v >= need && v < smallest {
			smallest = v
		}
	}
	return smallest
}

func addDirSizes(d *dir, sizeMap map[*dir]int) {
	sizeMap[d] = d.Size()
	for _, c := range d.Children() {
		if sub, ok := c.(*dir); ok {
			addDirSizes(sub, sizeMap)
		}
	}
}
