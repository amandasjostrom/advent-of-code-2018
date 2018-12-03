package day3

import (
	"reflect"
	"testing"
)

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

func Test_markOnFabric(t *testing.T) {
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
