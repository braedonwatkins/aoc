package main

import (
	"aoc/utils"
	"testing"
)

func TestOne(t *testing.T) {
	input := utils.FileParser("input.txt")
	expected := 1873376
	result := partOne(input)

	if result != expected {
		t.Errorf(
			"expected result was %d, but got %d instead",
			expected,
			result,
		)
	}
}

func TestTwo(t *testing.T) {
	input := utils.FileParser("input.txt")
	expected := 18997088
	result := partTwo(input)

	if result != expected {
		t.Errorf(
			"expected result was %d, but got %d instead",
			expected,
			result,
		)
	}
}
