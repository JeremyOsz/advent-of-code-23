package main

import (
	"fmt"
	io_helpers "jeremyosz/go-advent-2023/2023/helpers/io"
	"testing"
)

func TestReadCalibrate(t *testing.T) {
	input := io_helpers.ReadFileLines("./input.txt")
	// Check if empty arr
	if len(input) == 0 || input[0] == "" {
		t.Errorf("Expected len(input) to be > 0, got %d", len(input))
	}

	fmt.Println(input)

}
func TestReadInput(t *testing.T) {
	input := io_helpers.ReadFileLines("./input.txt")
	// Check if empty arr
	if len(input) == 0 || input[0] == "" {
		t.Errorf("Expected len(input) to be > 0, got %d", len(input))
	}
}

func TestCalibrate(t *testing.T) {
	input := io_helpers.ReadFileLines("./calibrate.txt")
	races := getRaceData(input)

	fmt.Println(races)

	// Race 1 - Range should be 2-5
	recordRange := getRecordRange(races[0])
	if recordRange[0] != 2 || recordRange[1] != 5 {
		t.Errorf("Expected recordRange to be [2, 5], got %d", recordRange)
	}

	// Race 2 - Range should be 4-11
	recordRange = getRecordRange(races[1])
	if recordRange[0] != 4 || recordRange[1] != 11 {
		t.Errorf("Expected recordRange to be [4, 11], got %d", recordRange)
	}

	// Race 3 - Range should be 11-19
	recordRange = getRecordRange(races[2])
	if recordRange[0] != 11 || recordRange[1] != 19 {
		t.Error("Expected recordRange to be [11, 19], got", recordRange)
	}

	// Total margin of error is 4*8*9 = 288
	ranges := getRecordRanges(races)
	rangeLengths := getRangeLengths(ranges)
	fmt.Println("rangeLengths", rangeLengths)
	totalMarginOfError := getMargin(rangeLengths)

	if totalMarginOfError != 288 {
		t.Errorf("Expected totalMarginOfError to be 288, got %d", totalMarginOfError)
	}
}
