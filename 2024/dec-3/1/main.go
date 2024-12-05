package main

import (
	"fmt"
	io_helpers "jeremyosz/go-advent-2023/2024/helpers/io"
	string_helpers "jeremyosz/go-advent-2023/2024/helpers/strings"
	"regexp"
)

func main() {
	fmt.Println("Hello, World!")

	result := multiplyNumbers("./input.txt")

	fmt.Printf(`
	------------------------------

	Result: %d

	------------------------------
	`, result)
}

func mul(x, y int) int {
	if x > 999 || y > 999 {
		return 0
	}
	return x * y
}

func multiplyNumbers(filename string) int {
	// read input as single string
	input := io_helpers.ReadFileString(filename)

	// Use REGEX to extract subsets of input WHERE it looks like mul(1, 2), mul(123, 243), etc.
	re := regexp.MustCompile(`mul\((\d+),\s*(\d+)\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	total := 0
	for _, match := range matches {
		x := string_helpers.ConvertToInt(match[1])
		y := string_helpers.ConvertToInt(match[2])
		total = total + mul(x, y)
	}
	// fmt.Println(input)
	fmt.Println(total)
	return total
}
