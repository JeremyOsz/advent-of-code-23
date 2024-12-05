package main

import (
	"fmt"
	io_helpers "jeremyosz/go-advent-2023/2024/helpers/io"
	string_helpers "jeremyosz/go-advent-2023/2024/helpers/strings"
	"regexp"
	"strings"
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
	fmt.Println(`Multiplying`, x, `and`, y)
	if x > 999 || y > 999 {
		return 0
	}
	return x * y
}

func multiplyNumbers(filename string) int {
	// read input as single string
	input := io_helpers.ReadFileString(filename)

	// split input string by do()
	dos := []string{}
	donts := strings.Split(input, "don't()")

	// Extract all do()s from dont blocks
	for i, dont := range donts {
		if i == 0 {
			dos = append(dos, dont)
			continue
		}

		insetDos := strings.Split(dont, "do()")
		dos = append(dos, insetDos[1:]...)
	}

	// Process all do()s
	total := 0
	for _, do := range dos {
		total = total + readMulFuncs(do)
	}

	return total
}

func readMulFuncs(input string) int {
	// read input from filename and return as []string
	re := regexp.MustCompile(`mul\((\d+),\s*(\d+)\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	total := 0
	for _, match := range matches {
		x := string_helpers.ConvertToInt(match[1])
		y := string_helpers.ConvertToInt(match[2])
		total = total + mul(x, y)
	}

	return total
}
