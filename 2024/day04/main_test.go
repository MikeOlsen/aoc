package main

import (
	"testing"

	"github.com/MikeOlsen/aoc/2024/common"
)

var test_data = common.LoadInputLines("test", "04")
var prod_data = common.LoadInputLines("prod", "04")

func TestCheckDirection(t *testing.T) {
	grid := [][]rune{
		{'X', 'M', 'A', 'S'},
		{'S', 'A', 'M', 'X'},
		{'X', 'X', 'M', 'A'},
		{'A', 'S', 'X', 'M'},
	}

	got := checkDirection(grid, 0, 0, 1, 0, "XMAS")
	if got != true {
		t.Errorf("Want: %t, Got: %t", true, got)
	}

	got = checkDirection(grid, 0, 0, 1, 1, "XAMM")
	if got != true {
		t.Errorf("Want: %t, Got: %t", true, got)
	}
}

func TestCheckCross(t *testing.T) {
	grid := [][]rune{
		{'M', '.', 'S'},
		{'.', 'A', '.'},
		{'M', '.', 'S'},
	}

	got := checkCross(grid, 0, 0, "MAS")
	if got != true {
		t.Errorf("Want: %t, Got: %t", true, got)
	}
}

func TestPart1(t *testing.T) {

	// Test
	want := 18
	got := part1(test_data)
	if got != want {
		t.Errorf("Want: %d, Got: %d", want, got)
	}

	// Prod
	want = 2434
	got = part1(prod_data)
	if got != want {
		t.Errorf("Want: %d, Got: %d", want, got)
	}

}

func TestPart2(t *testing.T) {

	// Test
	want := 9
	got := part2(test_data)
	if got != want {
		t.Errorf("Want: %d, Got: %d", want, got)
	}

	// Prod
	want = 1835
	got = part2(prod_data)
	if got != want {
		t.Errorf("Want: %d, Got: %d", want, got)
	}

}
