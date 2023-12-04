package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	round2()
}

func round2() int {
	input := readInput("./input.txt")

	total := 0

	// loop through lines
	for r, row := range input {
		for c, char := range row {
			// if char is a digit or .
			if char != '*' {
				continue
			}

			coordSet := make(map[[2]int]bool)

			// if char is *
			for _, x := range []int{r - 1, r, r + 1} {
				for _, y := range []int{c - 1, c, c + 1} {
					// Do something with cr and cc
					if isOutOfBounds(x, y, input) || !isDigit(input[x][y]) {
						continue
					}
					// Print digit
					fmt.Println("digit: ", string(char))
					//
					for y > 0 && isDigit(input[x][y-1]) {
						y--
					}
					coordSet[[2]int{x, y}] = true
				}
			}

			if len(coordSet) < 2 {
				continue
			}

			// loop through coordSet
			numbers := []int{}

			for coord := range coordSet {
				numberString := ""
				// Look up number in input
				// Get x coordinates from key

				line := input[coord[0]]
				for i := coord[1]; i < len(line); i++ {

					if isDigit(byte(line[i])) {
						numberString += string(line[i])
					} else {
						break
					}
				}
				// Convert number to int
				number, err := strconv.Atoi(numberString)
				if err != nil {
					log.Fatal(err)
				}
				// Add number to numbers
				numbers = append(numbers, number)

				if len(numbers) == 2 {
					total += numbers[0] * numbers[1]
				}
			}
		}
	}

	return total
}

func sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func isOutOfBounds(r, c int, input []string) bool {
	return r < 0 || r >= len(input) || c < 0 || c >= len(input[r])
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
