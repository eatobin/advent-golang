package main

import (
	"fmt"
	"github.com/eatobin/advent-golang/intcode"
)

const fp = "02/resources02/advent02.csv"

func updatedMemory(memory intcode.Memory, noun int, verb int) intcode.Memory {
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
			tv := intcode.MakeMemory(fp)
			icP := &intcode.IntCode{
				Pointer: 0,
				Memory:  updatedMemory(tv, noun, verb),
			}
			icReturn := 1
			for icReturn == 1 {
				icReturn = icP.OpCode()
			}
			candidate := icP.Memory[0]
			if candidate == 19690720 {
				break out
			}
		}
	}
	return (100 * noun) + verb
}

func main() {
	tv := intcode.MakeMemory(fp)
	icP := &intcode.IntCode{
		Pointer: 0,
		Memory:  updatedMemory(tv, 12, 2),
	}
	icReturn := 1
	for icReturn == 1 {
		icReturn = icP.OpCode()
	}
	fmt.Printf("Part A answer = %d\n", icP.Memory[0]) // Part A answer = 2890696
	fmt.Printf("Part B answer = %d", nounVerb())      // Part B answer = 8226
}
