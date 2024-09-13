package main

import (
	"fmt"
	"os"
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
	return text
}

func main() {
	fmt.Printf("%s\n", FileToString("advent02.csv"))
}
