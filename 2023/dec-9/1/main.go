package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
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

func analyseInput(filename string) ([]int, int) {
	// analyseInput returns the values and sum of the input
	input := readInput(filename)
	nextValues := []int{}
	sum := 0

	for _, line := range input {

	}

	return nextValues, sum
}
