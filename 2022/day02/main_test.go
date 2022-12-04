package main

import (
	"testing"

	"github.com/MikeOlsen/aoc/2022/common"
)

var test_data = common.LoadInput("test", "02")
var prod_data = common.LoadInput("prod", "02")

func TestPart1(t *testing.T) {

	// Test
	want := 15
	got := play(test_data, outcomes_part1, false)
	if got != want {
		t.Errorf("Want: %d, Got: %d", want, got)
	}

	// Prod
	want = 14163
	got = play(prod_data, outcomes_part1, false)
	if got != want {
		t.Errorf("Want: %d, Got: %d", want, got)
	}

}

func TestPart2(t *testing.T) {

	// Test
	want := 12
	got := play(test_data, outcomes_part2, false)
	if got != want {
		t.Errorf("Want: %d, Got: %d", want, got)
	}

	// Prod
	want = 12091
	got = play(prod_data, outcomes_part2, false)
	if got != want {
		t.Errorf("Want: %d, Got: %d", want, got)
	}

}
