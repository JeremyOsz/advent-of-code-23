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
	fmt.Println(instructions)

	// Create a map like this:
	// [AAA: [BBB,  CCC]]
	// [BBB: [DDD,  EEE]]
	nodes := make(map[string][]string)
	startingNodes := []string{}
	for _, line := range input[2:] { // skip the first 2 lines
		// fmt.Println(line)
		// Parse AAA = (BBB, CCC) as AAA: [BBB, CCC]
		lineSplit := strings.Split(line, " = ")
		key := lineSplit[0]

		// Parse (BBB, CCC) as [BBB, CCC]
		value := strings.Split(lineSplit[1], ", ")
		L := value[0][1:]
		R := value[1][:len(value[1])-1]

		// Add node to map
		nodes[key] = []string{L, R}

		// If node is starting node, add key to startingNodes
		// get last digit of key
		if key[2] == 'A' {
			startingNodes = append(startingNodes, key)
		}
	}

	// fmt.Println("nodes", nodes)
	// fmt.Println("startingNodes: ", startingNodes)

	steps := 1

	return stepThroughNodes(startingNodes, instructions, nodes, steps)

}

func stepThroughNodes(
	currentNodes []string,
	instructions string,
	nodes map[string][]string,
	steps int) int {
	fmt.Println("currentNodes: ", currentNodes)
	for _, instruction := range instructions {

		nextNodes := currentNodes
		atZ := []bool{}
		for i := range currentNodes {
			if instruction == 'L' {
				nextNodes[i] = nodes[currentNodes[i]][0]
			} else {
				nextNodes[i] = nodes[currentNodes[i]][1]
			}

			if nextNodes[i][2] == 'Z' {
				atZ = append(atZ, true)
			} else {
				atZ = append(atZ, false)
			}
		}

		// if all in atZ are true, return steps
		if allTrue(atZ) {
			fmt.Println("Reached Z in ", steps, " steps")
			return steps
		} else {
			currentNodes = nextNodes
			steps++
		}
	}
	return stepThroughNodes(currentNodes, instructions, nodes, steps)
	// return steps
}

func allTrue(arr []bool) bool {
	for _, v := range arr {
		if !v {
			return false
		}
	}
	return true
}
