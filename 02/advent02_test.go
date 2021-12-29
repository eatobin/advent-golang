package main

import (
	"testing"
)

type Fixtures struct {
	Value    IntCode
	Expected IntCode
}

func TestOpCode(t *testing.T) {
	fixtures := []Fixtures{
		{Value: IntCode{pointer: 0, memory: map[int]int{0: 1, 1: 0, 2: 0, 3: 0, 4: 99}}, Expected: IntCode{pointer: 4, memory: map[int]int{0: 2, 1: 0, 2: 0, 3: 0, 4: 99}}},
		{IntCode{pointer: 0, memory: map[int]int{0: 2, 1: 3, 2: 0, 3: 3, 4: 99}}, IntCode{pointer: 4, memory: map[int]int{0: 2, 1: 3, 2: 0, 3: 6, 4: 99}}},
		//{IntCode{pointer: 0, memory: []int{2, 4, 4, 5, 99, 0}}, IntCode{pointer: 4, memory: []int{2, 4, 4, 5, 99, 9801}}},
		//{IntCode{pointer: 0, memory: []int{1, 1, 1, 4, 99, 5, 6, 0, 99}}, IntCode{pointer: 8, memory: []int{30, 1, 1, 4, 2, 5, 6, 0, 99}}},
		//{IntCode{pointer: 0, memory: []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}},
		//	IntCode{pointer: 8, memory: []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}}},
	}

	for _, fixture := range fixtures {
		icReturn := 1
		for icReturn == 1 {
			icReturn = OpCode(&fixture.Value)
		}
		if !CompareIntCode(fixture.Value, fixture.Expected) {
			t.Errorf("Got %v; want\n %v", fixture.Value, fixture.Expected)
		}
	}
}
