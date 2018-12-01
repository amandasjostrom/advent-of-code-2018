package main

import (
	"testing"
)

func Test_hasMoreThanOneDiff(t *testing.T) {
	type args struct {
		box1 string
		box2 string
	}
	tests := []struct {
		name                   string
		args                   args
		wantMoreThanOneDiff    bool
		wantCharactersInCommon string
	}{
		{"Finds exactly one difference",
			args{"amandd", "amanda"},
			false,
			"amand",
		},
		{"Finds no difference",
			args{"amaddd", "amanda"},
			true,
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMoreThanOneDiff, gotCharactersInCommon := hasMoreThanOneDiff(tt.args.box1, tt.args.box2)
			if gotMoreThanOneDiff != tt.wantMoreThanOneDiff {
				t.Errorf("hasMoreThanOneDiff() gotMoreThanOneDiff = %v, want %v", gotMoreThanOneDiff, tt.wantMoreThanOneDiff)
			}
			if gotCharactersInCommon != tt.wantCharactersInCommon {
				t.Errorf("hasMoreThanOneDiff() gotCharactersInCommon = %v, want %v", gotCharactersInCommon, tt.wantCharactersInCommon)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	type args struct {
		boxIds []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example input", args{[]string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.boxIds); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		boxIds []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Finds neighbour", args{[]string {"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}}, "fgij"},
		{"No neighbour", args{[]string {"abcde", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.boxIds); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
