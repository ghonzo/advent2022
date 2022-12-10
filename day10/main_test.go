// Advent of Code 2022, Day 10
package main

import (
	"testing"

	"github.com/ghonzo/advent2022/common"
)

func Test_part1(t *testing.T) {
	type args struct {
		entries []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{common.ReadStringsFromFile("testdata/example.txt")}, 13140},
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
	}{
		{"example", args{common.ReadStringsFromFile("testdata/example.txt")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			part2(tt.args.entries)
		})
	}
}
