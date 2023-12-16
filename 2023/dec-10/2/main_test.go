package main

import (
	io_helpers "jeremyosz/go-advent-2023/2023/helpers/io"
	"testing"
)

func TestReadCalibrate(t *testing.T) {
	input := io_helpers.ReadFileLines("./input.txt")
	// Check if empty arr
	if len(input) == 0 || input[0] == "" {
		t.Errorf("Expected len(input) to be > 0, got %d", len(input))
	}

}
func TestReadInput(t *testing.T) {
	input := io_helpers.ReadFileLines("./input.txt")
	// Check if empty arr
	if len(input) == 0 || input[0] == "" {
		t.Errorf("Expected len(input) to be > 0, got %d", len(input))
	}
}

func TestCalibrate(t *testing.T) {
	input1 := readInput("./calibrate.txt")
	test1 := getTilesEnclosed(input1)

	if test1 != 4 {
		t.Errorf("Expected test1 to be 4, got %d", test1)
	}

	input2 := readInput("./calibrate2.txt")
	test2 := getTilesEnclosed(input2)

	if test2 != 8 {
		t.Errorf("Expected test2 to be 8, got %d", test2)
	}

}
