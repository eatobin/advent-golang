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

func makePath(move string, start []int) [][]int {
	direction := []byte(direction(move))[0]
	distance := distance(move)
	xStart := start[0]
	yStart := start[1]
	path := make([][]int, distance+1)

	switch direction {
	case 'R':
		for i := 0; i < distance+1; i++ {
			path[i] = make([]int, 2)
		}
		for i := 0; i < distance+1; i++ {
			path[i] = []int{xStart, yStart}
			xStart++
		}
	case 'U':
		for i := 0; i < distance+1; i++ {
			path[i] = make([]int, 2)
		}
		for i := 0; i < distance+1; i++ {
			path[i] = []int{xStart, yStart}
			yStart++
		}
	case 'L':
		for i := 0; i < distance+1; i++ {
			path[i] = make([]int, 2)
		}
		for i := 0; i < distance+1; i++ {
			path[i] = []int{xStart, yStart}
			xStart--
		}
	case 'D':
		for i := 0; i < distance+1; i++ {
			path[i] = make([]int, 2)
		}
		for i := 0; i < distance+1; i++ {
			path[i] = []int{xStart, yStart}
			yStart--
		}
	}

	return path
}

func makePaths(start []int, moves []string) [][]int {
	var flatPaths [][]int
	paths := make([][][]int, len(moves))
	pathStart := start
	for i, move := range moves {
		path := makePath(move, pathStart)
		paths[i] = path
		pathStart = path[len(path)-1]
	}
	for _, path := range paths {
		for _, move := range path {
			flatPaths = append(flatPaths, move)
		}
	}
	return flatPaths
}

var red []string
var blue []string

func main() {
	both := make([][]string, 2)
	// open file
	f, err := os.Open("03/cc.csv")
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

	red = both[0]
	blue = both[1]

	fmt.Printf("%+v\n", red)
	fmt.Printf("%+v\n", blue)

	sliceOfSlices := makePath("L3", []int{1, 1})
	fmt.Println("\nSlice of slices: ", sliceOfSlices)

	paths := makePaths([]int{0, 0}, red)
	fmt.Println("\nPaths: ", paths)

	fmt.Println("\nPathsA: ", paths[0])
	fmt.Println("\nPathsB: ", paths[0][0])
	//fmt.Println("\nPathsC: ", paths[0][0][0])
}
