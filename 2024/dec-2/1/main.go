package main

import (
	"fmt"
	io_helpers "jeremyosz/go-advent-2023/2023/helpers/io"
	string_helpers "jeremyosz/go-advent-2023/2023/helpers/strings"
	"math"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")

	safeReports := testSafety("./input.txt")

	fmt.Printf(`
	------------------------------

	Safe reports: %d

	------------------------------
	`, safeReports)
}

func readInput(filename string) []string {
	// read input from filename and return as []string
	return io_helpers.ReadFileLines(filename)
}

func testSafety(filename string) int {
	// read input from filename
	input := readInput(filename)
	safeReports := 0

	for _, line := range input {
		numbers := string_helpers.ConvertSliceToInts(strings.Split(line, " "))
		isSafe := true

		distances := []int{}

		for i := 1; i < len(numbers); i++ {
			// convert to int
			distance := numbers[i-1] - numbers[i]
			distances = append(distances, distance)
		}

		fmt.Println(distances)

		var negative bool
		for i, distance := range distances {

			if distance == 0 {
				isSafe = false
				break
			}

			// Set direction
			if i == 0 {
				if distance < 0 {
					negative = true
				} else {
					negative = false
				}
			} else {
				// Check if direction changes
				if negative && distance > 0 {
					isSafe = false
					break
				}
				if !negative && distance < 0 {
					isSafe = false
					break
				}
			}

			// Check if distance is within bounds
			if math.Abs(float64(distance)) < 0 {
				isSafe = false
				break
			}

			if math.Abs(float64(distance)) > 3 {
				isSafe = false
				break
			}

		}

		fmt.Println(isSafe)
		if isSafe {
			safeReports++
		}
	}

	return safeReports
}
