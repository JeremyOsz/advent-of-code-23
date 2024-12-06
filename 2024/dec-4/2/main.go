package main

import (
	"fmt"
	io_helpers "jeremyosz/go-advent-2023/2023/helpers/io"
)

func main() {
	fmt.Println("Hello, World!")

	result := countXmas("./input.txt")
	fmt.Printf(
		`
	------------------------------

	Result: %d

	------------------------------
	`, result)
}

func readInput(filename string) []string {
	// read input from filename and return as []string
	return io_helpers.ReadFileLines(filename)

}

func countXmas(filename string) int {
	// read input from filename
	input := readInput(filename)
	xmas := 0

	for i, line := range input {
		for j, char := range line {
			if char == 'A' {
				// check diagonals
				if i-1 >= 0 && j-1 >= 0 && i+1 < len(input) && j+1 < len(line) {
					// M . S
					// . A .
					// M . S
					if input[i-1][j-1] == 'M' && input[i-1][j+1] == 'S' &&
						input[i+1][j-1] == 'M' && input[i+1][j+1] == 'S' {
						xmas++
					}

					// S . M
					// . A .
					// S . M
					if input[i-1][j-1] == 'S' && input[i-1][j+1] == 'M' &&
						input[i+1][j-1] == 'S' && input[i+1][j+1] == 'M' {
						xmas++
					}

					// M . M
					// . A .
					// S . S
					if input[i-1][j-1] == 'M' && input[i-1][j+1] == 'M' &&
						input[i+1][j-1] == 'S' && input[i+1][j+1] == 'S' {
						xmas++
					}

					// S . S
					// . A .
					// M . M
					if input[i-1][j-1] == 'S' && input[i-1][j+1] == 'S' &&
						input[i+1][j-1] == 'M' && input[i+1][j+1] == 'M' {
						xmas++
					}

				}
			}
		}
	}

	return xmas
}
