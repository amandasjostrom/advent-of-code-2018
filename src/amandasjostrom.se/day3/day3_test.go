package day3

import (
	"reflect"
	"testing"
	"amandasjostrom.se/common"
)

func TestRun(t *testing.T) {
	common.TestRun(t,"\n", Run,105047, "#658")
}

func Test_claims(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want []Claim
	}{
		{"Normal input", args{[]string{"#240 @ 131,366: 21x14"}}, []Claim{{"#240", 131, 366, 21, 14}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := claims(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("claims() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_markClaimsOnFabric(t *testing.T) {
	type args struct {
		claims []Claim
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{"Handles first", args{[]Claim{{"id", 1, 2, 1, 1}}}, map[string]int{"1+2": 1}},
		{"Handles second", args{[]Claim{
			{"id", 1, 2, 1, 1},
			{"id", 1, 2, 1, 1},
		}}, map[string]int{"1+2": 2}},
		{"Handles overlapping areas", args{[]Claim{
			{"id", 1, 2, 2, 2},
			{"id", 1, 2, 1, 1},
		}}, map[string]int{
			"1+2": 2,
			"2+2": 1,
			"1+3": 1,
			"2+3": 1,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := markClaimsOnFabric(tt.args.claims); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("markClaimsOnFabric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isOverLapping(t *testing.T) {
	type args struct {
		claim  Claim
		fabric map[string]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"not overlapping 1x1", args{Claim{"id", 1,1,1,1}, map[string]int{"1+1":1}}, false },
		{"not overlapping 2x2", args{Claim{"id", 1,1,2,2}, map[string]int{"1+1":1, "1+2":1, "2+2":1, "2+1":1}}, false  },
		{"is overlapping 1x1", args{Claim{"id", 1,1,1,1}, map[string]int{"1+1":2}}, true  },
		{"is overlapping 2x2", args{Claim{"id", 1,1,2,2}, map[string]int{"1+1":2, "1+2":1, "2+2":1, "2+1":1}}, true  },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOverLapping(tt.args.claim, tt.args.fabric); got != tt.want {
				t.Errorf("isOverLapping() = %v, want %v", got, tt.want)
			}
		})
	}
}
