package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	getSteps("./input.txt")
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

func getSteps(filename string) int {
	// getSteps returns the number of steps to reach the exit
	// from the input file
	input := readInput(filename)

	instructions := input[0]

	// Create a map like this:
	// [AAA: [BBB,  CCC]]
	// [BBB: [DDD,  EEE]]
	nodes := make(map[string][]string)
	for _, line := range input[2:] { // skip the first 2 lines
		// fmt.Println(line)
		// Parse AAA = (BBB, CCC) as AAA: [BBB, CCC]
		lineSplit := strings.Split(line, " = ")
		key := lineSplit[0]

		// Parse (BBB, CCC) as [BBB, CCC]
		value := strings.Split(lineSplit[1], ", ")
		L := value[0][1:]
		R := value[1][:len(value[1])-1]

		// Add to map
		nodes[key] = []string{L, R}
	}

	fmt.Println(nodes)

	// Follow the instructions
	currentNode := nodes["AAA"]
	steps := 1
	return stepThroughNodes(currentNode, instructions, nodes, steps)

}

func stepThroughNodes(
	currentNode []string,
	instructions string,
	nodes map[string][]string,
	steps int) int {
	for _, instruction := range instructions {
		nextNode := ""
		if instruction == 'L' {
			nextNode = currentNode[0]
		} else {
			nextNode = currentNode[1]
		}

		// If key of currentNode is ZZZ, we're done
		fmt.Println("Current Node: ", currentNode)
		fmt.Println("Next Node: ", nextNode)

		if nextNode == "ZZZ" {
			fmt.Println("Reached ZZZ in ", steps, " steps")
			return steps
		} else {
			currentNode = nodes[nextNode]
			steps++
		}
	}
	return stepThroughNodes(currentNode, instructions, nodes, steps)
}
