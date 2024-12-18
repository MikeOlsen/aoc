package main

import (
	"fmt"
	"strconv"

	"github.com/MikeOlsen/aoc/2024/common"
)

type Block struct {
	index  int
	length int
}

type Filesystem struct {
	size  int
	files map[int]Block
	space map[int]Block
}

func main() {
	test_data := common.LoadInputLines("test", "09")
	prod_data := common.LoadInputLines("prod", "09")

	fmt.Println("=============== PART 1 ==================")
	fmt.Println("Test:", part1(test_data))
	fmt.Println("Prod:", part1(prod_data))

	fmt.Println("=============== PART 2 ==================")
	fmt.Println("Test:", part2(test_data))
	fmt.Println("Prod:", part2(prod_data))
}

func part1(input []string) int {
	filesystem := parseInput(input[0])
	representation := getRepresentation(filesystem)
	blockCompact(representation)
	return calculateChecksum(representation)
}

func part2(input []string) int {
	filesystem := parseInput(input[0])
	filesystem = fileCompact(filesystem)
	representation := getRepresentation(filesystem)
	return calculateChecksum(representation)
}

func parseInput(input string) Filesystem {
	count := 0
	index := 0
	file := true

	filesystem := Filesystem{
		files: make(map[int]Block),
		space: make(map[int]Block),
	}

	for _, char := range input {
		length := parseBlockLength(char)
		filesystem.size += length
		if file {
			filesystem.files[count] = Block{index, length}
			count++
		} else {
			filesystem.space[count] = Block{index, length}
		}
		index += length
		file = !file
	}
	return filesystem
}

func parseBlockLength(char rune) int {
	if char < '0' || char > '9' {
		panic("Not a digit")
	}
	return int(char - '0')
}

func getRepresentation(filesystem Filesystem) []string {
	representation := make([]string, filesystem.size)
	for i := 0; i < filesystem.size; i++ {
		representation[i] = "."
	}
	for idx, block := range filesystem.files {
		for i := 0; i < block.length; i++ {
			representation[block.index+i] = strconv.Itoa(idx)
		}
	}
	return representation
}

func blockCompact(representation []string) {
	for i := 0; i < findLastFreeSpaceIndex(representation); i++ {
		if representation[i] == "." {
			val := popLastBlock(representation)
			representation[i] = val
		}
	}
}

func fileCompact(filesystem Filesystem) Filesystem {
	for i := len(filesystem.files) - 1; i > 0; i-- {
		for j := 0; j < len(filesystem.space); j++ {
			file := filesystem.files[i]
			space := filesystem.space[j]
			if space.index >= file.index {
				break
			}
			if file.length <= space.length {

				// Move file
				file.index = space.index

				// Shrink space
				space.index += file.length
				space.length -= file.length

				filesystem.files[i] = file
				filesystem.space[j] = space
				break
			}
		}
	}
	return filesystem
}

func popLastBlock(representation []string) string {
	for i := len(representation) - 1; i > 0; i-- {
		if "." != representation[i] {
			value := representation[i]
			representation[i] = "."
			return value
		}
	}
	panic("Not found")
}

func findLastFreeSpaceIndex(representation []string) int {
	skip := false
	for i := len(representation) - 1; i > 0; i-- {
		if "." != representation[i] {
			if !skip {
				skip = true
				continue
			}
			return i
		}
	}
	panic("Not found")
}

func calculateChecksum(representation []string) int {
	sum := 0
	for i, block := range representation {
		if block != "." {
			val, err := strconv.Atoi(block)
			if err != nil {
				panic(err)
			}
			sum += i * val
		}
	}
	return sum
}
