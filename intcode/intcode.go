package main

import (
	"fmt"
)

func main() {
	elements := [4]string{"abc", "def", "fgi", "adi"}
	elementMap := make(map[int]string)
	for i := 0; i < 4; i++ {
		elementMap[i] = elements[i]
	}
	fmt.Println(elementMap)
}
