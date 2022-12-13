// Advent of Code 2022, Day 13
package main

import (
	"testing"

	"github.com/ghonzo/advent2022/common"
)

func Test_inOrder(t *testing.T) {
	type args struct {
		left  string
		right string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"[1,1,3,1,1]", "[1,1,10,1,1]"}, true},
		{"2", args{"[[1],[2,3,4]]", "[[1],4]"}, true},
		{"3", args{"[9]", "[[8,7,6]]"}, false},
		{"4", args{"[[4,4],4,4]", "[[4,4],4,4,4]"}, true},
		{"5", args{"[7,7,7,7]", "[7,7,7]"}, false},
		{"6", args{"[]", "[3]"}, true},
		{"7", args{"[[[]]]", "[[]]"}, false},
		{"8", args{"[1,[2,[3,[4,[5,6,7]]]],8,9]", "[1,[2,[3,[4,[5,6,0]]]],8,9]"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inOrder(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("inOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	type args struct {
		entries []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{common.ReadStringsFromFile("testdata/example.txt")}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.entries); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		entries []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{common.ReadStringsFromFile("testdata/example.txt")}, 140},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.entries); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
