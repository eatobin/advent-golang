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

//const fp = "advent07.csv"

type IntCode struct {
	input     int
	output    int
	phase     int
	pointer   int
	memory    [678]int
	isStopped bool
	doesRecur bool
}

const offsetC int = 1
const offsetB int = 2
const offsetA int = 3

var memoryConstant = [678]int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1102, 79, 14, 225, 1101, 17, 42, 225, 2, 74, 69, 224, 1001, 224, -5733, 224, 4, 224, 1002, 223, 8, 223, 101, 4, 224, 224, 1, 223, 224, 223, 1002, 191, 83, 224, 1001, 224, -2407, 224, 4, 224, 102, 8, 223, 223, 101, 2, 224, 224, 1, 223, 224, 223, 1101, 18, 64, 225, 1102, 63, 22, 225, 1101, 31, 91, 225, 1001, 65, 26, 224, 101, -44, 224, 224, 4, 224, 102, 8, 223, 223, 101, 3, 224, 224, 1, 224, 223, 223, 101, 78, 13, 224, 101, -157, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 3, 224, 1, 224, 223, 223, 102, 87, 187, 224, 101, -4698, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 4, 224, 1, 223, 224, 223, 1102, 79, 85, 224, 101, -6715, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 2, 224, 1, 224, 223, 223, 1101, 43, 46, 224, 101, -89, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 1, 224, 224, 1, 223, 224, 223, 1101, 54, 12, 225, 1102, 29, 54, 225, 1, 17, 217, 224, 101, -37, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 3, 224, 1, 223, 224, 223, 1102, 20, 53, 225, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 107, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 329, 101, 1, 223, 223, 1108, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 344, 101, 1, 223, 223, 7, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 359, 101, 1, 223, 223, 108, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 374, 101, 1, 223, 223, 8, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 389, 101, 1, 223, 223, 1108, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 404, 101, 1, 223, 223, 1007, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 419, 101, 1, 223, 223, 8, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 434, 1001, 223, 1, 223, 1008, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 449, 1001, 223, 1, 223, 1008, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 464, 101, 1, 223, 223, 1107, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 479, 101, 1, 223, 223, 107, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 494, 1001, 223, 1, 223, 1107, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 509, 101, 1, 223, 223, 1108, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 524, 101, 1, 223, 223, 7, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 539, 101, 1, 223, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 554, 101, 1, 223, 223, 8, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 569, 1001, 223, 1, 223, 1008, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 584, 101, 1, 223, 223, 107, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 599, 1001, 223, 1, 223, 7, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 614, 101, 1, 223, 223, 1007, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 629, 101, 1, 223, 223, 1107, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 644, 101, 1, 223, 223, 108, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 659, 101, 1, 223, 223, 1007, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 674, 101, 1, 223, 223, 4, 223, 99, 226}

func makeIntcode() IntCode {
	intcode := IntCode{
		input:     1,
		output:    0,
		phase:     -1,
		pointer:   0,
		memory:    memoryConstant,
		isStopped: false,
		doesRecur: true,
	}
	return intcode
}

func pad5(op int, instruction *[5]int) *[5]int {
	asString := fmt.Sprintf("%05d", op)
	asBytes := []byte(asString)
	for i := 0; i < 5; i++ {
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

func pass(candidate []int, commonMemory [678]int, instruction *[5]int) int {
	memA := commonMemory
	memB := commonMemory
	memC := commonMemory
	memD := commonMemory
	memE := commonMemory
	//for key, value := range commonMemory {
	//	memA[key] = value
	//	memB[key] = value
	//	memC[key] = value
	//	memD[key] = value
	//	memE[key] = value
	//}
	icpA := &IntCode{
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
		icReturn = opcode(icpA, instruction)
	}

	icpB := &IntCode{
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
		icReturn = opcode(icpB, instruction)
	}

	icpC := &IntCode{
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
		icReturn = opcode(icpC, instruction)
	}

	icpD := &IntCode{
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
		icReturn = opcode(icpD, instruction)
	}

	icpE := &IntCode{
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
		icReturn = opcode(icpE, instruction)
	}

	return icpE.output
}

func pass2(candidate []int, commonMemory [678]int, instruction *[5]int) int {
	memA := commonMemory
	memB := commonMemory
	memC := commonMemory
	memD := commonMemory
	memE := commonMemory
	//for key, value := range commonMemory {
	//	memA[key] = value
	//	memB[key] = value
	//	memC[key] = value
	//	memD[key] = value
	//	memE[key] = value
	//}
	eOutput := 0
	allStopped := false
	icpA := &IntCode{
		input:     0,
		output:    0,
		phase:     candidate[0],
		pointer:   0,
		memory:    memA,
		isStopped: false,
		doesRecur: false,
	}
	icpB := &IntCode{
		input:     0,
		output:    0,
		phase:     candidate[1],
		pointer:   0,
		memory:    memB,
		isStopped: false,
		doesRecur: false,
	}
	icpC := &IntCode{
		input:     0,
		output:    0,
		phase:     candidate[2],
		pointer:   0,
		memory:    memC,
		isStopped: false,
		doesRecur: false,
	}
	icpD := &IntCode{
		input:     0,
		output:    0,
		phase:     candidate[3],
		pointer:   0,
		memory:    memD,
		isStopped: false,
		doesRecur: false,
	}
	icpE := &IntCode{
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
			icReturn = opcode(icpA, instruction)
		}
		icpB.input = icpA.output
		icReturn = 1
		for icReturn == 1 {
			icReturn = opcode(icpB, instruction)
		}
		icpC.input = icpB.output
		icReturn = 1
		for icReturn == 1 {
			icReturn = opcode(icpC, instruction)
		}
		icpD.input = icpC.output
		icReturn = 1
		for icReturn == 1 {
			icReturn = opcode(icpD, instruction)
		}
		icpE.input = icpD.output
		icReturn = 1
		for icReturn == 1 {
			icReturn = opcode(icpE, instruction)
		}

		icpA.input = icpE.output
		eOutput = icpE.output
		allStopped = icpE.isStopped
	}

	return eOutput
}

func passes(candidates [][]int, memory [678]int, instruction *[5]int) []int {
	vcm := make([]int, len(candidates))
	for i, v := range candidates {
		vcm[i] = pass(v, memory, instruction)
	}
	return vcm
}

func passes2(candidates [][]int, memory [678]int, instruction *[5]int) []int {
	vcm := make([]int, len(candidates))
	for i, v := range candidates {
		vcm[i] = pass2(v, memory, instruction)
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

//func main() {
//	tv := intCodePkg.MakeMemory(fp)
//	A := []int{0, 1, 2, 3, 4}
//	permutations(len(A), A)
//	answer := passes(candidates, tv)
//	sort.Ints(answer)
//	fmt.Printf("Part A answer = %d. Correct = 368584\n", answer[len(answer)-1])
//
//	tv = intCodePkg.MakeMemory(fp)
//	candidates = nil
//	A = []int{5, 6, 7, 8, 9}
//	permutations(len(A), A)
//	answer2 := passes2(candidates, tv)
//	sort.Ints(answer2)
//	fmt.Printf("Part B answer = %d. Correct = 35993240\n", answer2[len(answer)-1])
//}

func main() {
	instruction := [5]int{}
	intcode := makeIntcode()
	var icReturn = 1

	for icReturn == 1 {
		icReturn = opcode(&intcode, &instruction)
	}

	fmt.Printf("\nPart A answer = %d. Correct = 9025675\n", intcode.output)

}
