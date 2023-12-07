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

	// Idea 1 - find distances from a time of 1ms+
	// at point distances are diminishing returns continue until below record distance

	distances := [][2]int{}
	passedRecord := false
	for buttonTime := 1; buttonTime <= raceTime; buttonTime++ {
		if passedRecord && buttonTime > record {
			break
		}

		timeLeft := raceTime - buttonTime
		distance := timeLeft * buttonTime
		distances = append(distances, [2]int{buttonTime, distance})

		if distance > record {
			passedRecord = true
		}
	}

	// Get number of distances higher than record
	fmt.Println("distances", distances)
	updatedDistances := [][2]int{}
	for _, distance := range distances {
		if distance[1] > record {
			updatedDistances = append(updatedDistances, distance)
		}
	}

	// Get the min and max times
	min := updatedDistances[0][0]
	max := updatedDistances[0][0]
	for _, distance := range updatedDistances {
		if distance[0] < min {
			min = distance[0]
		}
		if distance[0] > max {
			max = distance[0]
		}
	}

	fmt.Println("min", min)
	fmt.Println("max", max)

	return [2]int{min, max}

	// Idea 2 - Start at middle time and work outwards
	// stopping when the distance is below record distance
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
