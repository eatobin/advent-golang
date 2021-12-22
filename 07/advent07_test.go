package main

import (
	"sort"
	"testing"
)

type Fixtures struct {
	Phase    [][]int
	Memory   Memory
	Expected int
}

func TestOpCode(t *testing.T) {
	fixtures := []Fixtures{
		{[][]int{{4, 3, 2, 1, 0}}, []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}, 43210},
		{[][]int{{0, 1, 2, 3, 4}}, []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}, 54321},
	}

	for _, fixture := range fixtures {
		answer := passes(fixture.Phase, fixture.Memory)
		sort.Ints(answer)
		got := answer[len(answer)-1]
		if got != fixture.Expected {
			t.Errorf("Got %v; want\n %v", got, fixture.Expected)
		}
	}
}
