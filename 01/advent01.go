package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Modules = []int

const fp = "01/advent01.txt"

func MakeModules(fp string) Modules {
	dat, err := ioutil.ReadFile(fp)
	if err != nil {
		panic(err)
	}

	txt := string(dat)
	lines := strings.Split(txt, "\n")
	strValues := lines[:len(lines)-1]
	modules := make([]int, len(strValues))
	for i, str := range strValues {
		mass, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		modules[i] = mass
	}
	return modules
}

//	fuel := Fuel(mass)
//	sum += fuel
//}
//fmt.Println(sum)
//dat, err := ioutil.ReadFile(fp)
//if err != nil {
//	panic(err)
//}
//
//txt := string(dat)
//txt = strings.TrimRight(txt, "\n")
//strOps := strings.Split(txt, "\n")
//modules := make([]int, len(strOps))
//
//for i, strOp := range strOps {
//	op, err := strconv.Atoi(strOp)
//	if err != nil {
//		panic(err)
//	}
//	modules[i] = op
//}
//return modules
//}

// part a

func Fuel(m int) int {
	return (m / 3) - 2
}

func main() {
	sum := 0
	ms := MakeModules(fp)
	for _, v := range ms {
		fuel := Fuel(v)
		sum += fuel
	}
	fmt.Printf("Part A: %d", sum) // 3337766
}
