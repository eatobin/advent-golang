package main

import (
	"github.com/eatobin/advent-golang/intcode"
	"testing"
)

type Fixtures struct {
	Value    intcode.IntCode
	Expected intcode.IntCode
}

// CompareIntCode compares two IntCodes
func CompareIntCode(a, b intcode.IntCode) bool {
	if &a == &b {
		return true
	}
	if a.Pointer != b.Pointer {
		return false
	}
	if len(a.Memory) != len(b.Memory) {
		return false
	}
	for i, v := range a.Memory {
		if v != b.Memory[i] {
			return false
		}
	}
	return true
}

func TestOpCode(t *testing.T) {
	test01In := intcode.MakeMemory("resources02/test01In.csv")
	test01Out := intcode.MakeMemory("resources02/test01Out.csv")
	test02In := intcode.MakeMemory("resources02/test02In.csv")
	test02Out := intcode.MakeMemory("resources02/test02Out.csv")
	test03In := intcode.MakeMemory("resources02/test03In.csv")
	test03Out := intcode.MakeMemory("resources02/test03Out.csv")
	test04In := intcode.MakeMemory("resources02/test04In.csv")
	test04Out := intcode.MakeMemory("resources02/test04Out.csv")
	test05In := intcode.MakeMemory("resources02/test05In.csv")
	test05Out := intcode.MakeMemory("resources02/test05Out.csv")

	fixtures := []Fixtures{
		{intcode.IntCode{Pointer: 0, Memory: test01In}, intcode.IntCode{Pointer: 4, Memory: test01Out}},
		{intcode.IntCode{Pointer: 0, Memory: test02In}, intcode.IntCode{Pointer: 4, Memory: test02Out}},
		{intcode.IntCode{Pointer: 0, Memory: test03In}, intcode.IntCode{Pointer: 8, Memory: test03Out}},
		{intcode.IntCode{Pointer: 0, Memory: test04In}, intcode.IntCode{Pointer: 4, Memory: test04Out}},
		{intcode.IntCode{Pointer: 0, Memory: test05In}, intcode.IntCode{Pointer: 8, Memory: test05Out}},
	}

	for _, fixture := range fixtures {
		icReturn := 1
		for icReturn == 1 {
			icReturn = fixture.Value.OpCode()
		}
		if !CompareIntCode(fixture.Value, fixture.Expected) {
			t.Errorf("Got %v; want\n %v", fixture.Value, fixture.Expected)
		}
	}
}
