package main

import (
	"fmt"
)

// Instruction:
// ABCDE
// 01234
// 01002
// 34 - two-digit opcode,      02 == opcode 2
//  2 - mode of 1st parameter,  0 == position mode
//  1 - mode of 2nd parameter,  1 == immediate mode
//  0 - mode of 3rd parameter,  0 == position mode,
//                                   omitted due to being a leading zero
// 0 1 or 2 = left-to-right position after 2 digit opcode
// p i or r = position, immediate or relative mode
// r or w = read or write

type Intcode struct {
	input   int
	output  int
	pointer int
	memory  [11]int
}

type Instruction []int

func main() {
	var intcode Intcode
	//var icReturn int

	intcode = makeIntcode()
	//icReturn = 1
	//
	//for icReturn == 1 {
	//	icReturn = opcode(&intcode)
	//}

	fmt.Printf("\nPart A answer = %d. Correct = 2890696\n", intcode.memory[0])
}

func makeIntcode() Intcode {
	intcode := Intcode{
		input:   0,
		output:  0,
		pointer: 0,
		memory:  [11]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
	}
	return intcode
}

//func opcode(intCode *Intcode) int {
//	var action int
//	var address1 int
//	var address2 int
//	var address3 int
//
//	action = intCode.memory[intCode.pointer]
//	address1 = intCode.memory[intCode.pointer+1]
//	address2 = intCode.memory[intCode.pointer+2]
//	address3 = intCode.memory[intCode.pointer+3]
//
//	switch action {
//	case 1:
//		intCode.memory[address3] =
//			intCode.memory[address1] +
//				intCode.memory[address2]
//		intCode.pointer += 4
//		return 1
//	case 2:
//		intCode.memory[address3] =
//			intCode.memory[address1] *
//				intCode.memory[address2]
//		intCode.pointer += 4
//		return 1
//	default:
//		return 0
//	}
//}
