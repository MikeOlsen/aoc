package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/MikeOlsen/aoc/2024/common"
)

type Operator int

type equation struct {
	answer int
	values []int
}

var log common.Logger

func main() {
	test_data := common.LoadInputLines("test", "07")
	prod_data := common.LoadInputLines("prod", "07")

	fmt.Println("=============== PART 1 ==================")
	fmt.Println("Test:", part1(test_data, true))
	fmt.Println("Prod:", part1(prod_data, false))

	fmt.Println("=============== PART 2 ==================")
	fmt.Println("Test:", part2(test_data, true))
	fmt.Println("Prod:", part2(prod_data, false))
}

func part1(input []string, verbose bool) int64 {
	log = common.GetLogger(verbose)

	equations := parseInput(input)

	var sum int64 = 0
	for _, eq := range equations {
		sum += int64(solveEquation(eq, 2))
	}

	return sum
}

func part2(input []string, verbose bool) int64 {
	log = common.GetLogger(verbose)

	equations := parseInput(input)

	var sum int64 = 0
	for _, eq := range equations {
		sum += int64(solveEquation(eq, 3))
	}

	return sum
}

func solveEquation(eq equation, operators int) int {
	permutations := generatePermutations(len(eq.values)-1, operators)
	log("\nAnswer: ", eq.answer, "\tValues: ", eq.values, "\n")
	for _, perm := range permutations {
		sum := evaluatePermutation(perm, eq.values)
		if sum == eq.answer {
			return sum
		}
	}
	return 0
}

func evaluatePermutation(permutation string, values []int) int {
	sum := values[0]
	for i := 1; i < len(values); i++ {

		val, err := strconv.Atoi(string(permutation[i-1]))
		if err != nil {
			panic(err)
		}
		if val == 0 {
			log("  ", sum, "+", values[i])
			sum = sum + values[i]
		} else if val == 1 {
			log("  ", sum, "*", values[i])
			sum = sum * values[i]
		} else if val == 2 {
			log("  ", sum, "|", values[i])
			sum = concatenate(sum, values[i])
		}
	}
	log(" = ", sum, "\n")
	return sum
}

func generatePermutations(n int, o int) []string {
	combinations := int(math.Pow(float64(o), float64(n)))
	var permutations []string
	for i := 0; i < combinations; i++ {
		permutations = append(permutations, ZeroPad(strconv.FormatInt(int64(i), o), n))
	}
	return permutations
}

func ZeroPad(input string, n int) string {
	if len(input) >= n {
		return input
	}
	return strings.Repeat("0", n-len(input)) + input
}

func concatenate(n1, n2 int) int {
	sum, err := strconv.Atoi(strconv.FormatInt(int64(n1), 10) + strconv.FormatInt(int64(n2), 10))
	if err != nil {
		panic(err)
	}
	return sum
}

func parseInput(input []string) []equation {
	var result []equation
	for _, line := range input {
		parts := strings.Fields(line)
		answer, err := strconv.Atoi(strings.Split(parts[0], ":")[0])
		if err != nil {
			panic(err)
		}
		values := common.SliceAtoi(parts[1:])
		result = append(result, equation{answer, values})
	}
	return result
}
