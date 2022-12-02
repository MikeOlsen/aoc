package main

import (
	"testing"

	"github.com/MikeOlsen/aoc/2022/common"
)

var test_data = common.LoadInput("test", "01")
var prod_data = common.LoadInput("prod", "01")

func TestPart1(t *testing.T) {

	// Test
	want := 24000
	got := findBiggestElves(test_data, 1, false)
	if got != want {
		t.Errorf("Want: %d, Got: %d", want, got)
	}

	// Prod
	want = 66306
	got = findBiggestElves(prod_data, 1, false)
	if got != want {
		t.Errorf("Want: %d, Got: %d", want, got)
	}

}

func TestPart2(t *testing.T) {

	// Test
	want := 45000
	got := findBiggestElves(test_data, 3, false)
	if got != want {
		t.Errorf("Want: %d, Got: %d", want, got)
	}

	// Prod
	want = 195292
	got = findBiggestElves(prod_data, 3, false)
	if got != want {
		t.Errorf("Want: %d, Got: %d", want, got)
	}

}
