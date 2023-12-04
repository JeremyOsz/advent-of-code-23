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
	points  int
}

func main() {
	input := readInput("./input.txt")
	scratchcards := getScratchcards(input)
	sum := sumPoints(scratchcards)

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

	card.winning = string_helpers.ConvertSliceToInt(
		strings.Split(numberSets[0], " "),
	)
	card.numbers = string_helpers.ConvertSliceToInt(
		strings.Split(numberSets[1], " "),
	)

	card.points = calculatePoints(card)

	fmt.Printf(`

		==========================
		NEW SCRATCH Card
		++++++++++++
		Card ID: %s
		Winning: %v
		Numbers: %v
		Points: %d
		++++++++++++
		==========================

	`,
		card.id,
		card.winning,
		card.numbers,
		card.points,
	)

	return card
}

func calculatePoints(card Card) int {
	points := 0
	// for each Card.number, check if it's in Card.winning
	// if it is, add 1 to point (or double it if it's not the first)
	for _, number := range card.numbers {
		if number_helpers.SliceContainsInt(card.winning, number) {
			fmt.Println("Number is in winning: ", number)
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
	}
	return points
}

func sumPoints(cards []Card) int {
	points := 0
	for _, card := range cards {
		points += card.points
	}
	return points
}
