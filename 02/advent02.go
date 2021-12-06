package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Memory []int
type Instruction map[byte]uint8

const fp = "advent02.csv"
const offsetC int = 1
const offsetB int = 2
const offsetA int = 3

type IntCode struct {
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
	// a-p-w
	case 0:
		choice = ic.memory[ic.pointer+offsetA]
	}
	return choice
}

func bParam(instruction Instruction, ic IntCode) int {
	var choice int
	switch instruction['b'] {
	// b-p-r
	case 0:
		choice = getOrElse(ic.pointer, offsetB, ic.memory)
	}
	return choice
}

func cParam(instruction Instruction, ic IntCode) int {
	var choice int
	switch instruction['c'] {
	// c-p-r
	case 0:
		choice = getOrElse(ic.pointer, offsetC, ic.memory)
	}
	return choice
}

func updateMemory(memory Memory, index int, value int) Memory {
	memory[index] = value
	return memory
}

func opCode(ic IntCode) IntCode {
	instruction := pad5(ic.memory[ic.pointer])
	switch instruction['e'] {
	case 9:
		break
	case 1:
		return opCode(IntCode{
			pointer: ic.pointer + 4,
			memory: updateMemory(ic.memory,
				aParam(instruction, ic),
				bParam(instruction, ic)+cParam(instruction, ic)),
		})
	case 2:
		return opCode(IntCode{
			pointer: ic.pointer + 4,
			memory: updateMemory(ic.memory,
				aParam(instruction, ic),
				bParam(instruction, ic)*cParam(instruction, ic)),
		})
	default:
		panic("opcode is not valid")
	}
	return ic
}

func updatedMemory(memory Memory, noun int, verb int) Memory {
	memory[1] = noun
	memory[2] = verb
	return memory
}

func nounVerb() int {
	var noun int
	var verb int

out:
	for noun = 0; noun < 101; noun++ {
		for verb = 0; verb < 101; verb++ {
			tv := MakeMemory(fp)
			candidate := opCode(IntCode{pointer: 0, memory: updatedMemory(tv, noun, verb)}).memory[0]
			if candidate == 19690720 {
				break out
			}
		}
	}
	return (100 * noun) + verb
}

func main() {
	tv := MakeMemory(fp)
	answer := opCode(IntCode{pointer: 0, memory: updatedMemory(tv, 12, 2)})
	fmt.Printf("Part A answer = %d\n", answer.memory[0]) // Part A answer = 2890696
	fmt.Printf("Part B answer = %d", nounVerb())         // Part B answer = 8226
}
