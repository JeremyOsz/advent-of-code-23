package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	minSteps := getSteps("./input.txt")
	fmt.Println("Min steps: ", minSteps)
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

	//  handle starting nodes concurrently
	pathSteps := make(chan int, len(startingNodes))
	var wg sync.WaitGroup
	for _, node := range startingNodes {
		wg.Add(1)
		go func(node string) {
			defer wg.Done()
			steps := 1
			minSteps := stepThroughNodes(nodes[node], instructions, nodes, steps)
			fmt.Println("Min steps for ", node, " is ", minSteps)
			pathSteps <- minSteps
		}(node)
	}
	wg.Wait()
	close(pathSteps)

	// Convert pathSteps to []int
	pathStepsArr := make([]int, len(startingNodes))
	i := 0
	for step := range pathSteps {
		pathStepsArr[i] = step
		i++
	}

	fmt.Println("Path steps: ", pathSteps)

	return LCM(pathStepsArr[0], pathStepsArr[1], pathStepsArr...)
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

		if nextNode[2] == 'Z' {
			fmt.Println("Reached Z in ", steps, " steps")
			return steps
		} else {
			currentNode = nodes[nextNode]
			steps++
		}
	}
	return stepThroughNodes(currentNode, instructions, nodes, steps)
}

func allTrue(arr []bool) bool {
	for _, v := range arr {
		if !v {
			return false
		}
	}
	return true
}

// Function to calculate GCD (Greatest Common Divisor)
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// Function to calculate LCM (Least Common Multiple)
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
