package main

import (
	"fmt"
	"os"
	"strings"
)

type Memory struct {
	Contents []int
	Length   int
}

type Intcode struct {
	Pointer int
	Memory  Memory
}

func FileToString(filename string) string {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	text := string(fileContent)
	text = strings.TrimRight(text, "\n")
	return text
}

func ReturnMemoryLength(string string) int {
	count := 0

	for i := 0; i < len(string); i++ {
		if string[i] == ',' {
			count++
		}
	}
	return count + 1
}

func main() {
	myString := FileToString("advent02.csv")
	length := ReturnMemoryLength(myString)
	fmt.Printf("%s\n", myString)
	fmt.Printf("%d\n", length)
}
