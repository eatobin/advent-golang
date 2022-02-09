package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Directions = []string

const fp = "03/day03a.csv"

func MakeDirections(fp string) Directions {
	dat, err := ioutil.ReadFile(fp)
	if err != nil {
		panic(err)
	}

	txt := string(dat)
	txt = strings.TrimRight(txt, "\n")
	directions := strings.Split(txt, "\n")

	return directions
}

// part a

func main() {
	ms := MakeDirections(fp)
	fmt.Printf("%v", ms[1]) // 3337766
}
