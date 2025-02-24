package main

import (
	"fmt"

	"github.com/MikeOlsen/aoc/2024/common"
)

type Grid map[int]map[int]rune
type Visited map[int]map[int]bool

var directions = [][2]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}
var cornerDirections = [][2][2]int{
	{{1, 0}, {0, -1}},
	{{1, 0}, {0, 1}},
	{{-1, 0}, {0, -1}},
	{{-1, 0}, {0, 1}},
}

type Area struct {
	name      string
	size      int
	perimeter int
	sides     int
}

func main() {
	test_data := common.LoadInputLines("test", "12")
	prod_data := common.LoadInputLines("prod", "12")

	fmt.Println("=============== PART 1 ==================")
	fmt.Println("Test:", part1(test_data))
	fmt.Println("Prod:", part1(prod_data))

	fmt.Println("=============== PART 2 ==================")
	fmt.Println("Test:", part2(test_data))
	fmt.Println("Prod:", part2(prod_data))
}

func part1(input []string) int {
	grid := parseGrid(input)
	areas := findAreas(grid)

	sum := 0
	for _, area := range areas {
		sum += area.size * area.perimeter
	}
	return sum
}

func part2(input []string) int {
	grid := parseGrid(input)
	areas := findAreas(grid)

	sum := 0
	for _, area := range areas {
		sum += area.size * area.sides
	}
	return sum
}

func findAreas(grid Grid) []Area {
	visited := make(map[int]map[int]bool)
	for i := 0; i < len(grid); i++ {
		visited[i] = make(map[int]bool)
	}

	var areas []Area
	for x := 0; x < len(grid[0]); x++ {
		for y := 0; y < len(grid); y++ {
			if !visited[x][y] {
				char := grid[x][y]
				area, perims, corners := floodFill(grid, visited, x, y, char)
				areas = append(areas, Area{string(char), area, perims, corners})
			}
		}
	}

	return areas
}

func floodFill(grid Grid, visited Visited, x, y int, char rune) (int, int, int) {
	_, exists := grid[x][y]
	if !exists || char != grid[x][y] {
		return 0, 1, 0
	}
	if visited[x][y] {
		return 0, 0, 0
	}

	visited[x][y] = true
	area, perims, corners := 1, 0, 0
	for _, dir := range cornerDirections {
		dir1, exists1 := grid[x+dir[0][0]][y+dir[0][1]]
		dir2, exists2 := grid[x+dir[1][0]][y+dir[1][1]]
		dir3, exists3 := grid[x+dir[0][0]][y+dir[1][1]]
		if (!exists1 || dir1 != char) && (!exists2 || dir2 != char) {
			corners++ // Inner corner
		} else if exists1 && dir1 == char && exists2 && dir2 == char {
			if exists3 && dir3 != char {
				corners++ // Outer corner
			}
		}
	}
	for _, dir := range directions {
		a, p, c := floodFill(grid, visited, x+dir[0], y+dir[1], char)
		area += a
		perims += p
		corners += c
	}

	return area, perims, corners
}

func parseGrid(input []string) Grid {
	grid := make(Grid)
	for y, line := range input {
		for x, char := range line {
			if _, exists := grid[x]; !exists {
				grid[x] = make(map[int]rune)
			}
			grid[x][y] = char
		}
	}
	return grid
}
