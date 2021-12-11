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
		{IntCode{pointer: 0, memory: []int{1, 0, 0, 0, 99}}, IntCode{pointer: 4, memory: []int{2, 0, 0, 0, 99}}},
		//{IntCode{input: 198, output: 0, pointer: 0, memory: []int{3, 0, 4, 0, 99}},
		//	IntCode{input: 198, output: 198, pointer: 0, memory: []int{3, 0, 4, 0, 99}}},
		//{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		//{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		//{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
		//{
		//	[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
		//	[]int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		//},
	}
	for _, fixture := range fixtures {
		value := OpCode(fixture.Value)
		if !CompareIntCode(value, fixture.Expected) {
			t.Errorf("Got %v; want\n %v", value, fixture.Expected)
		}
	}
}
