package main

import (
	"fmt"
)

func main() {
	keys := [4]byte{'a', 'b', 'c', 'd'}
	elements := [4]int8{55, 66, 77, 88}
	elementMap := make(map[byte]int8)
	for i := 0; i < 4; i++ {
		elementMap[keys[i]] = elements[i]
	}
	for key, value := range elementMap {
		fmt.Printf("%c value is %d\n", key, value)
	}
	fmt.Printf("\n%05d\n", 10)
	cc := fmt.Sprintf("%05d", 11)
	println(cc)
}
