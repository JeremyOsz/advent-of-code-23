package main

import (
	"fmt"
	number_helpers "jeremyosz/go-advent-2023/2023/helpers/numbers"
	string_helpers "jeremyosz/go-advent-2023/2023/helpers/strings"
	"log"
	"os"
	"strings"
)

type Card struct {
	id      string
	numbers []int
	winning []int
	wins    int
	copies  int
}

func main() {
	input := readInput("./input.txt")
	cards := getScratchcards(input)
	processedCards := returnResults(cards)
	sum := sumCards(processedCards)

	fmt.Println(processedCards[0])
	for _, card := range processedCards {
		// Print how many copies each card has
		fmt.Printf("Card %s has %d copies\n", card.id, card.copies)
	}

	// Return the sum of all points
	fmt.Printf(`
		!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

		SCRATCH CARD RESULTS
		++++++++++++

		Total cards: %d

		++++++++++++

		!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

	`,
		sum)
}

func readInput(filename string) []string {
	// read input from filename and return as []string
	input, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Split the input into lines
	return strings.Split(string(input), "\n")
}

func getScratchcards(input []string) []Card {
	cards := []Card{}

	for _, line := range input {
		if line == "" {
			continue
		}
		cards = append(cards, getCard(line))
	}
	return cards
}

func getCard(line string) Card {
	card := Card{}

	// Card ID is everthing before the first :
	card.id = strings.Split(line, ":")[0]

	// Remove the Card ID from the line
	line = strings.Split(line, ": ")[1]

	// Split line into two numbers based on |
	numberSets := strings.Split(line, "|")

	card.winning = string_helpers.ConvertSliceToInts(
		strings.Split(numberSets[0], " "),
	)
	card.numbers = string_helpers.ConvertSliceToInts(
		strings.Split(numberSets[1], " "),
	)

	card.wins = calculateWins(card)

	card.copies = 1

	fmt.Printf(`

		==========================
		NEW SCRATCH Card
		++++++++++++
		Card ID: %s
		Winning: %v
		Numbers: %v
		Wins: %d
		Copies: %d
		++++++++++++
		==========================

	`,
		card.id,
		card.winning,
		card.numbers,
		card.wins,
		card.copies,
	)

	return card
}

func calculateWins(card Card) int {
	wins := 0
	// for each Card.number, check if it's in Card.winning
	// if it is, add 1 win
	for _, number := range card.numbers {
		if number_helpers.SliceContainsInt(card.winning, number) {
			wins++
		}
	}
	return wins
}

func returnResults(cards []Card) []Card {
	// Loop through each card
	for i, card := range cards {
		// If card has no wins skip
		if card.wins == 0 {
			continue
		}

		// Loop for each copy of the card
		for c := 0; c < card.copies; c++ {
			// if card has 1 win add copy to card+1.copies

			if card.wins > 0 {
				// If card+1 doesn't exist do not add copies
				// For each win, add a copy to the next card
				for w := 0; w < card.wins; w++ {
					if len(cards) < 2 {
						continue
					}
					// add copy to next card - and go one card forward for each win
					copyIndex := i + 1 + w
					// if that card doesn't exist, skip
					if copyIndex > len(cards)-1 {
						continue
					}
					cards[copyIndex].copies++
				}
			}
		}

	}

	return cards
}

func sumCards(cards []Card) int {
	sum := 0
	for _, card := range cards {
		sum += card.copies
	}
	return sum
}
