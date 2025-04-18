package main

import (
	"fmt"
	"sort"
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

type IntCode struct {
	input     int
	output    int
	phase     int
	pointer   int
	memory    [523]int
	isStopped bool
	doesRecur bool
}

const offsetC int = 1
const offsetB int = 2
const offsetA int = 3

var memoryConstant = [523]int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 38, 55, 72, 93, 118, 199, 280, 361, 442, 99999, 3, 9, 1001, 9, 2, 9, 1002, 9, 5, 9, 101, 4, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 3, 9, 1001, 9, 5, 9, 1002, 9, 4, 9, 4, 9, 99, 3, 9, 101, 4, 9, 9, 1002, 9, 3, 9, 1001, 9, 4, 9, 4, 9, 99, 3, 9, 1002, 9, 4, 9, 1001, 9, 4, 9, 102, 5, 9, 9, 1001, 9, 4, 9, 4, 9, 99, 3, 9, 101, 3, 9, 9, 1002, 9, 3, 9, 1001, 9, 3, 9, 102, 5, 9, 9, 101, 4, 9, 9, 4, 9, 99, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99}

func pad5(op int, instruction *[5]int) *[5]int {
	asString := fmt.Sprintf("%05d", op)
	asBytes := []byte(asString)
	for i := range 5 {
		(*instruction)[i] = int(asBytes[i] - 48)
	}
	return instruction
}

func aParam(icP *IntCode, instruction *[5]int) int {
	switch instruction[0] {
	case 0: // a-p-w
		return icP.memory[icP.pointer+offsetA]
	default:
		panic("aParam is not valid")
	}
}

func bParam(icP *IntCode, instruction *[5]int) int {
	switch instruction[1] {
	case 0: // b-p-r
		return icP.memory[icP.memory[icP.pointer+offsetB]]
	case 1: // b-i-r
		return icP.memory[icP.pointer+offsetB]
	default:
		panic("bParam is not valid")
	}
}

func cParam(icP *IntCode, instruction *[5]int) int {
	if instruction[4] == 3 {
		switch instruction[2] {
		case 0: // c-p-w
			return icP.memory[icP.pointer+offsetC]
		default:
			panic("cParam is not valid")
		}
	}
	switch instruction[2] {
	case 0: // c-p-r
		return icP.memory[icP.memory[icP.pointer+offsetC]]
	case 1: // c-i-r
		return icP.memory[icP.pointer+offsetC]
	default:
		panic("cParam is not valid")
	}
}

var candidates [][]int

func addAPerm(perm []int) {
	tmp := make([]int, len(perm))
	copy(tmp, perm)
	candidates = append(candidates, tmp)
}

func permutations(k int, A []int) {
	if k == 1 {
		addAPerm(A)
	} else {
		for i := 0; i < k-1; i++ {
			permutations(k-1, A)
			if k%2 == 0 {
				A[i], A[k-1] = A[k-1], A[i]
			} else {
				A[0], A[k-1] = A[k-1], A[0]
			}
		}
		permutations(k-1, A)
	}
}

func pass(candidate []int, instruction *[5]int) int {
	memA := memoryConstant
	memB := memoryConstant
	memC := memoryConstant
	memD := memoryConstant
	memE := memoryConstant
	icpA := IntCode{
		input:     0,
		output:    0,
		phase:     candidate[0],
		pointer:   0,
		memory:    memA,
		isStopped: false,
		doesRecur: true,
	}

	icReturn := 1
	for icReturn == 1 {
		icReturn = opcode(&icpA, instruction)
	}

	icpB := IntCode{
		input:     icpA.output,
		output:    0,
		phase:     candidate[1],
		pointer:   0,
		memory:    memB,
		isStopped: false,
		doesRecur: true,
	}

	icReturn = 1
	for icReturn == 1 {
		icReturn = opcode(&icpB, instruction)
	}

	icpC := IntCode{
		input:     icpB.output,
		output:    0,
		phase:     candidate[2],
		pointer:   0,
		memory:    memC,
		isStopped: false,
		doesRecur: true,
	}

	icReturn = 1
	for icReturn == 1 {
		icReturn = opcode(&icpC, instruction)
	}

	icpD := IntCode{
		input:     icpC.output,
		output:    0,
		phase:     candidate[3],
		pointer:   0,
		memory:    memD,
		isStopped: false,
		doesRecur: true,
	}

	icReturn = 1
	for icReturn == 1 {
		icReturn = opcode(&icpD, instruction)
	}

	icpE := IntCode{
		input:     icpD.output,
		output:    0,
		phase:     candidate[4],
		pointer:   0,
		memory:    memE,
		isStopped: false,
		doesRecur: true,
	}

	icReturn = 1
	for icReturn == 1 {
		icReturn = opcode(&icpE, instruction)
	}

	return icpE.output
}

func pass2(candidate []int, instruction *[5]int) int {
	memA := memoryConstant
	memB := memoryConstant
	memC := memoryConstant
	memD := memoryConstant
	memE := memoryConstant
	eOutput := 0
	allStopped := false
	icpA := IntCode{
		input:     0,
		output:    0,
		phase:     candidate[0],
		pointer:   0,
		memory:    memA,
		isStopped: false,
		doesRecur: false,
	}
	icpB := IntCode{
		input:     0,
		output:    0,
		phase:     candidate[1],
		pointer:   0,
		memory:    memB,
		isStopped: false,
		doesRecur: false,
	}
	icpC := IntCode{
		input:     0,
		output:    0,
		phase:     candidate[2],
		pointer:   0,
		memory:    memC,
		isStopped: false,
		doesRecur: false,
	}
	icpD := IntCode{
		input:     0,
		output:    0,
		phase:     candidate[3],
		pointer:   0,
		memory:    memD,
		isStopped: false,
		doesRecur: false,
	}
	icpE := IntCode{
		input:     0,
		output:    0,
		phase:     candidate[4],
		pointer:   0,
		memory:    memE,
		isStopped: false,
		doesRecur: false,
	}

	for !allStopped {
		icReturn := 1
		for icReturn == 1 {
			icReturn = opcode(&icpA, instruction)
		}
		icpB.input = icpA.output
		icReturn = 1
		for icReturn == 1 {
			icReturn = opcode(&icpB, instruction)
		}
		icpC.input = icpB.output
		icReturn = 1
		for icReturn == 1 {
			icReturn = opcode(&icpC, instruction)
		}
		icpD.input = icpC.output
		icReturn = 1
		for icReturn == 1 {
			icReturn = opcode(&icpD, instruction)
		}
		icpE.input = icpD.output
		icReturn = 1
		for icReturn == 1 {
			icReturn = opcode(&icpE, instruction)
		}

		icpA.input = icpE.output
		eOutput = icpE.output
		allStopped = icpE.isStopped
	}

	return eOutput
}

func passes(instruction *[5]int) []int {
	vcm := make([]int, len(candidates))
	for i, v := range candidates {
		vcm[i] = pass(v, instruction)
	}
	return vcm
}

func passes2(instruction *[5]int) []int {
	vcm := make([]int, len(candidates))
	for i, v := range candidates {
		vcm[i] = pass2(v, instruction)
	}
	return vcm
}

func opcode(icP *IntCode, instruction *[5]int) int {
	if icP.isStopped {
		return 0
	} else {
		instruction = pad5(icP.memory[icP.pointer], instruction)
		if instruction[3] == 9 {
			icP.isStopped = true
			return 0
		} else {
			switch instruction[4] {
			case 1:
				icP.memory[aParam(icP, instruction)] = bParam(icP, instruction) + cParam(icP, instruction)
				icP.pointer += 4
				return 1
			case 2:
				icP.memory[aParam(icP, instruction)] = bParam(icP, instruction) * cParam(icP, instruction)
				icP.pointer += 4
				return 1
			case 3:
				if icP.phase == -1 {
					icP.memory[cParam(icP, instruction)] = icP.input
				} else {
					if icP.pointer == 0 {
						icP.memory[cParam(icP, instruction)] = icP.phase
					} else {
						icP.memory[cParam(icP, instruction)] = icP.input
					}
				}
				icP.pointer += 2
				return 1
			case 4:
				if icP.doesRecur {
					icP.output = cParam(icP, instruction)
					icP.pointer += 2
					return 1
				} else {
					icP.output = cParam(icP, instruction)
					icP.pointer += 2
					return 0
				}
			case 5:
				if cParam(icP, instruction) != 0 {
					icP.pointer = bParam(icP, instruction)
				} else {
					icP.pointer += 3
				}
				return 1
			case 6:
				if cParam(icP, instruction) == 0 {
					icP.pointer = bParam(icP, instruction)
				} else {
					icP.pointer += 3
				}
				return 1
			case 7:
				if cParam(icP, instruction) < bParam(icP, instruction) {
					icP.memory[aParam(icP, instruction)] = 1
				} else {
					icP.memory[aParam(icP, instruction)] = 0
				}
				icP.pointer += 4
				return 1
			case 8:
				if cParam(icP, instruction) == bParam(icP, instruction) {
					icP.memory[aParam(icP, instruction)] = 1
				} else {
					icP.memory[aParam(icP, instruction)] = 0
				}
				icP.pointer += 4
				return 1
			default:
				return 0
			}
		}
	}
}

func main() {
	phases := []int{0, 1, 2, 3, 4}
	permutations(len(phases), phases)
	instruction := [5]int{}
	answer := passes(&instruction)
	sort.Ints(answer)
	fmt.Printf("Part A answer = %d. Correct = 368584\n", answer[len(answer)-1])

	candidates = nil
	phases = []int{5, 6, 7, 8, 9}
	permutations(len(phases), phases)
	instruction = [5]int{}
	answer = passes2(&instruction)
	sort.Ints(answer)
	fmt.Printf("Part B answer = %d. Correct = 35993240\n", answer[len(answer)-1])
}
