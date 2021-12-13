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
		// in - out
		{IntCode{input: 198, output: 0, pointer: 0, memory: []int{3, 0, 4, 0, 99}},
			IntCode{input: 198, output: 198, pointer: 4, memory: []int{198, 0, 4, 0, 99}}},
		// 99 to end
		{IntCode{input: 0, output: 0, pointer: 0, memory: []int{1002, 4, 3, 4, 33}},
			IntCode{input: 0, output: 0, pointer: 4, memory: []int{1002, 4, 3, 4, 99}}},
		// can be negative
		{IntCode{input: 0, output: 0, pointer: 0, memory: []int{1101, 100, -1, 4, 0}},
			IntCode{input: 0, output: 0, pointer: 4, memory: []int{1101, 100, -1, 4, 99}}},
		// equal to 8 - position - pass
		{IntCode{input: 8, output: 0, pointer: 0, memory: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}},
			IntCode{input: 8, output: 1, pointer: 8, memory: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, 1, 8}}},
		// equal to 8 - position - fail
		{IntCode{input: 88, output: 0, pointer: 0, memory: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}},
			IntCode{input: 88, output: 0, pointer: 8, memory: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, 0, 8}}},
		// less than 8 - position
		{IntCode{input: 7, output: 0, pointer: 0, memory: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}},
			IntCode{input: 7, output: 1, pointer: 8, memory: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, 1, 8}}},
		// equal to 8 - immediate
		{IntCode{input: 8, output: 0, pointer: 0, memory: []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}},
			IntCode{input: 8, output: 1, pointer: 8, memory: []int{3, 3, 1108, 1, 8, 3, 4, 3, 99}}},
		// less than 8 - immediate
		{IntCode{input: 7, output: 0, pointer: 0, memory: []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}},
			IntCode{input: 7, output: 1, pointer: 8, memory: []int{3, 3, 1107, 1, 8, 3, 4, 3, 99}}},
		// jump on 0 - position
		{IntCode{input: 0, output: 0, pointer: 0, memory: []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}},
			IntCode{input: 0, output: 0, pointer: 11, memory: []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, 0, 0, 1, 9}}},
		// jump on 0 - immediate
		{IntCode{input: 0, output: 0, pointer: 0, memory: []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}},
			IntCode{input: 0, output: 0, pointer: 11, memory: []int{3, 3, 1105, 0, 9, 1101, 0, 0, 12, 4, 12, 99, 0}}},
		// input below, at or above 8
		{IntCode{input: 7, output: 0, pointer: 0, memory: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}},
			IntCode{input: 7, output: 999, pointer: 46, memory: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 7, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}}},
		// input below, at or above 8
		{IntCode{input: 8, output: 0, pointer: 0, memory: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}},
			IntCode{input: 8, output: 1000, pointer: 46, memory: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 1000, 8, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}}},
		// input below, at or above 8
		{IntCode{input: 9, output: 0, pointer: 0, memory: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}},
			IntCode{input: 9, output: 1001, pointer: 46, memory: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 1001, 9, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}}},
	}
	for _, fixture := range fixtures {
		value := OpCode(fixture.Value)
		if !CompareIntCode(value, fixture.Expected) {
			t.Errorf("Got %v; want\n %v", value, fixture.Expected)
		}
	}
}
