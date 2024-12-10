package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/MikeOlsen/aoc/2024/common"
)

func main() {
	test_data := common.LoadInput("test", "02")
	prod_data := common.LoadInput("prod", "02")

	fmt.Println("=============== PART 1 ==================")
	fmt.Println("Test:", validateReports(parseReports(test_data), false))
	fmt.Println("Prod:", validateReports(parseReports(prod_data), false))

	fmt.Println("=============== PART 2 ==================")
	fmt.Println("Test:", validateReports(parseReports(test_data), true))
	fmt.Println("Prod:", validateReports(parseReports(prod_data), true))
}

func validateReports(reports [][]int, dampener bool) int {
	count := 0

	for _, report := range reports {

		safe := validateReport(report)

		if !safe && dampener {
			safe = removeLevers(report)
		}

		if safe {
			count++
		}
	}

	return count
}

func validateReport(report []int) bool {
	rising := report[1] > report[0]

	for i := 1; i < len(report); i++ {
		dist := report[i] - report[i-1]

		// Statement 1: The levels are either all increasing or all decreasing
		if (rising && dist < 0) || (!rising && dist > 0) {
			return false
		}
		// Statement 2: Any two adjacent levels differ by at least one and at most three
		if common.Abs(dist) < 1 || common.Abs(dist) > 3 {
			return false
		}
	}
	return true
}

func removeLevers(report []int) bool {
	variations := generateVariations(report)
	for _, variation := range variations {
		if validateReport(variation) {
			return true
		}
	}
	return false
}

func generateVariations(report []int) [][]int {
	var variations [][]int

	for i := 0; i < len(report); i++ {
		modifiedReport := append([]int{}, report[:i]...)
		modifiedReport = append(modifiedReport, report[i+1:]...)
		variations = append(variations, modifiedReport)
	}

	return variations
}

func parseReports(input string) [][]int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var reports [][]int
	for scanner.Scan() {
		line := scanner.Text()
		parts := sliceAtoi(strings.Fields(line))
		reports = append(reports, parts)
	}
	return reports
}

func sliceAtoi(input []string) []int {
	var result []int
	for _, i := range input {
		val, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		result = append(result, val)
	}
	return result
}
