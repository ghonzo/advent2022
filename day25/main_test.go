// Advent of Code 2022, Day 25
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
		want string
	}{
		{"example", args{common.ReadStringsFromFile("testdata/example.txt")}, "2=-1=0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.entries); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
