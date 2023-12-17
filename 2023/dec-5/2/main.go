package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type ItemRange struct {
	start     int
	end       int
	processed bool
}

type MapFilter struct {
	filters []filter
	chain   *MapFilter
}

type filter struct {
	inputStart int
	inputEnd   int
	outputDiff int
}

func ReadLines(filename string) []string {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(file), "\n")
}

func processSeeds(input []string) int {
	var mapFilters []*MapFilter

	// build the filters
	for i := 2; i < len(input); i++ {
		// Ignore empty lines
		if len(input[i]) == 0 || input[i] == "" {
			continue
		}

		// Get seeds from first line by splitting on :

		// See if first char is a number - this means it is a filter rule
		if unicode.IsDigit(rune(input[i][0])) {

			// Add the filter
			values := strings.Fields(input[i])
			inputStart, _ := strconv.Atoi(values[1])
			filterRange, _ := strconv.Atoi(values[2])
			outputStart, _ := strconv.Atoi(values[0])

			// Convert filter to a range and add it to the current filter
			currentFilter := mapFilters[len(mapFilters)-1]
			currentFilter.filters = append(currentFilter.filters, filter{
				inputStart: inputStart,
				inputEnd:   inputStart + filterRange - 1,
				outputDiff: outputStart - inputStart,
			})
		} else {
			// It is a title - so create a new map filter
			newMapFilter := &MapFilter{}
			if len(mapFilters) > 0 {
				// sets the chain to the last filter so we can chain them together
				mapFilters[len(mapFilters)-1].chain = newMapFilter
			}
			mapFilters = append(mapFilters, newMapFilter)
		}
	}

	// Get maximum possible int32
	lowestNum := math.MaxInt64
	startMapFilter := mapFilters[0]
	seedsParts := strings.Fields(strings.Split(input[0], ":")[1])
	seedRanges := []*ItemRange{}

	// Build the ranges
	for i := 0; i < len(seedsParts)-1; i++ {
		seedVal, _ := strconv.Atoi(seedsParts[i])
		rangeVal, _ := strconv.Atoi(seedsParts[i+1])

		seedRanges = append(seedRanges, &ItemRange{
			start: seedVal,
			end:   seedVal + rangeVal - 1,
		})
		i++
	}

	result := startMapFilter.Calculate(seedRanges)

	for _, seedRange := range result {
		if seedRange.start < lowestNum {
			lowestNum = seedRange.start
		}
	}

	return lowestNum
}

// Calulate method
func (m *MapFilter) Calculate(input []*ItemRange) (outputRanges []*ItemRange) {

	for {
		allProcessed := true
		for _, inputRange := range input {
			if inputRange.processed {
				continue
			}
			allProcessed = false

			// Check if range is in any of the filters
			for _, filter := range m.filters {

				// Check exact fit first
				if inputRange.start >= filter.inputStart &&
					inputRange.end <= filter.inputEnd { // all below end
					// It is in the filter - so add the output range
					outputRanges = append(outputRanges, &ItemRange{
						start: inputRange.start + filter.outputDiff,
						end:   inputRange.end + filter.outputDiff,
					})
					inputRange.processed = true
					break
				}

				// Else Check overlap
				if inputRange.start >= filter.inputStart &&
					inputRange.start <= filter.inputEnd { // some below end
					outputRange := ItemRange{
						start: inputRange.start + filter.outputDiff,
						end:   filter.inputEnd + filter.outputDiff,
					}
					inputRange.processed = true

					// Split output ranges and add them to the output
					outputRanges = append(outputRanges, &outputRange)

					// Set remaining range FROM the end of the filter TO the end of the input range
					remainingRange := ItemRange{
						start: filter.inputEnd + 1,
						end:   inputRange.end,
					}

					// Add the remaining range back to the input
					input = append(input, &remainingRange)
					break
				}
			}
			if !inputRange.processed {
				inputRange.processed = true
				// If we're here it hasn't matched any filters to add it to the output filter directly and mark as complete
				outputRanges = append(outputRanges, &ItemRange{
					start:     inputRange.start,
					end:       inputRange.end,
					processed: false,
				})
			}
		}
		if allProcessed {
			break
		}
	}
	// now pass to the chain
	if m.chain != nil {
		outputRanges = m.chain.Calculate(outputRanges)
	}
	return outputRanges
}

func main() {
	start := time.Now()

	// Your existing code
	input := ReadLines("./input.txt")
	lowest := processSeeds(input)
	println(lowest)

	elapsed := time.Since(start)
	log.Printf("main took %s", elapsed)
}
