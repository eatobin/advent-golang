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
		// equal to 8 - position
		{IntCode{input: 8, output: 0, pointer: 0, memory: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}},
			IntCode{input: 8, output: 1, pointer: 8, memory: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, 1, 8}}},
		// less than 8 - position
		{IntCode{input: 7, output: 0, pointer: 0, memory: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}},
			IntCode{input: 7, output: 1, pointer: 8, memory: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, 1, 8}}},
		// equal to 8 - immediate
		{IntCode{input: 8, output: 0, pointer: 0, memory: []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}},
			IntCode{input: 8, output: 1, pointer: 8, memory: []int{3, 3, 1108, 1, 8, 3, 4, 3, 99}}},
	}
	for _, fixture := range fixtures {
		value := OpCode(fixture.Value)
		if !CompareIntCode(value, fixture.Expected) {
			t.Errorf("Got %v; want\n %v", value, fixture.Expected)
		}
	}
}
