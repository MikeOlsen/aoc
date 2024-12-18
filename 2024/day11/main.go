package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MikeOlsen/aoc/2024/common"
)

func main() {
	test_data := common.LoadInputLines("test", "11")
	prod_data := common.LoadInputLines("prod", "11")

	fmt.Println("=============== PART 1 ==================")
	fmt.Println("Test:", countStones(test_data, 25))
	fmt.Println("Prod:", countStones(prod_data, 25))

	fmt.Println("=============== PART 2 ==================")
	fmt.Println("Test:", countStones(test_data, 75))
	fmt.Println("Prod:", countStones(prod_data, 75))
}

func countStones(input []string, n int) int {
	parts := strings.Fields(input[0])
	var arrangement []int
	for _, part := range parts {
		stone, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
		arrangement = append(arrangement, stone)
	}

	freq := make(map[int]int)
	for _, stone := range arrangement {
		freq[stone]++
	}

	for i := 0; i < n; i++ {
		freq = blink(freq)
	}

	score := 0
	for _, val := range freq {
		score += val
	}

	return score
}

func blink(frequencies map[int]int) map[int]int {
	result := make(map[int]int)
	for stone, count := range frequencies {
		switch {

		case stone == 0:
			result[1] += count

		case countDigits(stone)%2 == 0:
			left, right := splitStone(stone)
			result[left] += count
			result[right] += count

		default:
			result[stone*2024] += count
		}
	}
	return result
}

func splitStone(stone int) (int, int) {
	number := strconv.Itoa(stone)
	mid := len(number) / 2
	left, _ := strconv.Atoi(number[:mid])
	right, _ := strconv.Atoi(number[mid:])
	return left, right
}

func countDigits(number int) int {
	if number == 0 {
		return 1
	}
	count := 0
	for number != 0 {
		number /= 10
		count++
	}
	return count
}
