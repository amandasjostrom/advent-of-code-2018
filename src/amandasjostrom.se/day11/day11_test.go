package day11

import (
	"testing"
	"amandasjostrom.se/common"
)

func TestRun(t *testing.T) {
	common.TestRun(t, "\n", Run, "20,51", "unknown")
}

func TestWithTestInput(t *testing.T) {
	result := Run([]string{"18"})
	wantPart1 := "33,45"
	wantPart2 := "90,269,16"
	if result.Part1 != wantPart1 {
		t.Errorf("claims() = %v, want %v", result.Part1, wantPart1)
	}
	if result.Part2 != wantPart2 {
		t.Errorf("claims() = %v, want %v", result.Part2, wantPart2)
	}
}

func Test_cellPower(t *testing.T) {
	type args struct {
		x                int
		y                int
		gridSerialNumber int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Fuel cell at  122,79, grid serial number 57: power level -5.", args{122, 79, 57}, -5},
		{"Fuel cell at 217,196, grid serial number 39: power level  0.", args{217, 196, 39}, 0},
		{"Fuel cell at 101,153, grid serial number 71: power level  4", args{101, 153, 71}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cellPower(tt.args.x, tt.args.y, tt.args.gridSerialNumber); got != tt.want {
				t.Errorf("cellPower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_squarePower(t *testing.T) {
	type args struct {
		topLeftX         int
		topLeftY         int
		gridSerialNumber int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"For grid serial number 18, the largest total 3x3 square has a top-left corner of 33,45 (with a total power of 29)",
			args{33, 45, 18}, 29},
		{"For grid serial number 42, the largest 3x3 square's top-left is 21,61 (with a total power of 30)",
			args{21, 61, 42}, 30},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := squarePower3By3(tt.args.topLeftX, tt.args.topLeftY, tt.args.gridSerialNumber); got != tt.want {
				t.Errorf("squarePower3By3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_squarePowerAnySize(t *testing.T) {
	type args struct {
		topLeftX         int
		topLeftY         int
		gridSerialNumber int
	}
	tests := []struct {
		name string
		args args
		power int
		size int
	}{
		{"For grid serial number 18, the largest total square (with a total power of 113) is 16x16 and has a top-left corner of 90,269",
			args{90, 269, 18}, 113, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPower, gotSize := squarePowerAnySize(tt.args.topLeftX, tt.args.topLeftY, tt.args.gridSerialNumber); gotPower != tt.power {
				t.Errorf("squarePowerAnySize() = %v, want %v", gotPower, tt.power)
				if gotSize != tt.size {
					t.Errorf("squarePowerAnySize() = %v, want %v", gotSize, tt.size)
				}
			}
		})
	}
}
