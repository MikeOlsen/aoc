package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/MikeOlsen/aoc/2022/common"
)

func main() {
	test_data := common.LoadInputLines("test", "04")
	prod_data := common.LoadInputLines("prod", "04")

	fmt.Println("=============== PART 1 ==================")
	fmt.Println("Test:", part1(test_data, 10, true))
	fmt.Println("Prod:", part1(prod_data, 100, false))

	fmt.Println("=============== PART 2 ==================")
	fmt.Println("Test:", part2(test_data, 10, true))
	fmt.Println("Prod:", part2(prod_data, 100, false))
}

func part1(input []string, size int, verbose bool) (count int) {

	var sections [][][]bool
	for _, line := range input {
		sections = append(sections, createSection(line, size))
	}

	for _, section := range sections {
		if contains(section[0], section[1]) {
			count++
		} else if contains(section[1], section[0]) {
			count++
		}
	}
	return count
}

func part2(input []string, size int, verbose bool) (count int) {

	var sections [][][]bool
	for _, line := range input {
		sections = append(sections, createSection(line, size))
	}

	for _, section := range sections {
		if overlaps(section[0], section[1]) {
			count++
		}
	}
	return count
}

// Creates a section pair
func createSection(line string, size int) [][]bool {
	elfs := strings.Split(line, ",")

	// Create a section for each elf
	section := make([][]bool, 2)
	for i := range section {
		section[i] = make([]bool, size)
	}

	// Populate sections with assignments
	for i, elf := range elfs {
		parts := strings.Split(elf, "-")

		start, err1 := strconv.Atoi(parts[0])
		end, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil {
			log.Fatalf("Failed to parse integers")
		}

		for j := start - 1; j < end; j++ {
			section[i][j] = true
		}
	}
	return section
}

func overlaps(section1 []bool, section2 []bool) bool {
	for i := range section2 {
		if section1[i] && section2[i] {
			return true
		}
	}
	return false
}

func contains(section1 []bool, section2 []bool) bool {
	for i, v := range section2 {
		if v == true && section1[i] != section2[i] {
			return false
		}
	}
	return true
}
