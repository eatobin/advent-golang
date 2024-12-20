package makeMemory

import (
	"os"
	"strconv"
	"strings"
)

type Memory []int

type Intcode struct {
	Pointer int
	Memory  Memory
}

func MakeMemory(fp string) Memory {
	dat, err := os.ReadFile(fp)
	if err != nil {
		panic(err)
	}

	txt := string(dat)
	txt = strings.TrimRight(txt, "\n")
	strOps := strings.Split(txt, ",")
	length := len(strOps)
	memory := make([]int, length)

	for i, strOp := range strOps {
		op, err := strconv.Atoi(strOp)
		if err != nil {
			panic(err)
		}
		memory[i] = op
	}
	return memory
}
