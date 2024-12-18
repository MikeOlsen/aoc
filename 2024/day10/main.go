package main

import (
	"fmt"

	"github.com/MikeOlsen/aoc/2024/common"
)

type Position struct {
	x int
	y int
}

type Trail struct {
	start Position
	peaks map[int]map[int]bool
}

type Topography map[int]map[int]int
type PeakMap map[int]map[int]bool

func main() {
	test_data := common.LoadInputLines("test", "10")
	prod_data := common.LoadInputLines("prod", "10")

	fmt.Println("=============== PART 1 ==================")
	fmt.Println("Test:", countTrails(test_data, true))
	fmt.Println("Prod:", countTrails(prod_data, true))

	fmt.Println("=============== PART 2 ==================")
	fmt.Println("Test:", countTrails(test_data, false))
	fmt.Println("Prod:", countTrails(prod_data, false))
}

func countTrails(input []string, checkIfVisited bool) int {
	topography, starts := parseTopography(input)

	count := 0
	for _, position := range starts {
		peaks := make(PeakMap)
		count += hike(topography, peaks, position.x, position.y, 0, checkIfVisited)
	}
	return count
}

func hike(topography Topography, peaks PeakMap, x, y, prevVal int, check bool) int {

	// Stop if peak is found
	if topography[x][y] == 9 {
		if check {
			return checkIfVisited(peaks, x, y)
		}
		return 1
	}

	// Test hiking all directions
	count := 0
	dx, dy := 1, 0
	for i := 0; i < 4; i++ {
		stepX, stepY := x+dx, y+dy
		val := topography[stepX][stepY]
		if val == prevVal+1 {
			count += hike(topography, peaks, stepX, stepY, val, check)
		}
		dx, dy = -dy, dx
	}
	return count
}

func parseTopography(input []string) (Topography, []Position) {
	starts := make([]Position, 0)
	topography := make(Topography)
	for y, line := range input {
		for x, char := range line {
			height := int(char - '0')
			placeMarker(topography, x, y, height)
			if height == 0 {
				starts = append(starts, Position{x, y})
			}
		}
	}
	return topography, starts
}

func checkIfVisited(peaks PeakMap, x, y int) int {
	if _, exists := peaks[x]; !exists {
		peaks[x] = make(map[int]bool)
	}
	if !peaks[x][y] {
		peaks[x][y] = true
		return 1
	}
	return 0
}

func placeMarker(input Topography, x, y, val int) {
	if _, exists := input[x]; !exists {
		input[x] = make(map[int]int)
	}
	input[x][y] = val
}
