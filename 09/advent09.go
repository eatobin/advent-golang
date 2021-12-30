package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Memory map[int]int
type Instruction map[byte]uint8

const fp = "advent09.csv"
const offsetC int = 1
const offsetB int = 2
const offsetA int = 3

type IntCode struct {
	input        int
	output       int
	phase        int
	pointer      int
	relativeBase int
	memory       Memory
	isStopped    bool
	doesRecur    bool
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
	if a.relativeBase != b.relativeBase {
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
	memory := make(map[int]int)

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

func (icP *IntCode) aParam(instruction Instruction) int {
	var choice int
	switch instruction['a'] {
	case 0: // a-p-w
		choice = icP.memory[icP.pointer+offsetA]
	case 2: // a-r-w
		choice = icP.memory[icP.pointer+offsetA] + icP.relativeBase
	}
	return choice
}

func (icP *IntCode) bParam(instruction Instruction) int {
	var choice int
	switch instruction['b'] {
	case 0: // b-p-r
		choice = icP.memory[icP.memory[icP.pointer+offsetB]]
	case 1: // b-i-r
		choice = icP.memory[icP.pointer+offsetB]
	case 2: // b-r-r
		choice = icP.memory[icP.memory[icP.pointer+offsetB]+icP.relativeBase]
	}
	return choice
}

func (icP *IntCode) cParam(instruction Instruction) int {
	var choice int
	if instruction['e'] == 3 {
		switch instruction['c'] {
		case 0: // c-p-w
			choice = icP.memory[icP.pointer+offsetC]
		case 2: // c-r-w
			choice = icP.memory[icP.pointer+offsetC] + icP.relativeBase
		}
	} else {
		switch instruction['c'] {
		case 0: // c-p-r
			choice = icP.memory[icP.memory[icP.pointer+offsetC]]
		case 1: // c-i-r
			choice = icP.memory[icP.pointer+offsetC]
		case 2: // c-r-r
			choice = icP.memory[icP.memory[icP.pointer+offsetC]+icP.relativeBase]
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
			case 9:
				icP.relativeBase += icP.memory[icP.cParam(instruction)]
				icP.pointer += 2
				return 1
			default:
				panic("opcode is not valid")
			}
		}
	}
}

func main() {
	tv := MakeMemory(fp)
	icP := &IntCode{
		input:        1,
		output:       0,
		phase:        -1,
		pointer:      0,
		relativeBase: 0,
		memory:       tv,
		isStopped:    false,
		doesRecur:    true,
	}
	icReturn := 1
	for icReturn == 1 {
		icReturn = icP.opCode()
	}
	fmt.Printf("Part A answer = %v\n", *icP) // Part A answer = 3780860499

	//tv = MakeMemory(fp)
	//fmt.Printf("Part B answer = %d", answer2[len(answer)-1]) // Part B answer = 35993240
}
