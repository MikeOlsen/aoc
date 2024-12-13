package common

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const PATH string = "../../input/2024/"

func LoadInput(env string, day string) string {
	data, err := os.ReadFile(PATH + env + "/" + day + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func LoadInputLines(env string, day string) (lines []string) {
	file, err := os.Open(PATH + env + "/" + day + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func SliceAtoi(input []string) []int {
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
