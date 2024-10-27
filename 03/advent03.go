package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
)

type visit struct {
	x, y int
}
type path = []visit
type route = []path
type flatRoute = []visit
type uniqueRoute = []visit

func MakeBoth(fp string) [][]string {
	both := make([][]string, 2)

	f, err1 := os.Open(fp)
	if err1 != nil {
		log.Fatal(err1)
	}

	defer func(f *os.File) {
		err2 := f.Close()
		if err2 != nil {
			return
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
	route := make(route, len(moves))
	pathStart := start
	for i, move := range moves {
		path := makePath(move, pathStart)
		route[i] = path
		pathStart = path[len(path)-1]
	}
	return route
}

func makeFlatRoute(route route) flatRoute {
	flatRoute := flatRoute{}

	for _, path := range route {
		flatRoute = append(flatRoute, path...)
	}
	return flatRoute
}

func makeUniqueRoute(start visit, moves []string) uniqueRoute {
	route := makeRoute(start, moves)
	flatRoute := makeFlatRoute(route)
	var unique uniqueRoute
	m := map[visit]bool{}

	for _, v := range flatRoute {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}
	return unique
}

// function for finding the intersection of two arrays
func findIntersection(uniqueRedRoute, uniqueBlueRoute uniqueRoute) []visit {
	intersection := make([]visit, 0)

	set := make(map[visit]bool)

	// Create a set from the first array
	for _, visit := range uniqueRedRoute {
		set[visit] = true // setting the initial value to true
	}

	// Check elements in the second array against the set
	for _, visit := range uniqueBlueRoute {
		if set[visit] {
			intersection = append(intersection, visit)
		}
	}

	return intersection[1:]
}

func manhattanizeAVisitFromOrigin(visit visit) int {
	return int(math.Abs(float64(visit.x-0)) + math.Abs(float64(visit.y-0)))
}

func manhattanizedSlice(intersection []visit) []int {
	manhattans := make([]int, len(intersection))

	for i, visit := range intersection {
		manhattans[i] = manhattanizeAVisitFromOrigin(visit)
	}
	return manhattans
}

func main() {
	var fp = "day03.csv"
	both := MakeBoth(fp)
	var red []string
	var blue []string

	red = both[0]
	blue = both[1]

	uniqueRedRoute := makeUniqueRoute(visit{x: 0, y: 0}, red)

	uniqueBlueRoute := makeUniqueRoute(visit{x: 0, y: 0}, blue)

	intersections := findIntersection(uniqueRedRoute, uniqueBlueRoute)
	// fmt.Println("\nRoutes Intersect at: ", intersections)

	manhattans := manhattanizedSlice(intersections)
	// fmt.Println("\nManhattanized slice: ", manhattans)

	fmt.Println("Minimum distance: ", slices.Min(manhattans)) // 2193
}
