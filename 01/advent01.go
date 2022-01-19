package main

import (
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
	txt = strings.TrimRight(txt, "\n")
	strOps := strings.Split(txt, "\n")
	modules := make([]int, len(strOps))

	for i, strOp := range strOps {
		op, err := strconv.Atoi(strOp)
		if err != nil {
			panic(err)
		}
		modules[i] = op
	}
	return modules
}

func main() {
	println(MakeModules(fp))
}
