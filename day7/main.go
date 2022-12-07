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
	root := constructFileSystem(entries)
	dirSizes := make(map[*dir]int)
	addDirSizes(root, dirSizes)
	var sum int
	for _, v := range dirSizes {
		if v <= 100000 {
			sum += v
		}
	}
	return sum
}

func constructFileSystem(entries []string) *dir {
	root := new(dir)
	cwd := root
	for _, s := range entries {
		if strings.HasPrefix(s, "$ cd ") {
			d := strings.TrimPrefix(s, "$ cd ")
			if d == "/" {
				cwd = root
			} else if d == ".." {
				cwd = cwd.Parent()
			} else {
				cwd = cwd.Cd(d)
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
		}
	}
	return root
}

func addDirSizes(d *dir, sizeMap map[*dir]int) {
	sizeMap[d] = d.Size()
	for _, c := range d.Children() {
		if sub, ok := c.(*dir); ok {
			addDirSizes(sub, sizeMap)
		}
	}
}

func part2(entries []string) int {
	root := constructFileSystem(entries)
	unused := 70000000 - root.Size()
	need := 30000000 - unused
	dirSizes := make(map[*dir]int)
	addDirSizes(root, dirSizes)
	smallest := 70000000
	for _, v := range dirSizes {
		if v >= need && v < smallest {
			smallest = v
		}
	}
	return smallest
}
