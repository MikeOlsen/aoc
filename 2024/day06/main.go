package main

import (
	"fmt"

	"github.com/MikeOlsen/aoc/2024/common"
)

type point struct {
	x int
	y int
}

type meta struct {
	position  point
	direction point
	bounds    point
	obstacles map[int]map[int]bool
	visited   map[int]map[int]map[point]bool
}

func main() {
	test_data := common.LoadInputLines("test", "06")
	prod_data := common.LoadInputLines("prod", "06")

	fmt.Println("=============== PART 1 ==================")
	fmt.Println("Test:", part1(test_data))
	fmt.Println("Prod:", part1(prod_data))

	fmt.Println("=============== PART 2 ==================")
	fmt.Println("Test:", part2(test_data))
	fmt.Println("Prod:", part2(prod_data))
}

func part1(input []string) int {
	data := parseInput(input)
	walk(&data)
	return countVisited(data)
}

func part2(input []string) int {
	count := 0
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {
			data := parseInput(input)
			// Skip start position and obstacles
			if x == data.position.x && y == data.position.y || data.obstacles[x][y] {
				continue
			}
			addObstacle(data, x, y)
			if walk(&data) {
				count++
			}
		}
	}
	return count
}

func addObstacle(data meta, x, y int) meta {
	data.obstacles[x][y] = true
	return data
}

func walk(data *meta) bool {
	for walk := true; walk; {
		collision, end := takeStep(data)
		if collision {
			rotateClockwize(data)
		}
		if checkIfVisited(data) {
			return true
		}
		if end {
			return false
		}
	}
	return true
}

func takeStep(data *meta) (bool, bool) {
	x := data.position.x + data.direction.x
	y := data.position.y + data.direction.y

	if x < 0 || y < 0 || x >= data.bounds.x || y >= data.bounds.y {
		return false, true
	}
	if data.obstacles[x][y] {
		return true, false
	}
	data.position.x = x
	data.position.y = y
	markVisited(data)
	return false, false
}

func parseInput(input []string) meta {
	var m meta

	m.visited = make(map[int]map[int]map[point]bool)
	m.position.x, m.position.y = findGuard(input)
	m.direction.x, m.direction.y = 0, -1 // Start going upwards
	m.obstacles = parseObstacles(input)

	m.bounds.x = len(input[0])
	m.bounds.y = len(input)

	markVisited(&m)

	return m
}

func markVisited(data *meta) {
	if _, exists := data.visited[data.position.x]; !exists {
		data.visited[data.position.x] = make(map[int]map[point]bool)
	}
	if _, exists := data.visited[data.position.x][data.position.y]; !exists {
		data.visited[data.position.x][data.position.y] = make(map[point]bool)
	}

	data.visited[data.position.x][data.position.y][data.direction] = true
}

func checkIfVisited(data *meta) bool {
	x := data.position.x + data.direction.x
	y := data.position.y + data.direction.y
	_, seen := data.visited[x][y][data.direction]
	return seen
}

func parseObstacles(input []string) map[int]map[int]bool {
	obstacles := make(map[int]map[int]bool)
	for y, line := range input {
		for x, val := range line {
			if _, exists := obstacles[x]; !exists {
				obstacles[x] = make(map[int]bool)
			}
			if val == '#' {
				obstacles[x][y] = true
			}
		}
	}
	return obstacles
}

func getBounds(input []string) (int, int) {
	return len(input), len(input[0])
}

func findGuard(input []string) (int, int) {
	for y, line := range input {
		for x, val := range line {

			if val == '^' {
				return x, y
			}
		}
	}
	panic("Guard not found")
}

func countVisited(data meta) int {
	visited := 0
	for y := 0; y < data.bounds.y; y++ {
		for x := 0; x < data.bounds.x; x++ {
			if _, exists := data.visited[x][y]; exists {
				visited++
			}
		}
	}
	return visited
}

func rotateClockwize(data *meta) {
	data.direction.x, data.direction.y = -data.direction.y, data.direction.x
}

func printWalk(data meta) {
	for y := 0; y < data.bounds.y; y++ {
		for x := 0; x < data.bounds.x; x++ {
			if _, exists := data.visited[x][y]; exists {
				fmt.Print("X")
				continue
			}
			if data.obstacles[x][y] {
				fmt.Print("#")
				continue
			}
			fmt.Print(".")
		}
		fmt.Println()
	}
	fmt.Println()
}
