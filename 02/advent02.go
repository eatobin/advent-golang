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
const offsetC = 1
const offsetB = 2
const offsetA = 3

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

func aParam(instruction Instruction, pointer int, memory Memory) int {
	var choice int
	switch instruction['a'] {
	// a-p-w
	case 0:
		choice = memory[pointer+offsetA]
	}
	return choice
}

func bParam(instruction Instruction, pointer int, memory Memory) int {
	var choice int
	switch instruction['b'] {
	// b-p-r
	case 0:
		choice = memory[memory[pointer+offsetB]]
	}
	return choice
}

func cParam(instruction Instruction, pointer int, memory Memory) int {
	var choice int
	switch instruction['c'] {
	// c-p-r
	case 0:
		choice = memory[memory[pointer+offsetC]]
	}
	return choice
}

func updateMemory(memory Memory, key int, value int) Memory {
	memory[key] = value
	return memory
}

func opCode(ic IntCode) IntCode {
	instruction := pad5(ic.memory[ic.pointer])
	switch instruction['e'] {
	case 9:
		break
	case 1:
		opCode(IntCode{
			pointer: ic.pointer + 4,
			memory: updateMemory(ic.memory, aParam(instruction, ic.pointer, ic.memory),
				bParam(instruction, ic.pointer, ic.memory)+cParam(instruction, ic.pointer, ic.memory)),
		})
	case 2:
		opCode(IntCode{
			pointer: ic.pointer + 4,
			memory: updateMemory(ic.memory, aParam(instruction, ic.pointer, ic.memory),
				bParam(instruction, ic.pointer, ic.memory)*cParam(instruction, ic.pointer, ic.memory)),
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

func nounVerb() {
	for noun := 0; noun < 101; noun++ {
		for verb := 0; verb < 101; verb++ {
			tv := MakeMemory(fp)
			candidate := opCode(IntCode{pointer: 0, memory: updatedMemory(tv, noun, verb)}).memory[0]
			if candidate == 19690720 {
				fmt.Printf("Part B answer = %d", (100*noun)+verb)
			}
		}
	}
}

func main() {
	tv := MakeMemory(fp)

	answer := opCode(IntCode{pointer: 0, memory: updatedMemory(tv, 12, 2)})
	fmt.Printf("Part A answer = %d\n", answer.memory[0]) // Part A answer = 2890696
	nounVerb()                                           // Part B answer = 8226
}
