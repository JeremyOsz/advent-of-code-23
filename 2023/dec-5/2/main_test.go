package main

import (
	io_helpers "jeremyosz/go-advent-2023/2023/helpers/io"
	"testing"
)

type seedTest struct {
	id          int
	soil        int
	fertilizer  int
	water       int
	light       int
	temperature int
	humidity    int
}

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

	seeds, Maps := parseInput("./input.txt")

	seedLocs := [][2]int{}
	for _, seed := range seeds {
		seedLocs = append(seedLocs, findLocation(seed, Maps))
	}

	lowest := getLowest(seedLocs)

	// Expect lowest to be 35
	if lowest != 46 {
		t.Errorf("Expected lowest to be 35, got %d", lowest)
	}
}
