package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
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

func direction(unit string) string {
	return unit[:1]
}

func distance(unit string) int {
	strDist := unit[1:]
	dist, err := strconv.Atoi(strDist)
	if err != nil {
		log.Fatal(err)
	}
	return dist
}

//func makePath(unit string, start []int) [][]int {
//	direction := []byte(direction(unit))[0]
//	distance := distance(unit)
//	xStart := start[0]
//	yStart := start[1]
//	path := make([][]int, distance+1)
//
//	switch direction {
//	case 'R':
//		//slice_of_slices := make([][]int , 3)
//
//		for i := 0; i < distance; i++ {
//			// looping through the slice to declare
//			// slice of slice of length 3
//			path[i] = make([]int, 2)
//
//			// assigning values to each
//			// slice of a slice
//			for x := xStart; x < xStart+distance+1; x++{
//				path[i][j] = i * j
//			}
//		}
//	}
//
//	return nil
//}

func main() {
	var red []string
	var blue []string

	both := MakeBoth(fp)

	red = both[0]
	blue = both[1]

	fmt.Printf("red = %+v\n", red)
	fmt.Printf("blue = %+v\n", blue)

	country := "London"
	firstCharacter := country[0:1]

	fmt.Println(firstCharacter)
}
