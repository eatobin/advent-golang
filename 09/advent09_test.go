package main

import (
	"testing"
)

type Fixtures struct {
	Value    IntCode
	Expected IntCode
}

func TestOpCode(t *testing.T) {
	test01In := MakeMemory("test01In.csv")
	test01Out := MakeMemory("test01Out.csv")
	test02In := MakeMemory("test02In.csv")
	test02Out := MakeMemory("test02Out.csv")
	test03In := MakeMemory("test03In.csv")
	test03Out := MakeMemory("test03Out.csv")
	test04In := MakeMemory("test04In.csv")
	test04Out := MakeMemory("test04Out.csv")
	test05In := MakeMemory("test05In.csv")
	test05Out := MakeMemory("test05Out.csv")

	fixtures := []Fixtures{
		{IntCode{pointer: 0, memory: test01In}, IntCode{pointer: 4, memory: test01Out}},
		{IntCode{pointer: 0, memory: test02In}, IntCode{pointer: 4, memory: test02Out}},
		{IntCode{pointer: 0, memory: test03In}, IntCode{pointer: 8, memory: test03Out}},
		{IntCode{pointer: 0, memory: test04In}, IntCode{pointer: 4, memory: test04Out}},
		{IntCode{pointer: 0, memory: test05In}, IntCode{pointer: 8, memory: test05Out}},
	}

	for _, fixture := range fixtures {
		icReturn := 1
		for icReturn == 1 {
			icReturn = OpCode(&fixture.Value)
		}
		if !CompareIntCode(fixture.Value, fixture.Expected) {
			t.Errorf("Got %v; want\n %v", fixture.Value, fixture.Expected)
		}
	}
}
