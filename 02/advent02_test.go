package main

import (
	"reflect"
	"testing"
)

type Fixtures struct {
	Value    []int
	Expected []int
}

func TestOpCode(t *testing.T) {
	fixtures := []Fixtures{
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
		{
			[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			[]int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
		},
	}
	for _, fixture := range fixtures {
		value := make([]int, len(fixture.Value))
		copy(value, fixture.Value)
		OpCode(IntCode{pointer: 0, memory: value})
		if !reflect.DeepEqual(value, fixture.Expected) {
			t.Errorf("OpCode(IntCode{_, memory: %d}) =\nOpCode(IntCode{_, memory: %d}); want\n"+
				"OpCode(IntCode{_, memory: %d})", fixture.Value, value, fixture.Expected)
		}
	}
}
