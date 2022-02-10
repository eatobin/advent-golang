package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

var fp = "03/day03a.csv"

func MakeBoth(fp string) [][]string {
	both := make([][]string, 2)

	f, err := os.Open(fp)
	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
		}
	}(f)

	csvReader := csv.NewReader(f)
	i := 0
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		both[i] = rec
		i++
	}
	return both
}

func main() {
	var red []string
	var blue []string

	both := MakeBoth(fp)

	red = both[0]
	blue = both[1]

	fmt.Printf("red = %+v\n", red)
	fmt.Printf("blue = %+v\n", blue)
}
