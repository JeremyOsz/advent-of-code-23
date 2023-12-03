package main

import (
	"fmt"
	"testing"
)

func TestReadSchematic(t *testing.T) {
	input := readInput("./calibrate.txt")
	result := readSchematic(input)

	// expect result to be a list of ints
	if len(result) < 1 {
		t.Errorf("Expected a list of ints, got %d", result)
	}
}

func TestCalibrate(t *testing.T) {
	input := readInput("./calibrate.txt")
	result := readSchematic(input)
	sum := schematicSum(result)

	// expect sum to be 4361
	if sum != 4361 {
		t.Errorf("Expected 4361, got %d", sum)
	} else {
		fmt.Print(
			"\n\n!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n\n",
			"Expected 4361, got ", sum,
			"\n\nGreat job!",
			"\n\n!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n\n",
		)
	}
}

func TestRealInput(t *testing.T) {
	input := readInput("./input.txt")
	result := readSchematic(input)
	sum := schematicSum(result)

	// expect sum to be 4361
	fmt.Print(
		"\n\n!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n\n",
		"We have a result.... ", sum,
		"\n\nCheck if it works!",
		"\n\n!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n\n",
	)
}
