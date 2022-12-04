package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/MikeOlsen/aoc/2022/common"
)

type Outcome string

const (
	WINN Outcome = "Win"
	DRAW Outcome = "Draw"
	LOSS Outcome = "Loss"
)

type Score struct {
	outcome     Outcome
	score       int
	description string
}

// First column:   A: Rock, B: Paper,  C: Scissors
// Second column:  X: Rock, Y: Paper,  Z: Scissors
var outcomes_part1 = map[string]Score{
	"A X": {DRAW, 4, "rock       1   +   draw    3   =   4"},
	"A Y": {WINN, 8, "paper      2   +   win     6   =   8"},
	"A Z": {LOSS, 3, "scissors   3   +   loss    0   =   3"},
	"B X": {LOSS, 1, "rock       1   +   loss    0   =   1"},
	"B Y": {DRAW, 5, "paper      2   +   draw    3   =   5"},
	"B Z": {WINN, 9, "scissors   3   +   win     6   =   9"},
	"C X": {WINN, 7, "rock       1   +   win     6   =   7"},
	"C Y": {LOSS, 2, "paper      2   +   loss    0   =   2"},
	"C Z": {DRAW, 6, "scissors   3   +   draw    3   =   6"},
}

// First column:   A: rock, B: paper,  C: scissors
// Second column:  X: loss, Y: draw,   Z: win
var outcomes_part2 = map[string]Score{
	"A X": {LOSS, 3, "Scissors   4   +   Loss    0   =   0"},
	"A Y": {DRAW, 4, "Rock       1   +   Draw    3   =   4"},
	"A Z": {WINN, 8, "Paper      2   +   Win     6   =   8"},
	"B X": {LOSS, 1, "Rock       1   +   Loss    0   =   1"},
	"B Y": {DRAW, 5, "Paper      2   +   Draw    3   =   5"},
	"B Z": {WINN, 9, "Scissors   3   +   Win     6   =   9"},
	"C X": {LOSS, 2, "Paper      2   +   Loss    0   =   2"},
	"C Y": {DRAW, 6, "Scissors   3   +   Draw    3   =   6"},
	"C Z": {WINN, 7, "Rock       1   +   Win     6   =   7"},
}

func main() {
	test_data := common.LoadInput("test", "02")
	prod_data := common.LoadInput("prod", "02")

	fmt.Println("=============== PART 1 ==================")
	fmt.Println("Test:", play(test_data, outcomes_part1, true))
	fmt.Println("Prod:", play(prod_data, outcomes_part1, false))

	fmt.Println("=============== PART 2 ==================")
	fmt.Println("Test:", play(test_data, outcomes_part2, true))
	fmt.Println("Prod:", play(prod_data, outcomes_part2, false))
}

func play(input string, mapping map[string]Score, verbose bool) int {
	score := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		score += mapping[line].score
		if verbose {
			fmt.Println(mapping[line].description)
		}
	}
	if verbose {
		fmt.Println(strings.Repeat("-", 41))
	}
	return score
}
