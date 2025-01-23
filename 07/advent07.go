package main

import (
	"fmt"
	"sort"

	"advent-golang/intCodePkg"
)

const fp = "advent07.csv"

func areUnique(si []int) bool {
	m := map[int]bool{}
	for _, v := range si {
		if m[v] {
			return false
		} else {
			m[v] = true
		}
	}
	return true
}

func candidates() [][]int {
	var winners [][]int
	var candidate []int
	for a := 0; a < 5; a++ {
		for b := 0; b < 5; b++ {
			for c := 0; c < 5; c++ {
				for d := 0; d < 5; d++ {
					for e := 0; e < 5; e++ {
						candidate = nil
						candidate = append(candidate, a, b, c, d, e)
						if areUnique(candidate) {
							winners = append(winners, candidate)
						}
					}
				}
			}
		}
	}
	return winners
}

func candidates2() [][]int {
	var winners [][]int
	var candidate []int
	for a := 5; a < 10; a++ {
		for b := 5; b < 10; b++ {
			for c := 5; c < 10; c++ {
				for d := 5; d < 10; d++ {
					for e := 5; e < 10; e++ {
						candidate = nil
						candidate = append(candidate, a, b, c, d, e)
						if areUnique(candidate) {
							winners = append(winners, candidate)
						}
					}
				}
			}
		}
	}
	return winners
}

func pass(candidate []int, commonMemory map[int]int) int {
	memA := make(map[int]int, len(commonMemory))
	memB := make(map[int]int, len(commonMemory))
	memC := make(map[int]int, len(commonMemory))
	memD := make(map[int]int, len(commonMemory))
	memE := make(map[int]int, len(commonMemory))
	for key, value := range commonMemory {
		memA[key] = value
		memB[key] = value
		memC[key] = value
		memD[key] = value
		memE[key] = value
	}
	icpA := &intCodePkg.IntCode{
		Input:     0,
		Output:    []int{},
		Phase:     candidate[0],
		Pointer:   0,
		Memory:    memA,
		IsStopped: false,
		DoesRecur: true,
	}

	icReturn := 1
	for icReturn == 1 {
		icReturn = icpA.OpCode()
	}

	icpB := &intCodePkg.IntCode{
		Input:     icpA.Output[len(icpA.Output)-1],
		Output:    []int{},
		Phase:     candidate[1],
		Pointer:   0,
		Memory:    memB,
		IsStopped: false,
		DoesRecur: true,
	}

	icReturn = 1
	for icReturn == 1 {
		icReturn = icpB.OpCode()
	}

	icpC := &intCodePkg.IntCode{
		Input:     icpB.Output[len(icpB.Output)-1],
		Output:    []int{},
		Phase:     candidate[2],
		Pointer:   0,
		Memory:    memC,
		IsStopped: false,
		DoesRecur: true,
	}

	icReturn = 1
	for icReturn == 1 {
		icReturn = icpC.OpCode()
	}

	icpD := &intCodePkg.IntCode{
		Input:     icpC.Output[len(icpC.Output)-1],
		Output:    []int{},
		Phase:     candidate[3],
		Pointer:   0,
		Memory:    memD,
		IsStopped: false,
		DoesRecur: true,
	}

	icReturn = 1
	for icReturn == 1 {
		icReturn = icpD.OpCode()
	}

	icpE := &intCodePkg.IntCode{
		Input:     icpD.Output[len(icpD.Output)-1],
		Output:    []int{},
		Phase:     candidate[4],
		Pointer:   0,
		Memory:    memE,
		IsStopped: false,
		DoesRecur: true,
	}

	icReturn = 1
	for icReturn == 1 {
		icReturn = icpE.OpCode()
	}

	return icpE.Output[len(icpE.Output)-1]
}

func pass2(candidate []int, commonMemory map[int]int) int {
	memA := make(map[int]int, len(commonMemory))
	memB := make(map[int]int, len(commonMemory))
	memC := make(map[int]int, len(commonMemory))
	memD := make(map[int]int, len(commonMemory))
	memE := make(map[int]int, len(commonMemory))
	for key, value := range commonMemory {
		memA[key] = value
		memB[key] = value
		memC[key] = value
		memD[key] = value
		memE[key] = value
	}
	eOutput := 0
	allStopped := false
	icpA := &intCodePkg.IntCode{
		Input:     0,
		Output:    []int{},
		Phase:     candidate[0],
		Pointer:   0,
		Memory:    memA,
		IsStopped: false,
		DoesRecur: false,
	}
	icpB := &intCodePkg.IntCode{
		Input:     0,
		Output:    []int{},
		Phase:     candidate[1],
		Pointer:   0,
		Memory:    memB,
		IsStopped: false,
		DoesRecur: false,
	}
	icpC := &intCodePkg.IntCode{
		Input:     0,
		Output:    []int{},
		Phase:     candidate[2],
		Pointer:   0,
		Memory:    memC,
		IsStopped: false,
		DoesRecur: false,
	}
	icpD := &intCodePkg.IntCode{
		Input:     0,
		Output:    []int{},
		Phase:     candidate[3],
		Pointer:   0,
		Memory:    memD,
		IsStopped: false,
		DoesRecur: false,
	}
	icpE := &intCodePkg.IntCode{
		Input:     0,
		Output:    []int{},
		Phase:     candidate[4],
		Pointer:   0,
		Memory:    memE,
		IsStopped: false,
		DoesRecur: false,
	}

	for !allStopped {
		icReturn := 1
		for icReturn == 1 {
			icReturn = icpA.OpCode()
		}
		icpB.Input = icpA.Output[len(icpA.Output)-1]
		icReturn = 1
		for icReturn == 1 {
			icReturn = icpB.OpCode()
		}
		icpC.Input = icpB.Output[len(icpB.Output)-1]
		icReturn = 1
		for icReturn == 1 {
			icReturn = icpC.OpCode()
		}
		icpD.Input = icpC.Output[len(icpC.Output)-1]
		icReturn = 1
		for icReturn == 1 {
			icReturn = icpD.OpCode()
		}
		icpE.Input = icpD.Output[len(icpD.Output)-1]
		icReturn = 1
		for icReturn == 1 {
			icReturn = icpE.OpCode()
		}

		icpA.Input = icpE.Output[len(icpE.Output)-1]
		eOutput = icpE.Output[len(icpE.Output)-1]
		allStopped = icpE.IsStopped
	}

	return eOutput
}

func passes(candidates [][]int, memory map[int]int) []int {
	vcm := make([]int, len(candidates))
	for i, v := range candidates {
		vcm[i] = pass(v, memory)
	}
	return vcm
}

func passes2(candidates [][]int, memory map[int]int) []int {
	vcm := make([]int, len(candidates))
	for i, v := range candidates {
		vcm[i] = pass2(v, memory)
	}
	return vcm
}

func main() {
	tv := intCodePkg.MakeMemory(fp)
	answer := passes(candidates(), tv)
	sort.Ints(answer)
	fmt.Printf("Part A answer = %d\n", answer[len(answer)-1]) // Part A answer = 368584

	tv = intCodePkg.MakeMemory(fp)
	answer2 := passes2(candidates2(), tv)
	sort.Ints(answer2)
	fmt.Printf("Part B answer = %d\n", answer2[len(answer)-1]) // Part B answer = 35993240
}
