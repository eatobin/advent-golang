package main

import (
	"fmt"
	"sort"

	"advent-golang/intCodePkg"
)

const fp = "advent07.csv"

var candidates [][]int

func addAPerm(perm []int) {
	tmp := make([]int, len(perm))
	copy(tmp, perm)
	candidates = append(candidates, tmp)
}

func permutations(k int, A []int) {
	if k == 1 {
		addAPerm(A)
	} else {
		permutations(k-1, A)
		for i := 0; i < k-1; i++ {
			if k%2 == 0 {
				A[i], A[k-1] = A[k-1], A[i]
			} else {
				A[0], A[k-1] = A[k-1], A[0]
			}
			permutations(k-1, A)
		}
	}
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
	A := []int{0, 1, 2, 3, 4}
	permutations(len(A), A)
	answer := passes(candidates, tv)
	sort.Ints(answer)
	fmt.Printf("Part A answer = %d. Correct = 368584\n", answer[len(answer)-1])

	tv = intCodePkg.MakeMemory(fp)
	candidates = nil
	A = []int{5, 6, 7, 8, 9}
	permutations(len(A), A)
	answer2 := passes2(candidates, tv)
	sort.Ints(answer2)
	fmt.Printf("Part B answer = %d. Correct = 35993240\n", answer2[len(answer)-1])
}
