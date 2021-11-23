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

func Pad5(op int) map[byte]byte {
	keys := [5]byte{'a', 'b', 'c', 'd', 'e'}
	instruction := make(map[byte]byte)
	asString := fmt.Sprintf("%05d", op)
	asBytes := []byte(asString)

	for i := 0; i < 5; i++ {
		instruction[keys[i]] = CharToInt(asBytes[i])
	}

	return instruction
}

func main() {
	myMap := Pad5(12345)
	for key, value := range myMap {
		fmt.Printf("%c :: %d\n", key, value)
	}
	//fmt.Printf("\nInt = %d", CharToInt('j'))
	fmt.Printf("\nInt = %d", CharToInt('0'))
	fmt.Printf("\nInt = %d", CharToInt('9'))
}
