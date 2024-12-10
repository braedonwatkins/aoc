package main

import (
	"aoc/utils"
	"testing"
)

func TestOne(t *testing.T) {
	input := utils.FileParse("test.txt")
	result := partOne(input)
	expected := 2

	if result != expected {
		t.Errorf(
			"expected result was %d, but got %d instead",
			expected,
			result,
		)
	}
}

func TestTwo(t *testing.T) {
	input := utils.FileParse("test.txt")
	result := partTwo(input)
	expected := 4

	if result != expected {
		t.Errorf(
			"expected result was %d, but got %d instead",
			expected,
			result,
		)
	}
}
