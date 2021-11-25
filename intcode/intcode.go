package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Memory map[int]int
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

func aParam(instruction Instruction, pointer int, memory Memory) int {
	switch instruction['a'] {
	// a-p-w
	case 0:
		return memory[pointer+offsetA]
	default:
		return 0
	}
}

func bParam(instruction Instruction, pointer int, memory Memory) int {
	switch instruction['b'] {
	// b-p-r
	case 0:
		return memory[memory[pointer+offsetB]]
	default:
		return 0
	}
}

func cParam(instruction Instruction, pointer int, memory Memory) int {
	switch instruction['c'] {
	// c-p-r
	case 0:
		return memory[memory[pointer+offsetC]]
	default:
		return 0
	}
}

func opCode(ic IntCode) IntCode {
	instruction := pad5(ic.memory[ic.pointer])
	switch instruction['e'] {
	case 1:
		a := aParam(instruction, ic.pointer, ic.memory)
		b := bParam(instruction, ic.pointer, ic.memory)
		c := cParam(instruction, ic.pointer, ic.memory)
		opCode(IntCode{
			pointer: ic.pointer + 4,
			//memory:  ic.memory[aParam(instruction, ic.pointer, ic.memory)],
			memory: ic.memory,
		})
	case 9:
		return IntCode{pointer: ic.pointer, memory: ic.memory}
	}
	return IntCode{}
}

func main() {
	for key, value := range MakeMemory(fp) {
		fmt.Printf("%3d :: %d\n", key, value)
	}
	fmt.Println(MakeMemory("advent02.csv")[120])

	myMap := pad5(12345)
	for key, value := range myMap {
		fmt.Printf("%c :: %d\n", key, value)
	}
	//fmt.Printf("\nInt = %d", charToInt('j'))
	fmt.Printf("\nInt = %d", charToInt('0'))
	fmt.Printf("\nInt = %d", charToInt('9'))
}
