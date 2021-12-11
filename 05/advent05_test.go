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
		//{IntCode{pointer: 0, memory: []int{1, 0, 0, 0, 99}}, IntCode{pointer: 4, memory: []int{2, 0, 0, 0, 99}}},
		{IntCode{input: 198, output: 0, pointer: 0, memory: []int{3, 0, 4, 0, 99}},
			IntCode{input: 198, output: 198, pointer: 4, memory: []int{198, 0, 4, 0, 99}}},
	}
	for _, fixture := range fixtures {
		value := OpCode(fixture.Value)
		if !CompareIntCode(value, fixture.Expected) {
			t.Errorf("Got %v; want\n %v", value, fixture.Expected)
		}
	}
}
