package main

import (
	"fmt"
	"testing"
)

func TestReadCalibrate(t *testing.T) {
	input := readInput("./calibrate.txt")
	// Check if empty arr
	if len(input) == 0 || input[0] == "" {
		t.Errorf("Expected len(input) to be > 0, got %d", len(input))
	} else {
		t.Logf("\n\nlen(input) is %d", len(input))
	}

}
func TestReadInput(t *testing.T) {
	input := readInput("./input.txt")
	// Check if empty arr
	if len(input) == 0 || input[0] == "" {
		t.Errorf("Expected len(input) to be > 0, got %d", len(input))
	}
}

func TestCalibrate(t *testing.T) {
	input := readInput("./calibrate.txt")
	scratchcards := getScratchcards(input)

	// Card 1 has three winning numbers (48, 83, 17, and 86), so it is worth 8 points.
	// Card 2 has two winning numbers (32 and 61), so it is worth 2 points.
	// Card 3 has two winning numbers (1 and 21), so it is worth 2 points.
	// Card 4 has one winning number (84), so it is worth 1 point.
	// Card 5 has no winning numbers, so it is worth no points.
	// Card 6 has no winning numbers, so it is worth no points.
	if scratchcards[0].points != 8 {
		t.Errorf("Expected scratchcards[0].points to be 8, got %d", scratchcards[0].points)
	}
	if scratchcards[1].points != 2 {
		t.Errorf("Expected scratchcards[1].points to be 2, got %d", scratchcards[1].points)
	}
	if scratchcards[2].points != 2 {
		t.Errorf("Expected scratchcards[2].points to be 2, got %d", scratchcards[2].points)
	}
	if scratchcards[3].points != 1 {
		t.Errorf("Expected scratchcards[3].points to be 1, got %d", scratchcards[3].points)
	}
	if scratchcards[4].points != 0 {
		t.Errorf("Expected scratchcards[4].points to be 0, got %d", scratchcards[4].points)
	}
	if scratchcards[5].points != 0 {
		t.Errorf("Expected scratchcards[5].points to be 0, got %d", scratchcards[5].points)
	}

	sum := sumPoints(scratchcards)
	if sum != 13 {
		t.Errorf("Expected scratchcards to be 13, got %d", sum)
	}

	// Return the sum of all points
	fmt.Printf(`
		!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!


		SCRATCH CARD RESULTS
		++++++++++++

		Total points: %d

		++++++++++++

		!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	`,
		sum)
}
