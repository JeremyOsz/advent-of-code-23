package main

import (
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

	// Card 1 has 4 wins
	if scratchcards[0].wins != 4 {
		t.Errorf("Expected scratchcards[0].wins to be 4, got %d", scratchcards[0].wins)
	}
	// Card 2 has 2 wins
	if scratchcards[1].wins != 2 {
		t.Errorf("Expected scratchcards[1].wins to be 2, got %d", scratchcards[1].wins)
	}
	// Card 3 has 2 wins
	if scratchcards[2].wins != 2 {
		t.Errorf("Expected scratchcards[2].wins to be 2, got %d", scratchcards[2].wins)
	}
	// Card 4 has 1 win
	if scratchcards[3].wins != 1 {
		t.Errorf("Expected scratchcards[3].wins to be 1, got %d", scratchcards[3].wins)
	}
	// Card 5 has 0 wins
	if scratchcards[4].wins != 0 {
		t.Errorf("Expected scratchcards[4].wins to be 0, got %d", scratchcards[4].wins)
	}
	// Card 6 has 0 wins
	if scratchcards[5].wins != 0 {
		t.Errorf("Expected scratchcards[5].wins to be 0, got %d", scratchcards[5].wins)
	}

	processedCards := returnResults(scratchcards)

	// Card 1 has 1 copy
	if processedCards[0].copies != 1 {
		t.Errorf("Expected processedCards[0].copies to be 1, got %d", processedCards[0].copies)
	}

	// Card 2 has 2 copies
	if processedCards[1].copies != 2 {
		t.Errorf("Expected processedCards[1].copies to be 2, got %d", processedCards[1].copies)
	}

	// Card 3 has 4 copies
	if processedCards[2].copies != 4 {
		t.Errorf("Expected processedCards[2].copies to be 4, got %d", processedCards[2].copies)
	}

	// Card 4 has 8 copies
	if processedCards[3].copies != 8 {
		t.Errorf("Expected processedCards[3].copies to be 8, got %d", processedCards[3].copies)
	}

	// Card 5 has 14 copies
	if processedCards[4].copies != 14 {
		t.Errorf("Expected processedCards[4].copies to be 14, got %d", processedCards[4].copies)
	}

	// Card 6 has 1 copy
	if processedCards[5].copies != 1 {
		t.Errorf("Expected processedCards[5].copies to be 1, got %d", processedCards[5].copies)
	}

	sumCards := sumCards(processedCards)

	// Sum of cards is 30
	if sumCards != 30 {
		t.Errorf("Expected sumCards to be 30, got %d", sumCards)
	}
}
