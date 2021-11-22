package main

import (
	"fmt"
)

func pad5(op int) map[byte]int8 {
	asString := fmt.Sprintf("%05d", op)
	asChars := []byte(asString)
	values := [5]int8{}
	for i := 0; i < 5; i++ {
		values[i] = int8(asChars[i] - 48)
	}
	keys := [5]byte{'a', 'b', 'c', 'd', 'e'}
	elementMap := make(map[byte]int8)
	for i := 0; i < 5; i++ {
		elementMap[keys[i]] = values[i]
	}
	return elementMap
}

func main() {
	myMap := pad5(10000)
	for key, value := range myMap {
		fmt.Printf("%c value is %d\n", key, value)
	}
}
