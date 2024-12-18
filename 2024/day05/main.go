package main

import (
	"fmt"
	"strings"

	"github.com/MikeOlsen/aoc/2024/common"
)

type Content struct {
	Rules     [][]int
	Updates   [][]int
	Unordered map[int]struct{}
}

func main() {
	test_data := common.LoadInputLines("test", "05")
	prod_data := common.LoadInputLines("prod", "05")

	fmt.Println("=============== PART 1 ==================")
	fmt.Println("Test:", part1(test_data))
	fmt.Println("Prod:", part1(prod_data))

	fmt.Println("=============== PART 2 ==================")
	fmt.Println("Test:", part2(test_data))
	fmt.Println("Prod:", part2(prod_data))
}

func part1(input []string) int {
	data := parseContent(input)

	var filtered [][]int
	for i, row := range data.Updates {
		if _, exists := data.Unordered[i]; !exists {
			filtered = append(filtered, row)
		}
	}

	return getMiddleSum(filtered)
}

func part2(input []string) int {
	data := parseContent(input)

	var filtered [][]int
	for i, row := range data.Updates {
		if _, exists := data.Unordered[i]; exists {
			filtered = append(filtered, row)
		}
	}

	fixed := fixUnorderedUpdates(filtered, data.Rules)

	return getMiddleSum(fixed)
}

func parseContent(input []string) Content {
	var rules, updates [][]int

	parsingRules := true

	for _, line := range input {
		if line == "" {
			parsingRules = false
			continue
		}
		if parsingRules {
			rules = append(rules, common.SliceAtoi(strings.Split(line, "|")))
		} else {
			updates = append(updates, common.SliceAtoi(strings.Split(line, ",")))
		}
	}
	unordered := findUnorderedUpdates(rules, updates)

	return Content{rules, updates, unordered}
}

func findUnorderedUpdates(rules, updates [][]int) map[int]struct{} {
	unordered := make(map[int]struct{})
	for updateIdx, update := range updates {
	out:
		for pageIdx, page := range update {
			for _, rule := range rules {
				if rule[0] == page {
					for i := 0; i < pageIdx; i++ {
						if update[i] == rule[1] {
							unordered[updateIdx] = struct{}{}
							break out
						}
					}
				}
			}
		}
	}
	return unordered
}

func fixUnorderedUpdates(updates, rules [][]int) [][]int {

	ordered := make([][]int, len(updates))
	for i := range updates {
		ordered[i] = make([]int, len(updates[i]))
		copy(ordered[i], updates[i])
	}

	for y, row := range ordered {

		// Continue until no more swaps
		for done := false; !done; {
			done = true
		out:
			for x, page := range row {

				// Apply rules
				for _, rule := range rules {
					if rule[0] == page {

						// Check if before
						for i := 0; i < x; i++ {

							if row[i] == rule[1] {

								// Swap pages
								ordered[y][x], ordered[y][i] = ordered[y][i], ordered[y][x]

								done = false
								break out
							}
						}
					}
				}
			}
		}

	}

	return ordered
}

func getMiddleSum(input [][]int) int {
	var middle []int
	for _, row := range input {

		if len(row)%2 == 0 {
			panic("Found an even row")
		}

		idx := len(row) / 2
		middle = append(middle, row[idx])

	}

	sum := 0
	for _, val := range middle {
		sum += val
	}
	return sum
}
