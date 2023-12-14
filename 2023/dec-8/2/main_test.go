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
	calibrate1 := getSteps("./calibrate.txt")

	if calibrate1 != 2 {
		t.Errorf("Expected calibrate1 to be 2, got %d", calibrate1)
	}

	calibrate2 := getSteps("./calibrate2.txt")
	if calibrate2 != 6 {
		t.Errorf("Expected calibrate2 to be 6, got %d", calibrate2)
	}

	calibrate3 := getSteps("./calibrate3.txt")
	if calibrate3 != 6 {
		t.Errorf("Expected calibrate3 to be 6, got %d", calibrate3)
	}
}
