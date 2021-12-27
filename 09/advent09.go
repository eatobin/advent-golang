package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Memory []int
type Instruction map[byte]uint8

const fp = "advent09.csv"
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
			return 0
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
				if icP.doesRecur {
					icP.output = icP.cParam(instruction)
					icP.pointer += 2
					return 1
				} else {
					icP.output = icP.cParam(instruction)
					icP.pointer += 2
					return 0
				}
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

func candidates2() [][]int {
	var winners [][]int
	var candidate []int
	for a := 5; a < 10; a++ {
		for b := 5; b < 10; b++ {
			for c := 5; c < 10; c++ {
				for d := 5; d < 10; d++ {
					for e := 5; e < 10; e++ {
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

func pass(candidate []int, commonMemory Memory) int {
	memA := make([]int, len(commonMemory))
	memB := make([]int, len(commonMemory))
	memC := make([]int, len(commonMemory))
	memD := make([]int, len(commonMemory))
	memE := make([]int, len(commonMemory))
	copy(memA, commonMemory)
	copy(memB, commonMemory)
	copy(memC, commonMemory)
	copy(memD, commonMemory)
	copy(memE, commonMemory)
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
		icReturn = icpA.opCode()
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
		icReturn = icpB.opCode()
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
		icReturn = icpC.opCode()
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
		icReturn = icpD.opCode()
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
		icReturn = icpE.opCode()
	}

	return icpE.output
}

func pass2(candidate []int, commonMemory Memory) int {
	memA := make([]int, len(commonMemory))
	memB := make([]int, len(commonMemory))
	memC := make([]int, len(commonMemory))
	memD := make([]int, len(commonMemory))
	memE := make([]int, len(commonMemory))
	copy(memA, commonMemory)
	copy(memB, commonMemory)
	copy(memC, commonMemory)
	copy(memD, commonMemory)
	copy(memE, commonMemory)
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
			icReturn = icpA.opCode()
		}
		icpB.input = icpA.output
		icReturn = 1
		for icReturn == 1 {
			icReturn = icpB.opCode()
		}
		icpC.input = icpB.output
		icReturn = 1
		for icReturn == 1 {
			icReturn = icpC.opCode()
		}
		icpD.input = icpC.output
		icReturn = 1
		for icReturn == 1 {
			icReturn = icpD.opCode()
		}
		icpE.input = icpD.output
		icReturn = 1
		for icReturn == 1 {
			icReturn = icpE.opCode()
		}

		icpA.input = icpE.output
		eOutput = icpE.output
		allStopped = icpE.isStopped
	}

	return eOutput
}

func passes(candidates [][]int, memory Memory) []int {
	vcm := make([]int, len(candidates))
	for i, v := range candidates {
		vcm[i] = pass(v, memory)
	}
	return vcm
}

func passes2(candidates [][]int, memory Memory) []int {
	vcm := make([]int, len(candidates))
	for i, v := range candidates {
		vcm[i] = pass2(v, memory)
	}
	return vcm
}

func main() {
	tv := MakeMemory(fp)
	answer := passes(candidates(), tv)
	sort.Ints(answer)
	fmt.Printf("Part A answer = %d\n", answer[len(answer)-1]) // Part A answer = ;368584

	tv = MakeMemory(fp)
	answer2 := passes2(candidates2(), tv)
	sort.Ints(answer2)
	fmt.Printf("Part B answer = %d", answer2[len(answer)-1]) // Part B answer = 35993240
}
