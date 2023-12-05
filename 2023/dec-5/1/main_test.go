package main

import (
	io_helpers "jeremyosz/go-advent-2023/2023/helpers/io"
	"testing"
)

type seedTest struct {
	id          int
	soil        int
	fertilizer  int
	water       int
	light       int
	temperature int
	humidity    int
}

func TestReadCalibrate(t *testing.T) {
	input := io_helpers.ReadFileLines("./calibrate.txt")
	// Check if empty arr
	if len(input) == 0 || input[0] == "" {
		t.Errorf("Expected len(input) to be > 0, got %d", len(input))
	}

}
func TestReadInput(t *testing.T) {
	input := io_helpers.ReadFileLines("./input.txt")
	// Check if empty arr
	if len(input) == 0 || input[0] == "" {
		t.Errorf("Expected len(input) to be > 0, got %d", len(input))
	}
}

func TestCalibrate(t *testing.T) {

	tests := []seedTest{
		{79, 81, 81, 81, 74, 78, 78},
		{14, 14, 53, 49, 42, 42, 43},
		{55, 57, 57, 53, 46, 82, 82},
		{13, 13, 52, 41, 34, 34, 35},
		// Add more tests as needed
	}

	chunks := readFileChunks("./calibrate.txt")
	seeds := createSeeds(chunks[0])
	maps := readMaps(chunks, seeds)
	seeds = analyseSeedMap(seeds, maps)

	for _, test := range tests {
		if seeds[test.id].soil != test.soil {
			t.Errorf("Expected seeds[%d].soil to be %d, got %d", test.id, test.soil, seeds[test.id].soil)
		}
		if seeds[test.id].fertilizer != test.fertilizer {
			t.Errorf("Expected seeds[%d].fertilizer to be %d, got %d", test.id, test.fertilizer, seeds[test.id].fertilizer)
		}
		if seeds[test.id].water != test.water {
			t.Errorf("Expected seeds[%d].water to be %d, got %d", test.id, test.water, seeds[test.id].water)
		}
		if seeds[test.id].light != test.light {
			t.Errorf("Expected seeds[%d].light to be %d, got %d", test.id, test.light, seeds[test.id].light)
		}
		if seeds[test.id].temperature != test.temperature {
			t.Errorf("Expected seeds[%d].temperature to be %d, got %d", test.id, test.temperature, seeds[test.id].temperature)
		}
		if seeds[test.id].humidity != test.humidity {
			t.Errorf("Expected seeds[%d].humidity to be %d, got %d", test.id, test.humidity, seeds[test.id].humidity)
		}
	}

	lowest := findLowestLocation(seeds)

	// Expect lowest to be 35
	if lowest != 35 {
		t.Errorf("Expected lowest to be 35, got %d", lowest)
	}
}
