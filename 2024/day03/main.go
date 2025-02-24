package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/MikeOlsen/aoc/2024/common"
)

func main() {
	test_data := common.LoadInputLines("test", "03")
	prod_data := common.LoadInputLines("prod", "03")

	test_data_1 := []string{test_data[0]}
	test_data_2 := []string{test_data[1]}

	fmt.Println("=============== PART 1 ==================")
	fmt.Println("Test:", part1(test_data_1))
	fmt.Println("Prod:", part1(prod_data))

	fmt.Println("=============== PART 2 ==================")
	fmt.Println("Test:", part2(test_data_2))
	fmt.Println("Prod:", part2(prod_data))
}

func part1(input []string) int {

	pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	sum := 0
	for _, line := range input {
		expressions := pattern.FindAllStringSubmatch(line, -1)
		for _, expression := range expressions {
			left, err := strconv.Atoi(expression[1])
			if err != nil {
				panic(err)
			}
			right, err := strconv.Atoi(expression[2])
			if err != nil {
				panic(err)
			}
			sum += left * right
		}
	}

	return sum
}

func part2(input []string) int {

	data := strings.Join(input, "")

	exclusionPattern := regexp.MustCompile(`don\'t\(\).*?do\(\)|don't\(\).*`)
	matchPattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	sum := 0

	data = exclusionPattern.ReplaceAllString(data, "")
	expressions := matchPattern.FindAllStringSubmatch(data, -1)

	for _, expression := range expressions {
		left, err := strconv.Atoi(expression[1])
		if err != nil {
			panic(err)
		}
		right, err := strconv.Atoi(expression[2])
		if err != nil {
			panic(err)
		}
		sum += left * right
	}
	return sum
}
