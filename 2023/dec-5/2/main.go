package main

import (
	"fmt"
	io_helpers "jeremyosz/go-advent-2023/2023/helpers/io"
	string_helpers "jeremyosz/go-advent-2023/2023/helpers/strings"
	"strings"
)

type Seed struct {
	id          int
	soil        int
	fertilizer  int
	water       int
	light       int
	temperature int
	humidity    int
	location    int
}

type MapRow struct {
	destination_range_index int
	source_range_index      int
	range_length            int
}

type Maps map[string]Map

type Map []MapRow

type Seeds map[int]Seed

func main() {
	chunks := readFileChunks("./input.txt")
	seeds := createSeeds(chunks[0])
	maps := readMaps(chunks, seeds)
	seeds = analyseSeedMap(seeds, maps)
	lowest := findLowestLocation(seeds)

	fmt.Printf(`
	
	!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!


	SEED ANALYSIS RESULTS
	++++++++++++++++++++

	lowest location: %d

	++++++++++++++++++++

	!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	
	`, lowest)
}

func findLowestLocation(seeds Seeds) int {
	// Arbitrary high number lowest will be compared to
	lowestLocation := 999999999999999999
	for _, seed := range seeds {
		if seed.location < lowestLocation {
			lowestLocation = seed.location
		}
	}

	return lowestLocation
}

func analyseSeedMap(seeds Seeds, maps Maps) Seeds {
	fmt.Printf(`
		!!!!!!!!!!!!!
		
		analyseSeedMap: 
		
		seeds
		%v

		maps
		%v

	`,
		seeds, maps)
	soilMap := maps["seed-to-soil"]
	fertilizerMap := maps["soil-to-fertilizer"]
	waterMap := maps["fertilizer-to-water"]
	lightMap := maps["water-to-light"]
	temperatureMap := maps["light-to-temperature"]
	humidityMap := maps["temperature-to-humidity"]
	locationMap := maps["humidity-to-location"]

	var processedSeeds = Seeds{}

	for _, seed := range seeds {

		// 1. Map seed to soil
		seed.soil = readMap(seed.id, &soilMap)
		// 2. Map soil to fertilizer
		seed.fertilizer = readMap(seed.soil, &fertilizerMap)
		// 3. Map fertilizer to water
		seed.water = readMap(seed.fertilizer, &waterMap)
		// 4. Map water to light
		seed.light = readMap(seed.water, &lightMap)
		// 5. Map light to temperature
		seed.temperature = readMap(seed.light, &temperatureMap)
		// 6. Map temperature to humidity
		seed.humidity = readMap(seed.temperature, &humidityMap)
		// 7. Map humidity to location
		seed.location = readMap(seed.humidity, &locationMap)

		// Add seed to processedSeeds
		processedSeeds[seed.id] = seed
	}

	return processedSeeds
}

func readMap(source int, destinationMap *Map) int {
	fmt.Printf(`
	
		============

		Mapping source %d

		Based on map:
		%v

	`,
		source,
		*(destinationMap))
	for _, row := range *(destinationMap) {
		// If within range
		if source >= row.source_range_index && source <= row.source_range_index+row.range_length {
			fmt.Println("Source is within range")

			// Calculate position
			// Range is destination_range_index to destination_range_index + range_length
			// SO
			// position = destination_range_index + (source - source_range_index)
			// Eg Source 79, range_length 48, source_range_index 50, destination_range_index 52
			// Because 79 is within range 52-100, the position is 52 + (79-50) = 81

			position := row.destination_range_index + (source - row.source_range_index) // Calculate position

			return position
		}
	}

	// Should only get here if there are no planters available or if the source is not within range
	fmt.Println("No planters available - returning source")
	return source
}

func readFileChunks(filepath string) []string {
	return strings.Split(io_helpers.ReadFileString(filepath), "\n\n")
}

func readMaps(chunks []string, seeds Seeds) Maps {

	var maps = Maps{}
	for _, chunk := range chunks[1:] {
		processChunk(chunk, seeds, &maps)
	}

	// Return the chunks
	return maps
}

func createSeeds(chunk string) Seeds {
	// Remove 'seeds: ' from the chunk and split into a slice
	numberStrings := strings.Split(chunk[7:], " ")
	numbers := string_helpers.ConvertSliceToInts(numberStrings)

	// convert numbers into pairs of ints
	numberPairs := [][]int{}
	fmt.Println("numbers", numbers)
	for i := 0; i < len(numbers); i += 2 {
		numberPairs = append(numberPairs, []int{numbers[i], numbers[i+1]})
	}
	fmt.Println("numberPairs", numberPairs)

	// Create a map of seeds
	seeds := Seeds{}

	for _, numberPair := range numberPairs {
		// Create a seed
		// fmt.Println("numberPair", numberPair)

		seedIndex := numberPair[0]

		// Loop as many times as numberPair[1]
		for i := 0; i < numberPair[1]; i++ {
			fmt.Println("seedIndex", seedIndex)
			// Create a seed with ID as seedIndex
			seed := Seed{
				id: seedIndex,
			}
			// Add seed to seeds
			seeds[seedIndex] = seed

			// Increment seedIndex
			seedIndex++
		}

		// SET SEED RANGE INSTEAD OF CREATE SEEDS
	}

	return seeds
}

func processChunk(chunk string, seeds Seeds, maps *Maps) {
	// Split the chunk into lines
	lines := strings.Split(chunk, "\n")

	// Remove last 4 chars of first line
	id := lines[0][:len(lines[0])-5]

	// Create an entry in the maps map with the id as the key
	(*maps)[id] = []MapRow{}

	for _, lines := range lines[1:] {
		values := strings.Split(lines, " ")
		destination_range_index := string_helpers.ConvertToInt(values[0])
		source_range_index := string_helpers.ConvertToInt(values[1])
		range_length := string_helpers.ConvertToInt(values[2])

		// add the MapRow to *(maps)[id]
		(*maps)[id] = append((*maps)[id], MapRow{
			destination_range_index,
			source_range_index,
			range_length,
		})
	}

}
