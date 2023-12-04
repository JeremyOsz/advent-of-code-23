package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// calorieCount calculates the sum of the top 3 calories for each elf
func calorieCount(input []string) int {
	// Break input into chunks - using whitespace as delimiter
	elves := chunkInput(input)

	// Sum the calories for each elf
	maxCalories := sumCalories(elves)

	// Sort the calories in descending order
	sort.Slice(maxCalories, func(i, j int) bool {
		return maxCalories[i] > maxCalories[j]
	})

	// Sum the top 3 calories
	sum := 0
	for i := 0; i < 3 && i < len(maxCalories); i++ {
		sum += maxCalories[i]
	}

	return sum
}

// chunkInput breaks input into chunks using whitespace as delimiter
func chunkInput(input []string) [][]int {
	elf := []int{}
	elves := [][]int{}

	for i, line := range input {

		// If line is not empty, add calorie to elf
		if line != "" {
			calories, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Add calories to elf
			elf = append(elf, calories)

			// If line is the last line, add elf to elves
			if i == len(input)-1 {
				elves = append(elves, elf)
			}
		} else {
			// Add elf to elves
			elves = append(elves, elf)

			// Reset elf
			elf = []int{}
		}

	}
	return elves
}

// sumCalories sums the calories for each elf
func sumCalories(elves [][]int) []int {
	maxCalories := []int{}
	for _, elf := range elves {
		sum := 0
		for _, calorie := range elf {
			sum += calorie
		}
		maxCalories = append(maxCalories, sum)
	}
	return maxCalories
}

// readLines reads a file into an array of strings
func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
