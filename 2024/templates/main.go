package main

import (
	"fmt"
	io_helpers "jeremyosz/go-advent-2023/2023/helpers/io"
)

func main() {
	fmt.Println("Hello, World!")
}

func readInput(filename string) []string {
	// read input from filename and return as []string
	return io_helpers.ReadFileLines(filename)

}
