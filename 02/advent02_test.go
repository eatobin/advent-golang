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
		{IntCode{pointer: 0, memory: map[int]int{0: 2, 1: 4, 2: 4, 3: 5, 4: 99, 5: 0}}, IntCode{pointer: 4, memory: map[int]int{0: 2, 1: 4, 2: 4, 3: 5, 4: 99, 5: 9801}}},
		{IntCode{pointer: 0, memory: map[int]int{0: 1, 1: 1, 2: 1, 3: 4, 4: 99, 5: 5, 6: 6, 7: 0, 8: 99}}, IntCode{pointer: 8, memory: map[int]int{0: 30, 1: 1, 2: 1, 3: 4, 4: 2, 5: 5, 6: 6, 7: 0, 8: 99}}},
		{IntCode{pointer: 0, memory: map[int]int{0: 1, 1: 9, 2: 10, 3: 3, 4: 2, 5: 3, 6: 11, 7: 0, 8: 99, 9: 30, 10: 40, 11: 50}},
			IntCode{pointer: 8, memory: map[int]int{0: 3500, 1: 9, 2: 10, 3: 70, 4: 2, 5: 3, 6: 11, 7: 0, 8: 99, 9: 30, 10: 40, 11: 50}}},
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
