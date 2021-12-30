package main

import (
	"testing"
)

type Fixtures struct {
	Value    IntCode
	Expected IntCode
}

// CompareIntCode compares two IntCodes
func CompareIntCode(a, b IntCode) bool {
	if &a == &b {
		return true
	}
	if a.input != b.input {
		return false
	}
	if a.output != b.output {
		return false
	}
	if a.phase != b.phase {
		return false
	}
	if a.pointer != b.pointer {
		return false
	}
	if a.relativeBase != b.relativeBase {
		return false
	}
	if len(a.memory) != len(b.memory) {
		return false
	}
	for i, v := range a.memory {
		if v != b.memory[i] {
			return false
		}
	}
	if a.isStopped != b.isStopped {
		return false
	}
	if a.doesRecur != b.doesRecur {
		return false
	}
	return true
}

func TestOpCode(t *testing.T) {
	test00In := MakeMemory("test00In.csv")
	test00Out := MakeMemory("test00Out.csv")
	test01In := MakeMemory("test01In.csv")
	test01Out := MakeMemory("test01Out.csv")
	test02In := MakeMemory("test02In.csv")
	test02Out := MakeMemory("test02Out.csv")

	fixtures := []Fixtures{
		{IntCode{memory: test00In}, IntCode{pointer: 4, output: 109, relativeBase: 1, memory: test00Out}},
		{IntCode{memory: test01In}, IntCode{pointer: 6, output: 1219070632396864, memory: test01Out}},
		{IntCode{memory: test02In}, IntCode{output: 1125899906842624, pointer: 2, memory: test02Out}},
	}

	for _, fixture := range fixtures {
		icReturn := 1
		for icReturn == 1 {
			icReturn = fixture.Value.opCode()
		}
		if !CompareIntCode(fixture.Value, fixture.Expected) {
			t.Errorf("Got %v; want\n %v", fixture.Value, fixture.Expected)
		}
	}
}
