package main

import (
	"fmt"
	io_helpers "jeremyosz/go-advent-2023/2023/helpers/io"
	"testing"
)

func TestReadCalibrate(t *testing.T) {
	input := io_helpers.ReadFileLines("./calibrate.txt")
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
	similarity := calculateSimilarity("./calibrate.txt")

	fmt.Println(similarity)
	if similarity != 31 {
		t.Errorf("Expected similarity to be 31, got %d", similarity)
	}
}

func TestInput(t *testing.T) {
	distance := calculateSimilarity("./input.txt")

	fmt.Println(distance)
}
