package main

// Instruction:
// ABCDE
// 01234
// 01002
// 34(DE) - two-digit opcode,      02 == opcode 2
//  2(C) - mode of 1st parameter,  0 == position mode
//  1(B) - mode of 2nd parameter,  1 == immediate mode
//  0(A) - mode of 3rd parameter,  0 == position mode,
//                                   omitted due to being a leading zero
// 0 1 or 2 = left-to-right position after 2 digit opcode
// p i or r = position, immediate or relative mode
// r or w = read or write

import (
	"fmt"
)

type Intcode struct {
	pointer int
	memory  [121]int
}

func main() {
	var intcode *Intcode
	var icReturn int

	intcode = newIntcode()
	icReturn = 1

	updatedMemory(intcode, 12, 2)

	for icReturn == 1 {
		icReturn = opcode(intcode)
	}

	fmt.Printf("\nPart A answer = %d. Correct = 2890696\n", intcode.memory[0])
	fmt.Printf("Part B answer = %d. Correct = 8226\n\n", nounVerb())
}

func newIntcode() *Intcode {
	intcode := Intcode{
		pointer: 0,
		memory: [121]int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 10, 1, 19, 2, 9, 19, 23, 2, 13, 23,
			27, 1, 6, 27, 31, 2, 6, 31, 35, 2, 13, 35, 39, 1, 39, 10, 43, 2, 43, 13, 47, 1, 9, 47,
			51, 1, 51, 13, 55, 1, 55, 13, 59, 2, 59, 13, 63, 1, 63, 6, 67, 2, 6, 67, 71, 1, 5, 71,
			75, 2, 6, 75, 79, 1, 5, 79, 83, 2, 83, 6, 87, 1, 5, 87, 91, 1, 6, 91, 95, 2, 95, 6, 99,
			1, 5, 99, 103, 1, 6, 103, 107, 1, 107, 2, 111, 1, 111, 5, 0, 99, 2, 14, 0, 0},
	}
	return &intcode
}

// For example, if your Intcode computer encounters 1,10,20,30,
// it should read the values at positions 10 and 20, add those values,
// and then overwrite the value at position 30 with their sum.

func opcode(intCode *Intcode) int {
	var action int
	var address1 int
	var address2 int
	var address3 int

	action = intCode.memory[intCode.pointer]     // action = 1
	address1 = intCode.memory[intCode.pointer+1] // address1 = 10 key = 10
	address2 = intCode.memory[intCode.pointer+2] // address2 = 20 key = 20
	address3 = intCode.memory[intCode.pointer+3] // address3 = 30 key = 30

	switch action {
	case 1:
		intCode.memory[address3] = // write to position 30
			intCode.memory[address1] + // read value at position 10
				intCode.memory[address2] // read value at position 20
		intCode.pointer += 4
		return 1
	case 2:
		intCode.memory[address3] =
			intCode.memory[address1] *
				intCode.memory[address2]
		intCode.pointer += 4
		return 1
	case 99:
		return 0
	default:
		panic("Something critical went wrong!")
	}
}

func updatedMemory(intcode *Intcode, noun int, verb int) {
	intcode.memory[1] = noun
	intcode.memory[2] = verb
}

func nounVerb() int {
	var noun int
	var verb int
	var intcode *Intcode
	var icReturn int
	var candidate int

	for noun = range 100 {
		for verb = range 100 {
			intcode = newIntcode()
			updatedMemory(intcode, noun, verb)

			icReturn = 1
			for icReturn == 1 {
				icReturn = opcode(intcode)
			}

			candidate = intcode.memory[0]
			if candidate == 19690720 {
				return 100*noun + verb
			}
		}
	}
	panic("Something critical went wrong!")
}
