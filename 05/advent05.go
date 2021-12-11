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
	input   int
	output  int
	pointer int
	memory  Memory
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

func aParam(instruction Instruction, ic IntCode) int {
	var choice int
	switch instruction['a'] {
	case 0: // a-p-w
		choice = ic.memory[ic.pointer+offsetA]
	}
	return choice
}

func bParam(instruction Instruction, ic IntCode) int {
	var choice int
	switch instruction['b'] {
	case 0: // b-p-r
		choice = getOrElse(ic.pointer, offsetB, ic.memory)
	case 1: // b-i-r
		choice = ic.memory[ic.pointer+offsetB]
	}
	return choice
}

func cParam(instruction Instruction, ic IntCode) int {
	var choice int
	switch instruction['e'] {
	case 3:
		switch instruction['c'] {
		case 0: // c-p-w
			choice = ic.memory[ic.pointer+offsetC]
		}
	default:
		switch instruction['c'] {
		case 0: // c-p-r
			choice = getOrElse(ic.pointer, offsetC, ic.memory)
		case 1: // c-i-r
			choice = ic.memory[ic.pointer+offsetC]
		}
	}
	return choice
}

func updateMemory(memory Memory, index int, value int) Memory {
	memory[index] = value
	return memory
}

func OpCode(ic IntCode) IntCode {
	instruction := pad5(ic.memory[ic.pointer])
	switch instruction['e'] {
	case 9:
		return ic
	case 1:
		return OpCode(IntCode{
			input:   ic.input,
			output:  ic.output,
			pointer: ic.pointer + 4,
			memory: updateMemory(ic.memory,
				aParam(instruction, ic),
				bParam(instruction, ic)+cParam(instruction, ic)),
		})
	case 2:
		return OpCode(IntCode{
			input:   ic.input,
			output:  ic.output,
			pointer: ic.pointer + 4,
			memory: updateMemory(ic.memory,
				aParam(instruction, ic),
				bParam(instruction, ic)*cParam(instruction, ic)),
		})
	case 3:
		return OpCode(IntCode{
			input:   ic.input,
			output:  ic.output,
			pointer: ic.pointer + 2,
			memory: updateMemory(ic.memory,
				cParam(instruction, ic),
				ic.input),
		})
	case 4:
		return OpCode(IntCode{
			input:   ic.input,
			output:  cParam(instruction, ic),
			pointer: ic.pointer + 2,
			memory:  ic.memory,
		})
	case 5:
		var newPointer = ic.pointer
		c := cParam(instruction, ic)
		if c == 0 {
			newPointer = newPointer + 3
		} else {
			newPointer = bParam(instruction, ic)
		}
		return OpCode(IntCode{
			input:   ic.input,
			output:  ic.output,
			pointer: newPointer,
			memory:  ic.memory,
		})
	case 6:
		var newPointer = ic.pointer
		c := cParam(instruction, ic)
		if c != 0 {
			newPointer = newPointer + 3
		} else {
			newPointer = bParam(instruction, ic)
		}
		return OpCode(IntCode{
			input:   ic.input,
			output:  ic.output,
			pointer: newPointer,
			memory:  ic.memory,
		})
	case 7:
		var newMemory = ic.memory
		c := cParam(instruction, ic)
		b := bParam(instruction, ic)
		if c < b {
			newMemory = updateMemory(newMemory,
				aParam(instruction, ic),
				1)
		} else {
			newMemory = updateMemory(newMemory,
				aParam(instruction, ic),
				0)
		}
		return OpCode(IntCode{
			input:   ic.input,
			output:  ic.output,
			pointer: ic.pointer + 4,
			memory:  newMemory,
		})
	case 8:
		var newMemory = ic.memory
		c := cParam(instruction, ic)
		b := bParam(instruction, ic)
		if c == b {
			newMemory = updateMemory(newMemory,
				aParam(instruction, ic),
				1)
		} else {
			newMemory = updateMemory(newMemory,
				aParam(instruction, ic),
				0)
		}
		return OpCode(IntCode{
			input:   ic.input,
			output:  ic.output,
			pointer: ic.pointer + 4,
			memory:  newMemory,
		})
	default:
		panic("opcode is not valid")
	}
}

func main() {
	tv := MakeMemory(fp)
	answer := OpCode(IntCode{input: 1, output: 0, pointer: 0, memory: tv})
	fmt.Printf("Part A answer = %d\n", answer.output) // Part A answer = 9025675

	tv = MakeMemory(fp)
	answer2 := OpCode(IntCode{input: 5, output: 0, pointer: 0, memory: tv})
	fmt.Printf("Part B answer = %d", answer2.output) // Part B answer = 11981754
}
