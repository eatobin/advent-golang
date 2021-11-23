package main

import (
	"fmt"
)

func CharToInt(char byte) byte {
	if char < 48 || char > 57 {
		panic("Char is not an integer")
	}
	return char - 48
}

func pad5(op int) map[byte]byte {
	keys := [5]byte{'a', 'b', 'c', 'd', 'e'}
	values := [5]byte{}
	instruction := make(map[byte]byte)
	asString := fmt.Sprintf("%05d", op)
	asBytes := []byte(asString)

	for i := 0; i < 5; i++ {
		values[i] = CharToInt(asBytes[i])
		instruction[keys[i]] = values[i]
	}

	return instruction
}

func main() {
	myMap := pad5(10009)
	for key, value := range myMap {
		fmt.Printf("%c :: %d\n", key, value)
	}
}
