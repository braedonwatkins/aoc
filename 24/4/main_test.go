package main

import (
	"aoc/utils"
	"testing"
)

func TestOne(t *testing.T) {
	input := utils.FileParse("test.txt")
	expected := 18
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
	input := utils.FileParse("test.txt")
	expected := 9
	result := partTwo(input)

	if result != expected {
		t.Errorf(
			"expected result was %d, but got %d instead",
			expected,
			result,
		)
	}
}