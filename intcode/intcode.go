package main

import (
	"fmt"
)

func main() {
	keys := [4]byte{'a', 'b', 'c', 'd'}
	elements := [4]byte{55, 66, 77, 88}
	elementMap := make(map[byte]byte)
	for i := 0; i < 4; i++ {
		elementMap[keys[i]] = elements[i]
	}
	for key, value := range elementMap {
		fmt.Printf("%c value is %d\n", key, value)
	}
}
