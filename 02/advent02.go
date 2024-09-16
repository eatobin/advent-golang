package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Memory []int

type Intcode struct {
	Pointer int
	Memory  Memory
}

func MakeMemory(fp string) Memory {
	dat, err := os.ReadFile(fp)
	if err != nil {
		panic(err)
	}

	txt := string(dat)
	txt = strings.TrimRight(txt, "\n")
	strOps := strings.Split(txt, ",")
	length := len(strOps)
	memory := make([]int, length)

	for i, strOp := range strOps {
		op, err := strconv.Atoi(strOp)
		if err != nil {
			panic(err)
		}
		memory[i] = op
	}
	return memory
}

func Opcode(intcode *Intcode) int {
	var action int
	var address1 int
	var address2 int
	var address3 int

	action = intcode.Memory[intcode.Pointer]
	address1 = intcode.Memory[intcode.Pointer+1]
	address2 = intcode.Memory[intcode.Pointer+2]
	address3 = intcode.Memory[intcode.Pointer+3]

	switch action {
	case 1:
		intcode.Memory[address3] =
			intcode.Memory[address1] +
				intcode.Memory[address2]
		intcode.Pointer += 4
		action = intcode.Memory[intcode.Pointer]
		return 1
	case 2:
		intcode.Memory[address3] =
			intcode.Memory[address1] *
				intcode.Memory[address2]
		intcode.Pointer += 4
		action = intcode.Memory[intcode.Pointer]
		return 1
	}
	return 0
}

func UpdatedMemory(intcode *Intcode, noun int, verb int) {
	intcode.Memory[1] = noun
	intcode.Memory[2] = verb
}

func nounVerb() int {
	var noun int
	var verb int
	var intcode Intcode
	var icReturn int
	var candidate int

	for noun = 0; noun < 100; noun++ {
		for verb = 0; verb < 100; verb++ {
			intcode.Pointer = 0
			intcode.Memory = MakeMemory("advent02.csv")
			UpdatedMemory(&intcode, noun, verb)

			icReturn = 1
			for icReturn == 1 {
				icReturn = Opcode(&intcode)
			}

			candidate = intcode.Memory[0]
			if candidate == 19690720 {
				goto end
			}
		}
	}
end:
	return (100 * noun) + verb
}

func main() {
	var intcode Intcode
	var icReturn int

	intcode.Pointer = 0
	intcode.Memory = MakeMemory("advent02.csv")
	icReturn = 1

	UpdatedMemory(&intcode, 12, 2)

	for icReturn == 1 {
		icReturn = Opcode(&intcode)
	}

	fmt.Printf("Memory length: %d\n\n", len(intcode.Memory))
	fmt.Printf("Part A answer = %d\n", intcode.Memory[0]) // Part A answer = 2890696
	fmt.Printf("Part B answer = %d\n", nounVerb())        // Part B answer = 8226
}
