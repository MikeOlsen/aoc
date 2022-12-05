package main

import (
	"testing"

	"github.com/MikeOlsen/aoc/2022/common"
)

var test_data = common.LoadInputLines("test", "04")
var prod_data = common.LoadInputLines("prod", "04")

func TestPart1(t *testing.T) {

	// Test
	want := 2
	got := part1(test_data, 10, false)
	if got != want {
		t.Errorf("Want: %d, Got: %d", want, got)
	}

	// Prod
	want = 487
	got = part1(prod_data, 100, false)
	if got != want {
		t.Errorf("Want: %d, Got: %d", want, got)
	}

}

func TestPart2(t *testing.T) {

	// Test
	want := 4
	got := part2(test_data, 10, false)
	if got != want {
		t.Errorf("Want: %d, Got: %d", want, got)
	}

	// Prod
	want = 849
	got = part2(prod_data, 100, false)
	if got != want {
		t.Errorf("Want: %d, Got: %d", want, got)
	}

}
