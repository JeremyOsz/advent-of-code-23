package main

import (
	io_helpers "jeremyosz/go-advent-2023/2023/helpers/io"
)

func main() {
	input := readInput("./calibrate.txt")
	// iterate over input and split by space
	// this will make edges
	// Find max and min of each edge
	// place max where it should go in grid
	// find next and next

}

func readInput(filename string) []string {
	// read input from filename and return as []string
	return io_helpers.ReadFileLines(filename)

}
