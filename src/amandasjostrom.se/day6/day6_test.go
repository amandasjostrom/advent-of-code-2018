package day6

import (
	"testing"
	"amandasjostrom.se/common"
)

func TestRun(t *testing.T) {
	common.TestRun(t, "\n", Run, 4011, 46054)
}
func TestWithTestInput(t *testing.T) {
	result := Run(common.GetTestInputFromTest("\n"))
	want := 17
	if result.Part1 != want {
		t.Errorf("claims() = %v, want %v", result.Part1, want)
	}
	//Wont test part 2 for test data because other limit for size of region
}
