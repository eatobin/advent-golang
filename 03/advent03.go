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
type flatRoute = []visit
type uniqueRoute = []visit

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
		for _, visit := range path {
			flatRoute = append(flatRoute, visit)
		}
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
func findIntersection(redFlatRoute, blueFlatRoute uniqueRoute) []visit {
	intersection := make([]visit, 0)

	set := make(map[visit]bool)

	// Create a set from the first array
	for _, visit := range redFlatRoute {
		set[visit] = true // setting the initial value to true
	}

	// Check elements in the second array against the set
	for _, visit := range blueFlatRoute {
		if set[visit] {
			intersection = append(intersection, visit)
		}
	}

	return intersection[1:]
}

func main() {
	var fp = "day03a.csv"
	both := MakeBoth(fp)
	var red []string
	var blue []string

	red = both[0]
	blue = both[1]

	fmt.Printf("%+v\n", red)
	fmt.Printf("%+v\n", blue)

	redRoutes := makeRoute(visit{x: 0, y: 0}, red)
	fmt.Println("\nRedRoutes: ", redRoutes)

	redFlatRoute := makeFlatRoute(redRoutes)
	fmt.Println("\nRedFlatRoute: ", redFlatRoute)

	uniqueRedRoute := makeUniqueRoute(visit{x: 0, y: 0}, red)
	fmt.Println("\nUniqueRedRoute: ", uniqueRedRoute)

	blueRoutes := makeRoute(visit{x: 0, y: 0}, blue)
	fmt.Println("\nBlueRoutes: ", blueRoutes)

	blueFlatRoute := makeFlatRoute(blueRoutes)
	fmt.Println("\nBlueFlatRoute: ", blueFlatRoute)

	uniqueBlueRoute := makeUniqueRoute(visit{x: 0, y: 0}, blue)
	fmt.Println("\nUniqueBlueRoute: ", uniqueBlueRoute)

	intersections := findIntersection(redFlatRoute, blueFlatRoute)
	fmt.Println("\nRoutes Intersect at: ", intersections)
}
