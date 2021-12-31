package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Memory map[int]int
type Instruction map[byte]uint8

const offsetC int = 1
const offsetB int = 2
const offsetA int = 3

type IntCode struct {
	Input        int
	Output       int
	Phase        int
	Pointer      int
	RelativeBase int
	Memory       Memory
	IsStopped    bool
	DoesRecur    bool
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
		choice = icP.Memory[icP.Pointer+offsetA]
	case 2: // a-r-w
		choice = icP.Memory[icP.Pointer+offsetA] + icP.RelativeBase
	}
	return choice
}

func (icP *IntCode) bParam(instruction Instruction) int {
	var choice int
	switch instruction['b'] {
	case 0: // b-p-r
		choice = icP.Memory[icP.Memory[icP.Pointer+offsetB]]
	case 1: // b-i-r
		choice = icP.Memory[icP.Pointer+offsetB]
	case 2: // b-r-r
		choice = icP.Memory[icP.Memory[icP.Pointer+offsetB]+icP.RelativeBase]
	}
	return choice
}

func (icP *IntCode) cParam(instruction Instruction) int {
	var choice int
	if instruction['e'] == 3 {
		switch instruction['c'] {
		case 0: // c-p-w
			choice = icP.Memory[icP.Pointer+offsetC]
		case 2: // c-r-w
			choice = icP.Memory[icP.Pointer+offsetC] + icP.RelativeBase
		}
	} else {
		switch instruction['c'] {
		case 0: // c-p-r
			choice = icP.Memory[icP.Memory[icP.Pointer+offsetC]]
		case 1: // c-i-r
			choice = icP.Memory[icP.Pointer+offsetC]
		case 2: // c-r-r
			choice = icP.Memory[icP.Memory[icP.Pointer+offsetC]+icP.RelativeBase]
		}
	}
	return choice
}

func (icP *IntCode) OpCode() int {
	if icP.IsStopped {
		return 0
	} else {
		instruction := pad5(icP.Memory[icP.Pointer])
		if instruction['d'] == 9 {
			icP.IsStopped = true
			return 0
		} else {
			switch instruction['e'] {
			case 1:
				icP.Memory[icP.aParam(instruction)] = icP.bParam(instruction) + icP.cParam(instruction)
				icP.Pointer += 4
				return 1
			case 2:
				icP.Memory[icP.aParam(instruction)] = icP.bParam(instruction) * icP.cParam(instruction)
				icP.Pointer += 4
				return 1
			case 3:
				if icP.Phase != -1 {
					if icP.Pointer == 0 {
						icP.Memory[icP.cParam(instruction)] = icP.Phase
					} else {
						icP.Memory[icP.cParam(instruction)] = icP.Input
					}
				} else {
					icP.Memory[icP.cParam(instruction)] = icP.Input
				}
				icP.Pointer += 2
				return 1
			case 4:
				if icP.DoesRecur {
					icP.Output = icP.cParam(instruction)
					icP.Pointer += 2
					return 1
				} else {
					icP.Output = icP.cParam(instruction)
					icP.Pointer += 2
					return 0
				}
			case 5:
				if icP.cParam(instruction) == 0 {
					icP.Pointer += 3
				} else {
					icP.Pointer = icP.bParam(instruction)
				}
				return 1
			case 6:
				if icP.cParam(instruction) != 0 {
					icP.Pointer += 3
				} else {
					icP.Pointer = icP.bParam(instruction)
				}
				return 1
			case 7:
				if icP.cParam(instruction) < icP.bParam(instruction) {
					icP.Memory[icP.aParam(instruction)] = 1
				} else {
					icP.Memory[icP.aParam(instruction)] = 0
				}
				icP.Pointer += 4
				return 1
			case 8:
				if icP.cParam(instruction) == icP.bParam(instruction) {
					icP.Memory[icP.aParam(instruction)] = 1
				} else {
					icP.Memory[icP.aParam(instruction)] = 0
				}
				icP.Pointer += 4
				return 1
			case 9:
				icP.RelativeBase += icP.cParam(instruction)
				icP.Pointer += 2
				return 1
			default:
				panic("opcode is not valid")
			}
		}
	}
}
