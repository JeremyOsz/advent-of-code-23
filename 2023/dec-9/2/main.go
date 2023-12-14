package main

import (
	"fmt"
	string_helpers "jeremyosz/go-advent-2023/2023/helpers/strings"
	"log"
	"os"
	"strings"
)

type Pyramid [][]int

func main() {
	nextValues, sum := analyseInput("./input.txt")

	fmt.Printf(`
	=====================================

	nextValues: %v
	sum: %d

	=====================================
	`, nextValues, sum)
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
		// fmt.Println(line)
		values := string_helpers.ConvertSliceToInts(strings.Split(line, " "))
		pyramid := buildPyramid(Pyramid{values})
		nextValue := getNextValue(pyramid)
		// fmt.Println(pyramid)
		// fmt.Println(nextValue)
		nextValues = append(nextValues, nextValue)
		sum += nextValue
	}

	return nextValues, sum
}

func buildPyramid(pyramid Pyramid) Pyramid {
	lastLine := pyramid[len(pyramid)-1]
	nextLine := []int{}
	zeroes := 0
	for i := 0; i < len(lastLine)-1; i++ {
		difference := lastLine[i+1] - lastLine[i]
		nextLine = append(nextLine, difference)
		if difference == 0 {
			zeroes++
		}
	}

	pyramid = append(pyramid, nextLine)

	if zeroes == len(lastLine)-1 {
		return pyramid
	}

	return buildPyramid(pyramid)
}

func getNextValue(pyramid Pyramid) int {
	// fmt.Println("=====================================")
	// fmt.Println("Getting next value for pyramid: ", pyramid)
	sum := 0
	sums := []int{}
	// reverse loop through pyramid
	for i := len(pyramid) - 1; i > 0; i-- {
		x := pyramid[i][len(pyramid[i])-1]
		y := pyramid[i-1][len(pyramid[i-1])-1]
		sum = x + y
		pyramid[i-1] = append(pyramid[i-1], sum)
		sums = append(sums, sum)
	}

	return sum
}
