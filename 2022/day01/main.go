package main

import (
	"bufio"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/MikeOlsen/aoc/2022/common"
)

func main() {
	test_data := common.LoadInput("test", "01")
	prod_data := common.LoadInput("prod", "01")

	fmt.Println("=============== PART 1 ==================")
	fmt.Println("Test:", findBiggestElves(test_data, 3, false))
	fmt.Println("Prod:", findBiggestElves(prod_data, 3, false))

	fmt.Println("=============== PART 2 ==================")
	fmt.Println("Test:", findBiggestElves(test_data, 3, false))
	fmt.Println("Prod:", findBiggestElves(prod_data, 3, false))
}

func findBiggestElves(input string, count int, verbose bool) int {
	elves := createElves(input, verbose)
	sort.Ints(elves)
	if verbose {
		fmt.Println(elves)
	}
	top := elves[len(elves)-count:]
	return sum(top)
}

func createElves(input string, verbose bool) []int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var s []int
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			s = append(s, count)
			count = 0
		} else {
			calories, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal("Not a number", err)
			}
			count += calories
		}
	}
	s = append(s, count)
	return s
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
