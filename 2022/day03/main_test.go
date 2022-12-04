package main

import (
	"testing"

	"github.com/MikeOlsen/aoc/2022/common"
)

var test_data = common.LoadInputLines("test", "03")
var prod_data = common.LoadInputLines("prod", "03")

func TestFindCommonItem(t *testing.T) {
	var input = []string{"adAC", "beAC", "cfBC"}

	want := "C"
	got := string(findCommonItem(input, false))
	if got != want {
		t.Errorf("Want: %s, Got: %s", want, got)
	}
}

func TestPart1(t *testing.T) {

	// Test
	want := 157
	got := play_part1(test_data, false)
	if got != want {
		t.Errorf("Want: %d, Got: %d", want, got)
	}

	// Prod
	want = 2342
	got = play_part2(prod_data, false)
	if got != want {
		t.Errorf("Want: %d, Got: %d", want, got)
	}

}

func TestPart2(t *testing.T) {

	// Test
	want := 70
	got := play_part2(test_data, false)
	if got != want {
		t.Errorf("Want: %d, Got: %d", want, got)
	}

	// Prod
	want = 2342
	got = play_part2(prod_data, false)
	if got != want {
		t.Errorf("Want: %d, Got: %d", want, got)
	}

}
