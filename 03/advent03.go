package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type visit struct {
	x, y int
}
type path = []visit
type route = []path

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

func makePath(move string, start visit) path {
	direction := []byte(direction(move))[0]
	distance := distance(move)
	xStart := start.x
	yStart := start.y
	path := make(path, distance+1)

	switch direction {
	case 'R':
		for i := 0; i < distance+1; i++ {
			path[i] = visit{xStart, yStart}
			xStart++
		}
	case 'U':
		for i := 0; i < distance+1; i++ {
			path[i] = visit{xStart, yStart}
			yStart++
		}
	case 'L':
		for i := 0; i < distance+1; i++ {
			path[i] = visit{xStart, yStart}
			xStart--
		}
	case 'D':
		for i := 0; i < distance+1; i++ {
			path[i] = visit{xStart, yStart}
			yStart--
		}
	}
	return path
}

func makeRoute(start visit, moves []string) route {
	//	var flatPaths [][]int
	route := make(route, len(moves))
	//	paths := make([][][]int, len(moves))
	pathStart := start
	for i, move := range moves {
		path := makePath(move, pathStart)
		route[i] = path
		pathStart = path[len(path)-1]
	}
	//	for _, path := range paths {
	//		for _, move := range path {
	//			flatPaths = append(flatPaths, move)
	//		}
	//	}
	return route
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

	sliceOfSlices := makePath("D3", visit{x: 0, y: 1})
	fmt.Println("\nSlice of slices: ", sliceOfSlices)

	paths := makeRoute(visit{x: 0, y: 0}, red)
	fmt.Println("\nPaths: ", paths)

	//fmt.Println("\nPathsA: ", paths[0])
	//fmt.Println("\nPathsB: ", paths[0][0])
	////fmt.Println("\nPathsC: ", paths[0][0][0])
}

//type visit struct {
//	x, y int
//}
//
//func main() {
//	visited := []visit{
//		visit{1, 100},
//		visit{2, 2},
//		visit{1, 100},
//		visit{1, 1},
//	}
//
//	var unique []visit
//	m := map[visit]bool{}
//
//	for _, v := range visited {
//		if !m[v] {
//			m[v] = true
//			unique = append(unique, v)
//		}
//	}
//
//	fmt.Println(unique)
//}
