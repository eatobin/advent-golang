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
