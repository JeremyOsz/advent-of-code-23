package main

import (
	"fmt"
	io_helpers "jeremyosz/go-advent-2023/2023/helpers/io"
	string_helpers "jeremyosz/go-advent-2023/2023/helpers/strings"
	"math"
	"strings"
)

func main() {

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
		if isSafe(numbers) {
			safeReports++
		} else {
			for i := 1; i < len(numbers); i++ {
				numbersExcluded := append(numbers[:i], numbers[i+1:]...)
				if isSafe(numbersExcluded) {
					safeReports++
				}
			}
		}
	}
	return safeReports
}

func isSafe(numbers []int) bool {
	for i := 1; i < len(numbers)-1; i++ {

		a, b := numbers[i], numbers[i+1]
		if math.Abs(float64(a-b)) > 3 {
			return false
		}

		if i == len(numbers)-2 {
			continue
		}

		c := numbers[i+2]

		if !((a > b && b > c) || (a < b && b < c)) {
			return false
		}
	}
	return true
}
