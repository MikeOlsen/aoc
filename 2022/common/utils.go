package common

import (
	"bufio"
	"log"
	"os"
)

const PATH string = "../../input/2022/"

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
