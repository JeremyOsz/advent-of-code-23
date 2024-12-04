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
	file := readInput("calibrate.txt")
	// 	Output
	// 220 * 5
	// 483 * 4
	// 684 * 3
	// 28 * 2
	// 765 * 1
	expectedWinnings := []int{765, 28 * 2, 684 * 3, 483 * 4, 220 * 5}
	expectedTotalWinnings := 765 + 28*2 + 684*3 + 483*4 + 220*5

	totalWinnings, hands := getWinnings(file)

	// check total winnings
	if totalWinnings != expectedTotalWinnings {
		t.Errorf("Expected total winnings to be %d, got %d", expectedTotalWinnings, totalWinnings)
	}

	// check winnings
	for i, hand := range hands {
		if hand.winnings != expectedWinnings[i] {
			t.Errorf("Expected winnings to be %d, got %d", expectedWinnings[i], hand.winnings)
		}
	}

}
