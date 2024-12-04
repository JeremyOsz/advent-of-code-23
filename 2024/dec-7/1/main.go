package main

import (
	"fmt"
	string_helpers "jeremyosz/go-advent-2023/2023/helpers/strings"
	"log"
	"os"
	"sort"
	"strings"
)

// Map of HandTypes from best to worst
type HandType int

const (
	FiveOfAKind HandType = iota + 1
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

type Card rune

var cardRanks = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

type Hand struct {
	cards      []rune
	cardGroups [][]rune
	bet        int
	handType   HandType
	high       rune
	rank       int
	winnings   int
}

func main() {
	hands := readInput("input.txt")

	// Get winnings
	totalWinnings, hands := getWinnings(hands)

	// Print hands
	printHands(hands)

	fmt.Println("Total Winnings: ", totalWinnings)

}

func printHands(hands []Hand) {
	for _, hand := range hands {
		fmt.Println("===========================")
		fmt.Println("Hand: ", string(hand.cards))
		fmt.Println("Hand Groups: ", hand.cardGroups)
		fmt.Println("Bet: ", hand.bet)
		fmt.Println("Hand Type: ", hand.handType)
		fmt.Println("High Card: ", string(hand.high))
		fmt.Println("Rank: ", hand.rank)
		fmt.Println("Winnings: ", hand.winnings)
		fmt.Println("===========================")
	}

}

func getHandTypes(hands []Hand) []Hand {
	for i, hand := range hands {
		hand = getHandType(hand)
		hands[i] = hand
	}
	return hands
}

func orderHands(hands []Hand) []Hand {
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType == hands[j].handType {
			return compareHighCards(hands[i], hands[j])
		}
		return hands[i].handType > hands[j].handType
	})
	return hands
}

func compareHighCards(a Hand, b Hand) bool {

	for i := range a.cards {
		if a.cards[i] != b.cards[i] {
			return cardRanks[a.cards[i]] < cardRanks[b.cards[i]]
		}
	}

	return false
}

func calculateWinnings(hands []Hand) (int, []Hand) {
	totalWinnings := 0
	for i := range hands {
		// fmt.Println("Setting winnings for hand: ", string(hands[i].cards))
		hands[i].rank = i + 1
		hands[i].winnings = (i + 1) * hands[i].bet
		// fmt.Println("Winnings: ", hands[i].winnings)
		totalWinnings += hands[i].winnings
	}
	return totalWinnings, hands
}

func getWinnings(hands []Hand) (int, []Hand) {
	hands = getHandTypes(hands)
	hands = orderHands(hands)
	return calculateWinnings(hands)
}

func readInput(filename string) []Hand {
	// read input from filename and return as []string
	input, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Split the input into lines
	lines := strings.Split(string(input), "\n")

	// Split each line into 2 parts by the first space
	hands := []Hand{}
	for _, line := range lines {
		x := strings.Split(line, " ")
		bid := string_helpers.ConvertToInt(x[1])
		hand := Hand{
			cards: []rune(x[0]),
			bet:   bid,
		}
		hands = append(hands, hand)
	}

	// Split the input into lines
	return hands
}

func getHandType(hand Hand) Hand {
	// Return the hand type of a given hand
	// 1. Five of a Kind
	// 2. Four of a Kind
	// 3. Full House
	// 4. Three of a Kind
	// 5. Two Pair
	// 6. One Pair
	// 7. High Card

	// Group like cards together - [['A',2], ['Q',2], ['J',1]]
	groups := [][]rune{}
	high := '2'
	for _, card := range hand.cards {

		// Get high card
		if cardRanks[card] > cardRanks[high] {
			high = card
		}

		// if groups is not empty
		if len(groups) > 0 {
			// Check if card is in groups
			for i, group := range groups {
				if group[0] == card {
					groups[i][1] += 1
					break
				}
			}
		}
		// If groups is empty or card is not in groups
		// Add card to groups
		groups = append(groups, []rune{card, 1})
	}

	hand.cardGroups = groups

	// Sort groups by count in descending order
	sort.Slice(groups, func(i, j int) bool {
		return groups[i][1] > groups[j][1]
	})

	// fmt.Println("groups: ", groups)

	// Set high card
	hand.high = high

	switch {
	case groups[0][1] == 5:
		hand.handType = FiveOfAKind
	case groups[0][1] == 4:
		hand.handType = FourOfAKind
	case groups[0][1] == 3 && groups[1][1] == 2:
		hand.handType = FullHouse
	case groups[0][1] == 3:
		hand.handType = ThreeOfAKind
	case groups[0][1] == 2 && groups[1][1] == 2:
		hand.handType = TwoPair
	case groups[0][1] == 2:
		hand.handType = OnePair
	case groups[0][1] == 1:
		hand.handType = HighCard
	}

	return hand
}
