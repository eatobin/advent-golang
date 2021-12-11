package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	expected := OpCode(IntCode{input: 19, output: 19, pointer: 0, memory: []int{3, 0, 4, 0, 99}})
	value := OpCode(IntCode{input: 19, output: 0, pointer: 0, memory: []int{3, 0, 4, 0, 99}})
	if value.output != expected.output {
		t.Errorf("Expected value.output == 19, but got = %v", value.output)
	}
}

//	r = []int{1101, 100, -1, 4, 0}
//	Run(r, nil)
//	if r[4] != 99 {
//		t.Errorf("Expected r[4] == 99, but r = %v", r)
//	}
//}
//
//func TestInputOutput(t *testing.T) {
//	inputs := []int{666}
//	r := []int{3, 0, 4, 0, 99}
//	outputs := Run(r, inputs)
//	if len(outputs) != 1 || outputs[0] != 666 {
//		t.Errorf("Expected outputs == [666], but outputs = %v", outputs)
//	}
//}
//
//func TestEq8(t *testing.T) {
//	r := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
//
//	outputs := Run(r, []int{666})
//	if len(outputs) != 1 || outputs[0] != 0 {
//		t.Errorf("Expected outputs == [0], but outputs = %v", outputs)
//	}
//
//	outputs = Run(r, []int{8})
//	if len(outputs) != 1 || outputs[0] != 1 {
//		t.Errorf("Expected outputs == [1], but outputs = %v", outputs)
//	}
//}
//
//func TestLt8(t *testing.T) {
//	r := []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
//
//	outputs := Run(r, []int{8})
//	if len(outputs) != 1 || outputs[0] != 0 {
//		t.Errorf("Expected outputs == [0], but outputs = %v", outputs)
//	}
//
//	outputs = Run(r, []int{7})
//	if len(outputs) != 1 || outputs[0] != 1 {
//		t.Errorf("Expected outputs == [1], but outputs = %v", outputs)
//	}
//}
//
//func TestEq8Im(t *testing.T) {
//	r := []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}
//
//	outputs := Run(r, []int{666})
//	if len(outputs) != 1 || outputs[0] != 0 {
//		t.Errorf("Expected outputs == [0], but outputs = %v", outputs)
//	}
//
//	outputs = Run(r, []int{8})
//	if len(outputs) != 1 || outputs[0] != 1 {
//		t.Errorf("Expected outputs == [1], but outputs = %v", outputs)
//	}
//}
//
//func TestLt8Im(t *testing.T) {
//	r := []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}
//
//	outputs := Run(r, []int{8})
//	if len(outputs) != 1 || outputs[0] != 0 {
//		t.Errorf("Expected outputs == [0], but outputs = %v", outputs)
//	}
//
//	outputs = Run(r, []int{7})
//	if len(outputs) != 1 || outputs[0] != 1 {
//		t.Errorf("Expected outputs == [1], but outputs = %v", outputs)
//	}
//}
//
//func TestJump(t *testing.T) {
//	r := []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
//
//	outputs := Run(r, []int{0})
//	if len(outputs) != 1 || outputs[0] != 0 {
//		t.Errorf("Expected outputs == [0], but outputs = %v", outputs)
//	}
//
//	outputs = Run(r, []int{-1})
//	if len(outputs) != 1 || outputs[0] != 1 {
//		t.Errorf("Expected outputs == [1], but outputs = %v", outputs)
//	}
//}
//
//func TestJumpIm(t *testing.T) {
//	r := []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
//
//	outputs := Run(r, []int{0})
//	if len(outputs) != 1 || outputs[0] != 0 {
//		t.Errorf("Expected outputs == [0], but outputs = %v", outputs)
//	}
//
//	outputs = Run(r, []int{-1})
//	if len(outputs) != 1 || outputs[0] != 1 {
//		t.Errorf("Expected outputs == [1], but outputs = %v", outputs)
//	}
//}
//
//func TestBigger(t *testing.T) {
//	r := []int{
//		3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
//		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
//		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
//
//	outputs := Run(r, []int{7})
//	if len(outputs) != 1 || outputs[0] != 999 {
//		t.Errorf("Expected outputs == [999], but outputs = %v", outputs)
//	}
//
//	outputs = Run(r, []int{8})
//	if len(outputs) != 1 || outputs[0] != 1000 {
//		t.Errorf("Expected outputs == [1000], but outputs = %v", outputs)
//	}
//
//	outputs = Run(r, []int{9})
//	if len(outputs) != 1 || outputs[0] != 1001 {
//		t.Errorf("Expected outputs == [1001], but outputs = %v", outputs)
//	}
//}
