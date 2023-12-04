package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const symbol = "*"

func main() {
	input := readInput("./input.txt")
	result := readSchematic(input)
	sum := schematicSum(result)

	// expect sum to be 4361
	fmt.Print(
		"\n\n!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n\n",
		"We have a result.... ", sum,
		"\n\nCheck if it works!",
		"\n\n!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n\n",
	)
}

func readInput(filename string) []string {
	// read input from filename and return as []string
	input := []string{}
	// read file ./input.txt
	// file, err := os.Open("./calibrate.txt")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write each line of ./input.txt to input as a string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

func readSchematic(lines []string) []int {
	// Intialise list of ints to hold part numbers
	parts := []int{}

	// loop through lines
	for lineIndex := range lines {

		// Add parts in line to parts
		parts = append(parts, processLine(lineIndex, lines)...)

		// Print current parts
		fmt.Println("\n\nCurrent parts: ", parts)

	}

	return parts
}

// Take a pointer to a line and a pointer to a list of parts
// For each number in line check if there is an adjacent *
func processLine(index int, lines []string) []int {
	parts := []int{}
	currentLine := lines[index]

	fmt.Print(
		"\n\n===\n\n",
		"Current line \n"+currentLine, "\n\n",
	)

	// Get numbers from current line
	numbers := getNumbersFromLine(currentLine)

	// For each number in line check if there is an adjacent *
	for _, number := range numbers {
		fmt.Print("\n\n====================\n\n")
		fmt.Println("Checking number: ", number)

		num := number[0]
		// numIndex should be Int
		numIndex, _ := strconv.Atoi(number[1])
		start, end := getScanRangeFromIndex(currentLine, num, numIndex)
		// check if there is a symbol on same line adjacent to number
		if checkLine(currentLine, start, end) {
			fmt.Print(
				"\n===== \n\n",
				"Found a symbol on same line adjacent to number: "+num+" in line: "+currentLine+"\n\n",
			)
			parts = append(parts, processNumber(num))
			continue
		}

		// Check the line above if it's not the first line
		if index > 0 {
			previousLine := lines[index-1]
			if checkLine(previousLine, start, end) {
				fmt.Printf(
					"\n===== \n\nFound a symbol on line above adjacent to number: %s\n"+
						"%s\n"+
						"%s\n\n",
					number,
					previousLine,
					currentLine,
				)
				parts = append(parts, processNumber(num))
				continue
			}
		}

		// Check the line below if it's not the last line
		if index < len(lines)-1 {
			nextLine := lines[index+1]
			if checkLine(nextLine, start, end) {
				fmt.Printf(
					"\n===== \n\nFound a symbol on line below adjacent to number: %s\n"+
						"%s\n\n"+
						"%s\n\n",
					number,
					currentLine,
					nextLine,
				)
				parts = append(parts, processNumber(num))
				continue
			}
		}

	}

	// Print lines being added to parts
	fmt.Println("\n%%%%%%%%")
	fmt.Println("Parts being added: ", parts)
	fmt.Println("%%%%%%%%")

	return parts
}

func findAllIndices(s, sub string) []int {
	var indices []int
	for i := 0; i < len(s); i++ {
		if strings.HasPrefix(s[i:], sub) {
			indices = append(indices, i)
		}
	}
	return indices
}

func getScanRangeFromIndex(line string, number string, index int) (int, int) {
	if index == 0 {
		return 0, len(number) + 1
	}
	// If number is at the end of the line range is index - 1 to EOL
	if index+len(number) == len(line) {
		return index - 1, len(line)
	}
	return index - 1, index + len(number) + 1
}

func processNumber(number string) int {
	numberInt, err := strconv.Atoi(number)
	if err != nil {
		log.Fatal(err)
	}
	return numberInt
}

// get each number and their indices in line
func getNumbersFromLine(line string) [][]string {
	// regex to check for numbers
	numberCheck := regexp.MustCompile(`\d+`)

	// Create a slice of all matching numbers and their indices
	// with the structure [[number, index], [number, index]]

	// Find all the numbers in the line
	numbers := numberCheck.FindAllString(line, -1)
	indices := numberCheck.FindAllStringIndex(line, -1)

	fmt.Println("Numbers in line: ", numbers)

	// Create a slice of slices to hold numbers and their indices
	numbersAndIndices := [][]string{}
	for i, number := range numbers {
		numbersAndIndices = append(numbersAndIndices, []string{number, strconv.Itoa(indices[i][0])})
	}

	fmt.Println("Numbers and indices: ", numbersAndIndices)

	return numbersAndIndices
}

func getScanRange(line string, number string) (int, int) {

	var start int
	var end int

	// find start index of number in line
	startIndex := strings.Index(line, number)
	endIndex := startIndex + len(number)

	// match all indices of number in line
	indices := findAllIndices(line, number)

	// if there is more than one index, find the index that matches the number
	if len(indices) > 1 {
		// Print fatal error if there is more than one index
		log.Fatal("More than one index found for number: ", number, " in line: ", line)
	}

	if startIndex > 0 {
		start = startIndex - 1
	} else {
		start = startIndex
	}

	if endIndex < len(line) {
		end = endIndex + 1
	} else {
		end = endIndex
	}

	fmt.Println("Scan Range -- Start index: ", start, "End index: ", end)

	return start, end
}

func checkLine(line string, start int, end int) bool {
	// regex to check for symbols eg. * #  + etc (not ., letters or numbers)
	symbolCheck := regexp.MustCompile(`[^a-zA-Z0-9.]`)

	// check if there is a symbol between start and end
	if symbolCheck.MatchString(line[start:end]) {
		return true
	}
	return false
}

func schematicSum(parts []int) int {
	// add all ints in parts and return sum
	sum := 0
	for _, part := range parts {
		sum += part
	}
	return sum
}
