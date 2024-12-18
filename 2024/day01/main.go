package main

import (
	"bufio"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/MikeOlsen/aoc/2024/common"
)

func main() {
	test_data := common.LoadInput("test", "01")
	prod_data := common.LoadInput("prod", "01")

	fmt.Println("=============== PART 1 ==================")
	fmt.Println("Test:", calculateDistance(readSortedColumns(test_data)))
	fmt.Println("Prod:", calculateDistance(readSortedColumns(prod_data)))

	fmt.Println("=============== PART 2 ==================")
	fmt.Println("Test:", calculateSimilarityScore(readSortedColumns(test_data)))
	fmt.Println("Prod:", calculateSimilarityScore(readSortedColumns(prod_data)))
}

func calculateSimilarityScore(input [][]int) int {
	similarity := 0
	left, right := input[0], input[1]

	for i := range left {
		count := 0
		for j := range right {
			if left[i] == right[j] {
				count++
			}
		}
		similarity += left[i] * count
	}
	return similarity
}

func calculateDistance(input [][]int) int {
	left, right := input[0], input[1]
	dist := 0
	for i := range left {
		dist += abs(left[i] - right[i])
	}
	return dist
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func readSortedColumns(input string) [][]int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var left, right []int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		l, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("Invalid number in input: %v", err)
		}
		r, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("Invalid number in input: %v", err)
		}

		left = append(left, l)
		right = append(right, r)
	}

	sort.Ints(left)
	sort.Ints(right)

	return [][]int{left, right}
}
