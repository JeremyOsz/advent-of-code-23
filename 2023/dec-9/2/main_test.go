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

	nextValues, sum := analyseInput("./calibrate.txt")

	expectedValues := []int{3, 4, 5}
	expectedSum := 12

	if expectedSum != sum {
		t.Errorf("Expected %d, got %d", expectedSum, sum)
	}

	// Check if values are in expectedValues
	for i, value := range nextValues {
		if value != expectedValues[i] {
			t.Errorf("Expected %d, got %d", expectedValues[i], value)
		}
	}
}
