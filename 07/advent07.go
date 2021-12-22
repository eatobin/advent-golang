package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Memory []int
type Instruction map[byte]uint8

const fp = "advent05.csv"
const offsetC int = 1
const offsetB int = 2
const offsetA int = 3

type IntCode struct {
	input     int
	output    int
	phase     int
	pointer   int
	memory    Memory
	isStopped bool
	doesRecur bool
}

// CompareIntCode compares two IntCodes
func CompareIntCode(a, b IntCode) bool {
	if &a == &b {
		return true
	}
	if a.input != b.input {
		return false
	}
	if a.output != b.output {
		return false
	}
	if a.phase != b.phase {
		return false
	}
	if a.pointer != b.pointer {
		return false
	}
	if len(a.memory) != len(b.memory) {
		return false
	}
	for i, v := range a.memory {
		if v != b.memory[i] {
			return false
		}
	}
	if a.isStopped != b.isStopped {
		return false
	}
	if a.doesRecur != b.doesRecur {
		return false
	}
	return true
}

func MakeMemory(fp string) Memory {
	dat, err := ioutil.ReadFile(fp)
	if err != nil {
		panic(err)
	}

	txt := string(dat)
	txt = strings.TrimRight(txt, "\n")
	strOps := strings.Split(txt, ",")
	memory := make([]int, len(strOps))

	for i, strOp := range strOps {
		op, err := strconv.Atoi(strOp)
		if err != nil {
			panic(err)
		}
		memory[i] = op
	}
	return memory
}

func charToInt(char byte) uint8 {
	if char < 48 || char > 57 {
		panic("Char is not an integer")
	}
	return char - 48
}

func pad5(op int) Instruction {
	keys := [5]byte{'a', 'b', 'c', 'd', 'e'}
	instruction := make(map[byte]uint8)
	asString := fmt.Sprintf("%05d", op)
	asBytes := []byte(asString)

	for i := 0; i < 5; i++ {
		instruction[keys[i]] = charToInt(asBytes[i])
	}
	return instruction
}

func getOrElse(pointer int, offsetX int, memory Memory) int {
	if (pointer + offsetX) > len(memory)-1 {
		return 0
	} else {
		return memory[memory[pointer+offsetX]]
	}
}

func (icP *IntCode) aParam(instruction Instruction) int {
	var choice int
	switch instruction['a'] {
	case 0: // a-p-w
		choice = icP.memory[icP.pointer+offsetA]
	}
	return choice
}

func (icP *IntCode) bParam(instruction Instruction) int {
	var choice int
	switch instruction['b'] {
	case 0: // b-p-r
		choice = getOrElse(icP.pointer, offsetB, icP.memory)
	case 1: // b-i-r
		choice = icP.memory[icP.pointer+offsetB]
	}
	return choice
}

func (icP *IntCode) cParam(instruction Instruction) int {
	var choice int
	if instruction['e'] == 3 {
		switch instruction['c'] {
		case 0: // c-p-w
			choice = icP.memory[icP.pointer+offsetC]
		}
	} else {
		switch instruction['c'] {
		case 0: // c-p-r
			choice = getOrElse(icP.pointer, offsetC, icP.memory)
		case 1: // c-i-r
			choice = icP.memory[icP.pointer+offsetC]
		}
	}
	return choice
}

func (icP *IntCode) opCode() int {
	if icP.isStopped {
		return 0
	} else {
		instruction := pad5(icP.memory[icP.pointer])
		if instruction['d'] == 9 {
			icP.isStopped = true
			return 1
		} else {
			switch instruction['e'] {
			case 1:
				icP.memory[icP.aParam(instruction)] = icP.bParam(instruction) + icP.cParam(instruction)
				icP.pointer += 4
				return 1
			case 2:
				icP.memory[icP.aParam(instruction)] = icP.bParam(instruction) * icP.cParam(instruction)
				icP.pointer += 4
				return 1
			case 3:
				if icP.phase != -1 {
					if icP.pointer == 0 {
						icP.memory[icP.cParam(instruction)] = icP.phase
					} else {
						icP.memory[icP.cParam(instruction)] = icP.input
					}
				} else {
					icP.memory[icP.cParam(instruction)] = icP.input
				}
				icP.pointer += 2
				return 1
			case 4:
				icP.output = icP.cParam(instruction)
				icP.pointer += 2
				return 1
			case 5:
				if icP.cParam(instruction) == 0 {
					icP.pointer += 3
				} else {
					icP.pointer = icP.bParam(instruction)
				}
				return 1
			case 6:
				if icP.cParam(instruction) != 0 {
					icP.pointer += 3
				} else {
					icP.pointer = icP.bParam(instruction)
				}
				return 1
			case 7:
				if icP.cParam(instruction) < icP.bParam(instruction) {
					icP.memory[icP.aParam(instruction)] = 1
				} else {
					icP.memory[icP.aParam(instruction)] = 0
				}
				icP.pointer += 4
				return 1
			case 8:
				if icP.cParam(instruction) == icP.bParam(instruction) {
					icP.memory[icP.aParam(instruction)] = 1
				} else {
					icP.memory[icP.aParam(instruction)] = 0
				}
				icP.pointer += 4
				return 1
			default:
				panic("opcode is not valid")
			}
		}
	}
}

func areUnique(si []int) bool {
	m := map[int]bool{}
	for _, v := range si {
		if m[v] {
			return false
		} else {
			m[v] = true
		}
	}
	return true
}

func candidates() [][]int {
	var winners [][]int
	var candidate []int
	for a := 0; a < 5; a++ {
		for b := 0; b < 5; b++ {
			for c := 0; c < 5; c++ {
				for d := 0; d < 5; d++ {
					for e := 0; e < 5; e++ {
						candidate = nil
						candidate = append(candidate, a, b, c, d, e)
						if areUnique(candidate) {
							winners = append(winners, candidate)
						}
					}
				}
			}
		}
	}
	return winners
}

func pass(candidate []int, memory Memory) int {
	icpA := &IntCode{
		input:     0,
		output:    0,
		phase:     candidate[0],
		pointer:   0,
		memory:    memory,
		isStopped: false,
		doesRecur: true,
	}

	icReturn := 1
	for icReturn == 1 {
		icReturn = icpA.opCode()
	}

	icpB := &IntCode{
		input:     icpA.output,
		output:    0,
		phase:     candidate[1],
		pointer:   0,
		memory:    memory,
		isStopped: false,
		doesRecur: true,
	}

	icReturn = 1
	for icReturn == 1 {
		icReturn = icpB.opCode()
	}

	icpC := &IntCode{
		input:     icpB.output,
		output:    0,
		phase:     candidate[2],
		pointer:   0,
		memory:    memory,
		isStopped: false,
		doesRecur: true,
	}

	icReturn = 1
	for icReturn == 1 {
		icReturn = icpC.opCode()
	}

	icpD := &IntCode{
		input:     icpC.output,
		output:    0,
		phase:     candidate[3],
		pointer:   0,
		memory:    memory,
		isStopped: false,
		doesRecur: true,
	}

	icReturn = 1
	for icReturn == 1 {
		icReturn = icpD.opCode()
	}

	icpE := &IntCode{
		input:     icpD.output,
		output:    0,
		phase:     candidate[4],
		pointer:   0,
		memory:    memory,
		isStopped: false,
		doesRecur: true,
	}

	icReturn = 1
	for icReturn == 1 {
		icReturn = icpE.opCode()
	}

	return icpE.output
}

func main() {
	var mem = []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}
	var phases = []int{1, 0, 4, 3, 2}
	answer := pass(phases, mem)
	fmt.Printf("%v", answer)

	//fmt.Printf("%v", candidates())
	//fmt.Printf("\ncount = %d", len(candidates()))
	//fmt.Printf("\n%v", candidates()[119])
	//a := makeRange(0, 4)
	//fmt.Println(a)
	//visited := []int{
	//	1,
	//	2,
	//	88,
	//	22,
	//}
	//
	//fmt.Println(areUnique(visited))
}

//tv := MakeMemory(fp)
//icP := &IntCode{
//	input:     1,
//	output:    0,
//	phase:     -1,
//	pointer:   0,
//	memory:    tv,
//	isStopped: false,
//	doesRecur: true,
//}
//icReturn := 1
//for icReturn == 1 {
//	icReturn = icP.opCode()
//}
//fmt.Printf("Part A answer = %d\n", icP.output) // Part A answer = ;368584
//
//tv = MakeMemory(fp)
//icP = &IntCode{
//	input:     5,
//	output:    0,
//	phase:     -1,
//	pointer:   0,
//	memory:    tv,
//	isStopped: false,
//	doesRecur: true,
//}
//icReturn = 1
//for icReturn == 1 {
//	icReturn = icP.opCode()
//}
//fmt.Printf("Part B answer = %d", icP.output) // Part B answer = 11981754
//}
