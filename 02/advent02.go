package main

import (
	"advent-golang/intCodePkg"
	"fmt"
)

func opcode(intCode *intCodePkg.IntCode) int {
	var action int
	var address1 int
	var address2 int
	var address3 int

	action = intCode.Memory[intCode.Pointer]
	address1 = intCode.Memory[intCode.Pointer+1]
	address2 = intCode.Memory[intCode.Pointer+2]
	address3 = intCode.Memory[intCode.Pointer+3]

	switch action {
	case 1:
		intCode.Memory[address3] =
			intCode.Memory[address1] +
				intCode.Memory[address2]
		intCode.Pointer += 4
		return 1
	case 2:
		intCode.Memory[address3] =
			intCode.Memory[address1] *
				intCode.Memory[address2]
		intCode.Pointer += 4
		return 1
	default:
		return 0
	}
}

func updatedMemory(intcode *intCodePkg.IntCode, noun int, verb int) {
	intcode.Memory[1] = noun
	intcode.Memory[2] = verb
}

func nounVerb() int {
	var noun int
	var verb int
	var intcode intCodePkg.IntCode
	var icReturn int
	var candidate int

	for noun = 0; noun < 100; noun++ {
		for verb = 0; verb < 100; verb++ {
			intcode.Pointer = 0
			intcode.Memory = intCodePkg.MakeMemory("advent02.csv")
			updatedMemory(&intcode, noun, verb)

			icReturn = 1
			for icReturn == 1 {
				icReturn = opcode(&intcode)
			}

			candidate = intcode.Memory[0]
			if candidate == 19690720 {
				goto end
			}
		}
	}
end:
	return (100 * noun) + verb
}

func main() {
	var intcode intCodePkg.IntCode
	var icReturn int

	intcode.Pointer = 0
	intcode.Memory = intCodePkg.MakeMemory("advent02.csv")
	icReturn = 1

	updatedMemory(&intcode, 12, 2)

	for icReturn == 1 {
		icReturn = opcode(&intcode)
	}

	fmt.Printf("\nMemory length: %d\n\n", len(intcode.Memory))
	fmt.Printf("Part A answer = %d\n", intcode.Memory[0]) // Part A answer = 2890696
	fmt.Printf("Part B answer = %d\n", nounVerb())        // Part B answer = 8226
}
