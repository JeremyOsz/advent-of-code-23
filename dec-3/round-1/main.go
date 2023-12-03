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
	// read input from ./input.txt - can be any length
	input := readInput("calibrate.txt")

	// read schematic from input
	readSchematic(input)

	// // calculate power of schematic
	// power := calculatePower(schematic)

	// // print power
	// log.Println(power)
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
	for lineIndex, line := range lines {

		// Get numbers from each line
		numbers := getNumbersFromLine(line)

		// Print line and numbers
		fmt.Print(
			"\n\n*******************************************\n\n",
			"Processing line: \n ", line, "\n",
			"Numbers in line: ", strings.Join(numbers, ", "),
		)

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

		start, end := getScanRange(currentLine, number)

		// check if there is a symbol on same line adjacent to number
		if checkLine(currentLine, start, end) {
			fmt.Print(
				"\n===== \n\n",
				"Found a symbol on same line adjacent to number: "+number+" in line: "+currentLine+"\n\n",
			)
			parts = append(parts, processNumber(number))
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
				parts = append(parts, processNumber(number))
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
				parts = append(parts, processNumber(number))
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

func processNumber(number string) int {
	numberInt, err := strconv.Atoi(number)
	if err != nil {
		log.Fatal(err)
	}
	return numberInt
}

func getNumbersFromLine(line string) []string {
	// regex to check for numbers
	numberCheck := regexp.MustCompile(`\d+`)
	// find all numbers in line
	numbers := numberCheck.FindAllString(line, -1)

	fmt.Println("Numbers in line: ", numbers)

	return numbers

}

func getScanRange(line string, number string) (int, int) {

	var start int
	var end int

	// find start index of number in line
	startIndex := strings.Index(line, number)
	endIndex := startIndex + len(number)

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
