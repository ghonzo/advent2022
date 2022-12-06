// Advent of Code 2022, Day 6
package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{"mjqjpqmgbljsphdztnvjfqwrcgsmlb"}, 7},
		{"example 2", args{"bvwbjplbgvbhsrlpgdmjqwftvncz"}, 5},
		{"example 3", args{"nppdvjthqldpwncqszvftbrmjlhg"}, 6},
		{"example 4", args{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"}, 10},
		{"example 5", args{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.str); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{"mjqjpqmgbljsphdztnvjfqwrcgsmlb"}, 19},
		{"example 2", args{"bvwbjplbgvbhsrlpgdmjqwftvncz"}, 23},
		{"example 3", args{"nppdvjthqldpwncqszvftbrmjlhg"}, 23},
		{"example 4", args{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"}, 29},
		{"example 5", args{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"}, 26},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.str); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
