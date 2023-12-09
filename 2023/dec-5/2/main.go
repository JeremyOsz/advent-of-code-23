package main

import (
	"fmt"
	io_helpers "jeremyosz/go-advent-2023/2023/helpers/io"
	string_helpers "jeremyosz/go-advent-2023/2023/helpers/strings"
	"strings"
)

func main() {
	seeds, Maps := parseInput("./input.txt")

	seedLocs := []int{}
	for _, seed := range seeds {
		seedLocs = append(seedLocs, findLocation(seed, Maps))
	}

	fmt.Println(seedLocs)

	lowest := getLowest(seedLocs)

	fmt.Printf(`

	!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

	SEED ANALYSIS RESULTS
	++++++++++++++++++++

	lowest location: %d

	++++++++++++++++++++

	!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

	`, lowest)
}

type Seeds []int
type Map [][]int
type Maps []Map

func parseInput(filepath string) ([]int, Maps) {

	chunks := strings.Split(io_helpers.ReadFileString(filepath), "\n\n")

	// Seeds is values after "Seeds:"
	seeds := strings.Split(chunks[0], ":")[1]
	seedsInt := string_helpers.ConvertSliceToInts(strings.Split(seeds, " "))
	rest := chunks[1:]

	maps := Maps{}
	for _, chunk := range rest {
		maps = append(maps, chunkToMap(chunk))
	}

	fmt.Println(seedsInt)
	fmt.Println(maps)

	return seedsInt, maps
}

func chunkToMap(chunk string) [][]int {
	lines := strings.Split(chunk, "\n")[1:]

	currMap := [][]int{}
	// for each line in lines
	for _, line := range lines {
		// split line into ints
		// add ints to map
		intStrings := strings.Split(line, " ")
		ints := string_helpers.ConvertSliceToInts(intStrings)
		currMap = append(currMap, ints)
	}

	return currMap
}

func findLocation(seed int, maps Maps) int {
	currNum := seed

	for _, m := range maps {
		// for each row in map
		for _, row := range m {
			dstStart := row[0]
			srcStart := row[1]
			rangeLen := row[2]
			// if currNum is in range
			if currNum >= srcStart && currNum < srcStart+rangeLen {
				currNum = dstStart + (currNum - srcStart)
				break
			}
		}
	}

	return currNum
}

func getLowest(seeds Seeds) int {
	lowest := seeds[0]
	for _, seed := range seeds {
		if seed < lowest {
			lowest = seed
		}
	}
	return lowest
}
