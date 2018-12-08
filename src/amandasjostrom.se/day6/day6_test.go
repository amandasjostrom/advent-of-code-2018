package day6

import (
	"testing"
	"amandasjostrom.se/common"
)

func TestRun(t *testing.T) {
	common.TestRun(t, "\n", Run, 4011, 0)
}
func TestWithTestInput(t *testing.T) {
	result := Run(common.GetTestInputFromTest("\n"))
	want := 17
	if result.Part1 != want {
		t.Errorf("claims() = %v, want %v", result.Part1, want)
	}
}
