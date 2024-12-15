package main

import (
	"fmt"

	"github.com/MikeOlsen/aoc/2024/common"
)

type Point struct {
	x int
	y int
}

type Line struct {
	p1 Point
	p2 Point
}

func main() {
	test_data := common.LoadInputLines("test", "08")
	prod_data := common.LoadInputLines("prod", "08")

	fmt.Println("=============== PART 1 ==================")
	fmt.Println("Test:", part1(test_data))
	fmt.Println("Prod:", part1(prod_data))

	fmt.Println("=============== PART 2 ==================")
	fmt.Println("Test:", part2(test_data))
	fmt.Println("Prod:", part2(prod_data))
}

func part1(input []string) int {
	grid := parseGrid(input)
	var nodes = make(map[int]map[int]bool)
	count := 0
	// Find first antenna
	for y, row := range grid {
		for x, char := range row {
			if char != '.' {
				// Find matching antennas
				antennas := findAntennas(grid, char, x, y)
				for _, antenna := range antennas {
					line := Line{Point{x, y}, Point{antenna.x, antenna.y}}
					dx, dy := line.p2.x-line.p1.x, line.p2.y-line.p1.y
					node1 := Point{line.p1.x + 2*dx, line.p1.y + 2*dy}
					node2 := Point{line.p1.x - 1*dx, line.p1.y - 1*dy}
					if addAntiNode(grid, nodes, node1) {
						count++
					}
					if addAntiNode(grid, nodes, node2) {
						count++
					}
				}
			}
		}
	}
	return count
}

func part2(input []string) int {
	grid := parseGrid(input)
	var nodes = make(map[int]map[int]bool)
	count := 0
	// Find first antenna
	for y, row := range grid {
		for x, char := range row {
			if char != '.' {
				// Find matching antennas
				antennas := findAntennas(grid, char, x, y)
				for _, antenna := range antennas {
					line := Line{Point{x, y}, Point{antenna.x, antenna.y}}
					dx, dy := line.p2.x-line.p1.x, line.p2.y-line.p1.y
					count += addAntiNodes(grid, nodes, line, dx, dy)
				}
			}
		}
	}
	return count
}

func findAntennas(grid [][]rune, freq rune, startX, startY int) []Point {
	var antennas []Point
	for y, line := range grid {
		for x, char := range line {
			if y < startY || y == startY && x <= startX {
				continue
			}
			if char == freq {
				antennas = append(antennas, Point{x, y})
			}
		}
	}
	return antennas
}

func addAntiNodes(grid [][]rune, nodes map[int]map[int]bool, line Line, dx, dy int) int {

	count := 0

	// Forwards until out of bounds, starting at 0 to mark towers
	for n := 0; true; n++ {
		node := Point{line.p1.x + dx*n, line.p1.y + dy*n}
		if node.x >= len(grid[0]) || node.y >= len(grid) {
			break
		}
		if addAntiNode(grid, nodes, node) {
			count++
		}
	}

	// Backwards until out of bounds
	for n := 1; true; n++ {
		node := Point{line.p1.x - dx*n, line.p1.y - dy*n}
		if node.x < 0 || node.y < 0 {
			break
		}
		if addAntiNode(grid, nodes, node) {
			count++
		}

	}

	return count
}

func addAntiNode(grid [][]rune, nodes map[int]map[int]bool, node Point) bool {

	// Check bounds
	if node.x < 0 || node.y < 0 || node.x >= len(grid[0]) || node.y >= len(grid) {
		return false
	}

	// Check if overlap
	if nodes[node.y][node.x] {
		return false
	}

	// Create inner slice if missing
	if _, exists := nodes[node.y]; !exists {
		nodes[node.y] = make(map[int]bool)
	}

	nodes[node.y][node.x] = true
	return true
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
