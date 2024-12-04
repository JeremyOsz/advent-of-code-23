package main

import (
	"fmt"
	io_helpers "jeremyosz/go-advent-2023/2024/helpers/io"
	string_helpers "jeremyosz/go-advent-2023/2024/helpers/strings"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func readInput(filename string) []string {
	// read input from filename and return as []string
	return io_helpers.ReadFileLines(filename)

}

func calculateSimilarity(filename string) int {
	listA := []int{}
	listB := []int{}

	for _, line := range readInput(filename) {
		// iterate over input and split by space
		strings := strings.Split(line, "   ")
		listA = append(listA, string_helpers.ConvertToInt(strings[0]))
		listB = append(listB, string_helpers.ConvertToInt(strings[1]))
	}

	// sort listA and listB
	// sort.Ints(listA)
	// sort.Ints(listB)

	fmt.Println(listA)
	fmt.Println(listB)

	// calculate similarity

	similarity := 0
	for i := 0; i < len(listA); i++ {
		lineSimilarity := 0
		for j := 0; j < len(listB); j++ {
			if listA[i] == listB[j] {
				lineSimilarity++
			}
		}

		similarity = similarity + listA[i]*lineSimilarity

		fmt.Println(similarity)
	}

	return similarity
}
