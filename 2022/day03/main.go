package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/MikeOlsen/aoc/2022/common"
)

const LOWER_CASE_OFFSET = 96
const UPPER_CASE_OFFSET = 38

func main() {
	test_data := common.LoadInputLines("test", "03")
	prod_data := common.LoadInputLines("prod", "03")

	fmt.Println("=============== PART 1 ==================")
	fmt.Println("Test:", play_part1(test_data, false))
	fmt.Println("Prod:", play_part1(prod_data, false))

	fmt.Println("=============== PART 2 ==================")
	fmt.Println("Test:", play_part2(test_data, true))
	fmt.Println("Prod:", play_part2(prod_data, false))
}
func play_part1(input []string, verbose bool) int {
	var errors []byte
	for _, line := range input {
		parts := split(line, 2)
		errors = append(errors, findCommonItem(parts, verbose))
	}
	return calculateSum(errors, verbose)
}

func play_part2(input []string, verbose bool) int {
	chunks := chunk(input, 3)
	var errors []byte
	for _, chunk := range chunks {
		errors = append(errors, findCommonItem(chunk, verbose))
	}
	return calculateSum(errors, verbose)
}

func retrieveError(input []string) byte {
	fmt.Print("Batch:", input)
out:
	for i := 0; i < len(input[0]); i++ {
		for j := 0; j < len(input[1]); j++ {
			if input[0][i] == input[1][j] {
				fmt.Println(" -> ", string(input[0][i]))
				continue out
			}
		}
	}
	log.Fatalf("No error found")
	return 0
}

func findCommonItem(input []string, verbose bool) byte {
	if verbose {
		fmt.Println("Batch:", input)
	}
	for _, item := range input[0] {
		exists := true
		for sack := 1; sack < len(input); sack++ {
			if strings.Contains(input[sack], string(item)) {
				continue
			}
			exists = false
		}
		if exists {
			if verbose {
				fmt.Println("The correct item is:", string(item))
			}
			return byte(item)
		}
	}
	log.Fatalf("No error found")
	return 0
}

func calculateSum(letters []byte, verbose bool) int {
	sum := 0
	for i := 0; i < len(letters); i++ {
		prio := calculatePriority(letters[i])
		if verbose {
			fmt.Printf("%s -> %d\n", string(letters[i]), prio)
		}
		sum += prio
	}
	return sum
}
func calculatePriority(letter byte) int {
	if letter > 96 && letter < 123 {
		return int(letter) - LOWER_CASE_OFFSET
	} else {
		return int(letter) - UPPER_CASE_OFFSET
	}
}

func split(input string, times int) (parts []string) {
	size := len(input) / times
	for i := 0; i < times; i++ {
		parts = append(parts, input[i*size:(i+1)*size])
	}
	return parts
}

func chunk(items []string, size int) (chunks [][]string) {
	for i := 0; i < len(items); i += size {
		end := i + size
		if end > len(items) {
			end = len(items)
		}
		chunks = append(chunks, items[i:end])
	}
	return chunks
}
