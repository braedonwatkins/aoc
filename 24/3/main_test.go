package main

import (
	"aoc/utils"
	"testing"
)

func TestOne(t *testing.T) {
	input := utils.FileParser("test.txt")
	expected := 161
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
	input := utils.FileParser("test2.txt")
	expected := 48
	result := partTwo(input)

	if result != expected {
		t.Errorf(
			"expected result was %d, but got %d instead",
			expected,
			result,
		)
	}
}
