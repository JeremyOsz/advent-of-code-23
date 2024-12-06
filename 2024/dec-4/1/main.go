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
	// delta := []int{-3, 3, 3, -3}

	for i, line := range input {
		for j, char := range line {
			if char == 'X' {
				// check if the next 3 characters are M, A, S
				if j+3 < len(line) && line[j+1] == 'M' && line[j+2] == 'A' && line[j+3] == 'S' {
					xmas++
				}

				// check if the previous 3 characters are S, A, M
				if j-3 >= 0 && line[j-1] == 'M' && line[j-2] == 'A' && line[j-3] == 'S' {
					xmas++
				}

				// check diagonals
				if i+3 < len(input) && j+3 < len(line) {
					if input[i+1][j+1] == 'M' && input[i+2][j+2] == 'A' && input[i+3][j+3] == 'S' {
						xmas++
					}
				}

				if i+3 < len(input) && j-3 >= 0 {
					if input[i+1][j-1] == 'M' && input[i+2][j-2] == 'A' && input[i+3][j-3] == 'S' {
						xmas++
					}
				}

				if i-3 >= 0 && j-3 >= 0 {
					if input[i-1][j-1] == 'M' && input[i-2][j-2] == 'A' && input[i-3][j-3] == 'S' {
						xmas++
					}
				}

				if i-3 >= 0 && j+3 < len(line) {
					if input[i-1][j+1] == 'M' && input[i-2][j+2] == 'A' && input[i-3][j+3] == 'S' {
						xmas++
					}
				}

				// check verticals
				if i+3 < len(input) {
					if input[i+1][j] == 'M' && input[i+2][j] == 'A' && input[i+3][j] == 'S' {
						xmas++
					}
				}

				if i-3 >= 0 {
					if input[i-1][j] == 'M' && input[i-2][j] == 'A' && input[i-3][j] == 'S' {
						xmas++
					}
				}
			}
		}
	}

	return xmas
}
