package main

import (
	"fmt"
	io_helpers "jeremyosz/go-advent-2023/2023/helpers/io"
	string_helpers "jeremyosz/go-advent-2023/2023/helpers/strings"
	"strings"
)

func main() {
	seeds, Maps := parseInput("./input.txt")
	fmt.Println("seeds: ", seeds)

	new := processSeeds(seeds, Maps)
	lowest := getLowest(new)

	// To answer:
	// Why am I off by 1?
	// Why do I get 0 for the lowest location on input

	fmt.Printf(`

	!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

	SEED ANALYSIS RESULTS
	++++++++++++++++++++

	lowest location: %d

	++++++++++++++++++++

	!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

	`, lowest)
}

type Seeds [][2]int
type Map [][]int
type Maps []Map

func parseInput(filepath string) (Seeds, Maps) {

	chunks := strings.Split(io_helpers.ReadFileString(filepath), "\n\n")

	// Seeds is values after "Seeds:"
	seeds := strings.Split(chunks[0], ":")[1]
	seedsInt := string_helpers.ConvertSliceToInts(strings.Split(seeds, " "))

	seedRange := [][2]int{}
	for i := 0; i < len(seedsInt); i += 2 {
		seedRange = append(seedRange, [2]int{seedsInt[i], seedsInt[i] + seedsInt[i+1]})
	}

	rest := chunks[1:]

	maps := Maps{}
	for _, chunk := range rest {
		maps = append(maps, chunkToMap(chunk))
	}

	return seedRange, maps
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getLowest(seeds Seeds) int {
	// fmt.Println("Finding lowest in seeds: ", seeds)
	lowest := seeds[0][0]
	// fmt.Println("lowest: ", lowest)
	for _, seed := range seeds {
		if seed[0] < lowest {
			lowest = seed[0]
			// fmt.Println("lowest: ", lowest)
		}
	}
	return lowest
}

func processSeeds(seeds Seeds, maps Maps) Seeds {
	new := [][2]int{}
	// FOR EACH MAP
	for _, m := range maps {

		// fmt.Println("map: ", m)

		// LOOP THROUGH EACH SEED RANGE
		for _, seed := range seeds {
			start, end := seed[0], seed[1]
			for _, row := range m {
				dstStart, srcStart, rangeLen := row[0], row[1], row[2]

				// Find overlap between seed range and map range
				overlapStart := max(start, srcStart)
				overlapEnd := min(end, srcStart+rangeLen)

				if overlapStart < overlapEnd {
					shift := dstStart - srcStart
					newRange := [2]int{
						overlapStart - shift,
						overlapEnd - shift,
					}
					new = append(new, newRange)

					// Add seeds back to list if there is a gap between the start of the seed range and the start of the overlap
					if overlapStart > start {
						seeds = append(seeds, [2]int{start, overlapStart})
					}
					if end > overlapEnd {
						seeds = append(seeds, [2]int{overlapEnd, end})
					}
					break
				}
				new = append(new, seed)
			}
		}

	}

	return new
}
