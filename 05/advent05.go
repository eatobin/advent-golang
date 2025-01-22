package main

import (
	"advent-golang/intcode"
	"advent-golang/makeMemory"
	"fmt"
)

const fp = "advent05.csv"

func main() {
	tv := makeMemory.MakeMemory(fp)
	icP := &intcode.IntCode{
		Input:        1,
		Output:       []int{},
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
	fmt.Printf("Part A answer = %d\n", icP.Output[len(icP.Output)-1]) // Part A answer = 9025675

	tv = makeMemory.MakeMemory(fp)
	icP = &intcode.IntCode{
		Input:        5,
		Output:       []int{},
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
	fmt.Printf("Part B answer = %d\n", icP.Output[len(icP.Output)-1]) // Part B answer = 11981754
}
