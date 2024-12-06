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
	result := countXmas("./calibrate.txt")
	if result != 18 {
		t.Errorf("Expected 18, got %d", result)
	}
}

func TestMain(t *testing.T) {
	main()
}
