package main

import (
	"fmt"
	"github.com/eatobin/advent-golang/intcode"
)

const fp = "09/advent09.csv"

func main() {
	tv := intcode.MakeMemory(fp)
	icP := &intcode.IntCode{
		Input:        1,
		Output:       0,
		Phase:        -1,
		Pointer:      0,
		RelativeBase: 0,
		Memory:       tv,
		IsStopped:    false,
		DoesRecur:    true,
	}
	icReturn := 1
	for icReturn == 1 {
		icReturn = icP.OpCode()
	}
	fmt.Printf("Part A answer = %d\n", icP.Output) // Part A answer = 3780860499

	tv = intcode.MakeMemory(fp)
	icP = &intcode.IntCode{
		Input:        2,
		Output:       0,
		Phase:        -1,
		Pointer:      0,
		RelativeBase: 0,
		Memory:       tv,
		IsStopped:    false,
		DoesRecur:    true,
	}
	icReturn = 1
	for icReturn == 1 {
		icReturn = icP.OpCode()
	}
	fmt.Printf("Part B answer = %d\n", icP.Output) // Part A answer = 33343
}
