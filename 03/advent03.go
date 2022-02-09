package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	both := make([][]string, 2)
	// open file
	f, err := os.Open("03/day03a.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
		}
	}(f)

	// read csv values using csv.Reader
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
		// do something with read line
		both[i] = rec
		i++
	}
	fmt.Printf("%+v\n", both)
	fmt.Printf("%+v\n", both[0])
	fmt.Printf("%+v\n", both[1])
}
