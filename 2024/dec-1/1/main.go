package main

import (
	"fmt"
	io_helpers "jeremyosz/go-advent-2023/2024/helpers/io"
	string_helpers "jeremyosz/go-advent-2023/2024/helpers/strings"
	"math"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")

	distance := calculateDistance("./input.txt")

	fmt.Println(distance)

}

func readInput(filename string) []string {
	// read input from filename and return as []string
	return io_helpers.ReadFileLines(filename)

}

func calculateDistance(filename string) int {

	distance := 0
	listA := []int{}
	listB := []int{}

	for _, line := range readInput(filename) {
		// iterate over input and split by space
		strings := strings.Split(line, "   ")
		a := string_helpers.ConvertToInt(strings[0])
		b := string_helpers.ConvertToInt(strings[1])

		listA = insertSorted(listA, a)
		listB = insertSorted(listB, b)
	}

	fmt.Println(listA)
	fmt.Println(listB)

	// calculate distance
	for i := 0; i < len(listA); i++ {
		lineDist := math.Abs(float64(listB[i] - listA[i]))
		distance += int(lineDist)
	}

	return distance
}

func insertSorted(list []int, value int) []int {
	index := sort.SearchInts(list, value)
	list = append(list, 0)
	copy(list[index+1:], list[index:])
	list[index] = value
	return list
}
