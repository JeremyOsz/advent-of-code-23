package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello, World!")
}

func calorieCount(input []string) int {
	// Break input into chunks - using whitespace as delimiter

	elf := []int{}
	elves := [][]int{}

	for _, line := range input {
		fmt.Println(line)
		if line != "" {
			calories, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Add calories to elf
			elf = append(elf, calories)
		} else {
			// Add elf to elves
			elves = append(elves, elf)

			// Reset elf
			elf = []int{}
		}

	}

	// Sum the calories for each elf
	maxCalories := []int{}
	for _, elf := range elves {
		sum := 0
		for _, calorie := range elf {
			sum += calorie
		}
		maxCalories = append(maxCalories, sum)
	}

	fmt.Println(maxCalories)

	// Find the max
	max := 0
	for _, calorie := range maxCalories {
		if calorie > max {
			max = calorie
		}
	}

	return max
}

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
