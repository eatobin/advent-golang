package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

//fp = "advent02.csv"

func MakeMemory(fp string) map[int]int {
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

func CharToInt(char byte) uint8 {
	if char < 48 || char > 57 {
		panic("Char is not an integer")
	}
	return char - 48
}

func Pad5(op int) map[byte]uint8 {
	keys := [5]byte{'a', 'b', 'c', 'd', 'e'}
	instruction := make(map[byte]uint8)
	asString := fmt.Sprintf("%05d", op)
	asBytes := []byte(asString)

	for i := 0; i < 5; i++ {
		instruction[keys[i]] = CharToInt(asBytes[i])
	}
	return instruction
}

func main() {
	for key, value := range MakeMemory("advent02.csv") {
		fmt.Printf("%3d :: %d\n", key, value)
	}
	fmt.Println(MakeMemory("advent02.csv")[120])

	myMap := Pad5(12345)
	for key, value := range myMap {
		fmt.Printf("%c :: %d\n", key, value)
	}
	//fmt.Printf("\nInt = %d", CharToInt('j'))
	fmt.Printf("\nInt = %d", CharToInt('0'))
	fmt.Printf("\nInt = %d", CharToInt('9'))
}
