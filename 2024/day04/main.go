package main

import (
	"fmt"

	"github.com/MikeOlsen/aoc/2024/common"
)

func main() {
	test_data := common.LoadInputLines("test", "04")
	prod_data := common.LoadInputLines("prod", "04")

	fmt.Println("=============== PART 1 ==================")
	fmt.Println("Test:", part1(test_data))
	fmt.Println("Prod:", part1(prod_data))

	fmt.Println("=============== PART 2 ==================")
	fmt.Println("Test:", part2(test_data))
	fmt.Println("Prod:", part2(prod_data))
}

var directions = [][]int{
	{1, 0},   // →
	{1, 1},   // ↘
	{0, 1},   // ↓
	{-1, 1},  // ↙
	{-1, 0},  // ←
	{-1, -1}, // ↖
	{0, -1},  // ↑
	{1, -1},  // ↗
}

func part1(input []string) int {
	grid := parseGrid(input)

	count := 0
	for r, row := range grid {
		for c := range row {
			for _, dir := range directions {
				if checkDirection(grid, c, r, dir[0], dir[1], "XMAS") {
					count++
				}

			}
		}
	}

	return count
}

func part2(input []string) int {
	grid := parseGrid(input)
	count := 0
	for r, row := range grid {
		for c := range row {
			if checkCross(grid, c, r, "MAS") {
				count++
			}
		}
	}

	return count
}

func checkDirection(grid [][]rune, x, y, dx, dy int, word string) bool {
	rows := len(grid)
	cols := len(grid[1])

	for i, char := range word {
		curX := x + i*dx
		curY := y + i*dy

		if !checkBounds(curX, curY, cols, rows) || char != grid[curY][curX] {
			return false
		}
	}
	return true
}

func checkCross(grid [][]rune, x, y int, word string) bool {
	// Check ↘↖
	line1 := checkDirection(grid, x, y, 1, 1, word) || checkDirection(grid, x+2, y+2, -1, -1, word)
	// Check ↗↙
	line2 := checkDirection(grid, x, y+2, 1, -1, word) || checkDirection(grid, x+2, y, -1, 1, word)

	return line1 && line2
}

func checkBounds(x, y, xBound, yBound int) bool {
	return x >= 0 && y >= 0 && x < xBound && y < yBound
}

func parseGrid(input []string) [][]rune {
	var grid [][]rune
	for _, line := range input {
		var row []rune
		for _, char := range line {
			row = append(row, char)
		}
		grid = append(grid, row)
	}
	return grid
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		for _, r := range row {
			fmt.Printf("%c", r)
		}
		fmt.Println()
	}
}
