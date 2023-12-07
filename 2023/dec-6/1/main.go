package main

import (
	"fmt"
	io_helpers "jeremyosz/go-advent-2023/2023/helpers/io"
	string_helpers "jeremyosz/go-advent-2023/2023/helpers/strings"
	"log"
	"os"
	"strings"
)

type Races []Race

// Race is a tuple of time and distance - always 2 ints
type Race [2]int

func main() {
	input := io_helpers.ReadFileLines("./input.txt")
	races := getRaceData(input)
	ranges := getRecordRanges(races)
	rangeLengths := getRangeLengths(ranges)
	fmt.Println("rangeLengths", rangeLengths)
	totalMarginOfError := getMargin(rangeLengths)

	fmt.Printf(`

	!!!!!!!!!!!!!!!

	Total margin of error: %d

	!!!!!!!!!!!!!!!

	`, totalMarginOfError)
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

// Read input that looks like this:
// Time:      7  15   30
// Distance:  9  40  200
// Get slice of tuples like this:
// [[7, 9], [15, 40], [30, 200]]

func getRaceData(input []string) Races {

	raceTuple := Races{}

	// Create slices for time and distance
	// ASSUME: input is always 2 lines
	time := string_helpers.ConvertSliceToInts(strings.Split(input[0], " ")[1:])
	distance := string_helpers.ConvertSliceToInts(strings.Split(input[1], " ")[1:])

	// Create a tuple for each time and distance
	for i := 0; i < len(time); i++ {
		raceTuple = append(raceTuple, Race{time[i], distance[i]})
	}

	fmt.Println("raceTuple", raceTuple)
	return raceTuple
}

// TODO: Do not find all results - just find range
func getRecordRange(race Race) [2]int {
	// Time is race[0] and record distance is race[1]
	// Find possible record distances

	raceTime := race[0]
	record := race[1]

	right := [][2]int{}
	for buttonTime := raceTime / 2; buttonTime < raceTime; buttonTime++ {
		timeLeft := raceTime - buttonTime
		distance := timeLeft * buttonTime
		if distance > record {
			right = append(right, [2]int{buttonTime, distance})
		}
	}

	left := [][2]int{}
	for buttonTime := (raceTime / 2) - 1; buttonTime >= 0; buttonTime-- {
		timeLeft := raceTime - buttonTime
		distance := timeLeft * buttonTime
		if distance > record {
			left = append(left, [2]int{buttonTime, distance})
		}
	}

	fmt.Println("left", left)
	fmt.Println("right", right)

	// combine left and right

	// Get the min and max times
	return [2]int{left[len(left)-1][0], right[len(right)-1][0]}
}

// Get ranges for all races
func getRecordRanges(races Races) [][2]int {
	ranges := [][2]int{}

	// for each race apply getRecordRange
	for _, r := range races {
		ranges = append(ranges, getRecordRange(r))
	}

	return ranges
}

func getRangeLengths(ranges [][2]int) []int {
	lengths := []int{}
	for _, r := range ranges {
		lengths = append(lengths, r[1]-r[0]+1)
	}
	return lengths
}

func getMargin(winCases []int) int {
	accum := 1
	for _, winCase := range winCases {
		accum *= winCase
	}
	return accum
}
