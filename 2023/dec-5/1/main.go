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
	fmt.Println("Hello, World!")
	chunks := readFileChunks("./calibrate.txt")
	seeds := createSeeds(chunks[0])
	maps := readMaps(chunks, seeds)
	analyseSeedMap(seeds, maps)
}

func analyseSeedMap(seeds Seeds, maps Maps) Seeds {
	soilMap := maps["soil"]
	fertilizerMap := maps["fertilizer"]
	waterMap := maps["water"]
	lightMap := maps["light"]
	temperatureMap := maps["temperature"]
	humidityMap := maps["humidity"]
	locationMap := maps["location"]

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
	for _, row := range *(destinationMap) {
		// If within range
		if source >= row.source_range_index && source <= row.source_range_index+row.range_length {
			// save plant position
			plantPosition := row.destination_range_index
			// add 1 to the destination to mark it as planted
			row.destination_range_index++
			// remove 1 from the range to mark one less planter available
			row.range_length--

			return plantPosition
		}
	}

	// Should only get here if there are no planters available or if the source is not within range
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

	// Create a map of seeds
	seeds := Seeds{}
	for _, number := range numbers {
		seeds[number] = Seed{
			id: number,
		}
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
