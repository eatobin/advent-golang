package main

import (
	"github.com/eatobin/advent-golang/intcode"
	"testing"
)

type Fixtures struct {
	Value    intcode.IntCode
	Expected intcode.IntCode
}

// CompareIntCode compares two IntCodes
func CompareIntCode(a, b intcode.IntCode) bool {
	if &a == &b {
		return true
	}
	if a.Input != b.Input {
		return false
	}
	if a.Output != b.Output {
		return false
	}
	if a.Pointer != b.Pointer {
		return false
	}
	if len(a.Memory) != len(b.Memory) {
		return false
	}
	for i, v := range a.Memory {
		if v != b.Memory[i] {
			return false
		}
	}
	return true
}

func TestOpCode(t *testing.T) {
	fixtures := []Fixtures{
		// in - out
		{intcode.IntCode{Input: 198, Output: 0, Pointer: 0, Memory: []int{3, 0, 4, 0, 99}},
			intcode.IntCode{Input: 198, Output: 198, Pointer: 4, Memory: []int{198, 0, 4, 0, 99}}},
		// 99 to end
		{intcode.IntCode{Input: 0, Output: 0, Pointer: 0, Memory: []int{1002, 4, 3, 4, 33}},
			intcode.IntCode{Input: 0, Output: 0, Pointer: 4, Memory: []int{1002, 4, 3, 4, 99}}},
		// can be negative
		{intcode.IntCode{Input: 0, Output: 0, Pointer: 0, Memory: []int{1101, 100, -1, 4, 0}},
			intcode.IntCode{Input: 0, Output: 0, Pointer: 4, Memory: []int{1101, 100, -1, 4, 99}}},
		// equal to 8 - position - pass
		{intcode.IntCode{Input: 8, Output: 0, Pointer: 0, Memory: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}},
			intcode.IntCode{Input: 8, Output: 1, Pointer: 8, Memory: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, 1, 8}}},
		// equal to 8 - position - fail
		{intcode.IntCode{Input: 88, Output: 0, Pointer: 0, Memory: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}},
			intcode.IntCode{Input: 88, Output: 0, Pointer: 8, Memory: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, 0, 8}}},
		// less than 8 - position - pass
		{intcode.IntCode{Input: 7, Output: 0, Pointer: 0, Memory: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}},
			intcode.IntCode{Input: 7, Output: 1, Pointer: 8, Memory: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, 1, 8}}},
		// less than 8 - position - fail
		{intcode.IntCode{Input: 77, Output: 0, Pointer: 0, Memory: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}},
			intcode.IntCode{Input: 77, Output: 0, Pointer: 8, Memory: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, 0, 8}}},
		// equal to 8 - immediate - pass
		{intcode.IntCode{Input: 8, Output: 0, Pointer: 0, Memory: []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}},
			intcode.IntCode{Input: 8, Output: 1, Pointer: 8, Memory: []int{3, 3, 1108, 1, 8, 3, 4, 3, 99}}},
		// equal to 8 - immediate - fail
		{intcode.IntCode{Input: 88, Output: 0, Pointer: 0, Memory: []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}},
			intcode.IntCode{Input: 88, Output: 0, Pointer: 8, Memory: []int{3, 3, 1108, 0, 8, 3, 4, 3, 99}}},
		// less than 8 - immediate - pass
		{intcode.IntCode{Input: 7, Output: 0, Pointer: 0, Memory: []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}},
			intcode.IntCode{Input: 7, Output: 1, Pointer: 8, Memory: []int{3, 3, 1107, 1, 8, 3, 4, 3, 99}}},
		// less than 8 - immediate - fail
		{intcode.IntCode{Input: 77, Output: 0, Pointer: 0, Memory: []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}},
			intcode.IntCode{Input: 77, Output: 0, Pointer: 8, Memory: []int{3, 3, 1107, 0, 8, 3, 4, 3, 99}}},
		// jump on 0 - position
		{intcode.IntCode{Input: 0, Output: 0, Pointer: 0, Memory: []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}},
			intcode.IntCode{Input: 0, Output: 0, Pointer: 11, Memory: []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, 0, 0, 1, 9}}},
		// jump on 0 - immediate
		{intcode.IntCode{Input: 0, Output: 0, Pointer: 0, Memory: []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}},
			intcode.IntCode{Input: 0, Output: 0, Pointer: 11, Memory: []int{3, 3, 1105, 0, 9, 1101, 0, 0, 12, 4, 12, 99, 0}}},
		// Input below, at or above 8
		{intcode.IntCode{Input: 7, Output: 0, Pointer: 0, Memory: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}},
			intcode.IntCode{Input: 7, Output: 999, Pointer: 46, Memory: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 7, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}}},
		// Input below, at or above 8
		{intcode.IntCode{Input: 8, Output: 0, Pointer: 0, Memory: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}},
			intcode.IntCode{Input: 8, Output: 1000, Pointer: 46, Memory: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 1000, 8, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}}},
		// Input below, at or above 8
		{intcode.IntCode{Input: 9, Output: 0, Pointer: 0, Memory: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}},
			intcode.IntCode{Input: 9, Output: 1001, Pointer: 46, Memory: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 1001, 9, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}}},
	}

	for _, fixture := range fixtures {
		icReturn := 1
		for icReturn == 1 {
			icReturn = fixture.Value.opCode()
		}
		if !CompareIntCode(fixture.Value, fixture.Expected) {
			t.Errorf("Got %v; want\n %v", fixture.Value, fixture.Expected)
		}
	}
}
