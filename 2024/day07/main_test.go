package main

import (
	"reflect"
	"testing"

	"github.com/MikeOlsen/aoc/2024/common"
)

var test_data = common.LoadInputLines("test", "04")
var prod_data = common.LoadInputLines("prod", "04")

func TestGeneratePermutations(t *testing.T) {

	operators := 2

	want := []string{"00", "01", "10", "11"}
	got := generatePermutations(2, operators)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want: %s, Got: %s", want, got)
	}

	want = []string{"000", "001", "010", "011", "100", "101", "110", "111"}
	got = generatePermutations(3, operators)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want: %s, Got: %s", want, got)
	}

	want = []string{"0000", "0001", "0010", "0011", "0100", "0101", "0110", "0111", "1000", "1001", "1010", "1011", "1100", "1101", "1110", "1111"}
	got = generatePermutations(4, operators)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want: %s, Got: %s", want, got)
	}

	operators = 3

	want = []string{"00", "01", "02", "10", "11", "12", "20", "21", "22"}
	got = generatePermutations(2, operators)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want: %s, Got: %s", want, got)
	}

}
