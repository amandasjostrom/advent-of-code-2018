package common

import (
	"reflect"
	"testing"
)

type runner func([]string) Result

func TestRun(t *testing.T, separator string, run runner, part1 interface{}, part2 interface{} ) {
	test(t,separator, run,Result{part1,part2})
}

func test(t *testing.T, separator string, run runner, want Result ) {
	if got := run(GetRealInputFromTest(separator)); !reflect.DeepEqual(got, want) {
		t.Errorf("Run() = %v, want %v", got, want)
	}
}
