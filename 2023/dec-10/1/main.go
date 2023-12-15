package main

import (
	"fmt"
	io_helpers "jeremyosz/go-advent-2023/2023/helpers/io"
)

func main() {
	input := readInput("./input.txt")
	farthest := getFarthest(input)
	fmt.Println(farthest)
}

func readInput(filename string) []string {
	// read input from filename and return as []string
	return io_helpers.ReadFileLines(filename)

}

func getFarthest(input []string) int {
	mapMatrix := make([][]int, len(input))
	for y := 0; y < len(input); y++ {
		mapMatrix[y] = make([]int, len(input[0]))
		for x := 0; x < len(input[0]); x++ {
			if input[y][x] == 'S' {
				mapMatrix[y][x] = 0 // Weigh starting position as 0
			} else if input[y][x] == '.' {
				mapMatrix[y][x] = -2 // Weigh empty space as -2
			} else {
				mapMatrix[y][x] = -1 // Weigh all pipes as -1
			}
		}
	}

	mapMatrix = getMapDistance(input, mapMatrix, 0)

	// Print mapMatrix as a matrix
	for _, row := range mapMatrix {
		for _, value := range row {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}

	// Get max value
	max := 0
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {
			if mapMatrix[y][x] > max {
				max = mapMatrix[y][x]
			}
		}
	}

	return max
}

// Map distances from starting position to m
// This will baheave like a heatmap of distances from the starting position
//
func getMapDistance(input []string, m [][]int, level int) [][]int {
	nextLevel := level + 1 //Sets number to give next node
	found := false

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if m[y][x] == level {
				found = true
				switch input[y][x] {
				case 'S':
					if canMoveLeft(x, y, m) && contains([]byte{'-', 'L', 'F'}, input[y][x-1]) {
						m[y][x-1] = nextLevel
					}
					if canMoveRight(x, y, m) && contains([]byte{'-', '7', 'J'}, input[y][x+1]) {
						m[y][x+1] = nextLevel
					}
					if canMoveUp(x, y, m) && contains([]byte{'|', '7', 'F'}, input[y-1][x]) {
						m[y-1][x] = nextLevel
					}
					if canMoveDown(x, y, m) && contains([]byte{'|', 'J', 'L'}, input[y+1][x]) {
						m[y+1][x] = nextLevel
					}
				case '-':
					if canMoveLeft(x, y, m) && contains([]byte{'-', 'L', 'F'}, input[y][x-1]) {
						m[y][x-1] = nextLevel
					}
					if canMoveRight(x, y, m) && contains([]byte{'-', '7', 'J'}, input[y][x+1]) {
						m[y][x+1] = nextLevel
					}
				case '|':
					if canMoveUp(x, y, m) && contains([]byte{'|', '7', 'F'}, input[y-1][x]) {
						m[y-1][x] = nextLevel
					}
					if canMoveDown(x, y, m) && contains([]byte{'|', 'J', 'L'}, input[y+1][x]) {
						m[y+1][x] = nextLevel
					}
				case '7':
					if canMoveLeft(x, y, m) && contains([]byte{'-', 'L', 'F'}, input[y][x-1]) {
						m[y][x-1] = nextLevel
					}
					if canMoveDown(x, y, m) && contains([]byte{'|', 'J', 'L'}, input[y+1][x]) {
						m[y+1][x] = nextLevel
					}
				case 'F':
					if canMoveRight(x, y, m) && contains([]byte{'-', '7', 'J'}, input[y][x+1]) {
						m[y][x+1] = nextLevel
					}
					if canMoveDown(x, y, m) && contains([]byte{'|', 'J', 'L'}, input[y+1][x]) {
						m[y+1][x] = nextLevel
					}
				case 'J':
					if canMoveLeft(x, y, m) && contains([]byte{'-', 'L', 'F'}, input[y][x-1]) {
						m[y][x-1] = nextLevel
					}
					if canMoveUp(x, y, m) && contains([]byte{'|', '7', 'F'}, input[y-1][x]) {
						m[y-1][x] = nextLevel
					}
				case 'L':
					if canMoveRight(x, y, m) && contains([]byte{'-', '7', 'J'}, input[y][x+1]) {
						m[y][x+1] = nextLevel
					}
					if canMoveUp(x, y, m) && contains([]byte{'|', '7', 'F'}, input[y-1][x]) {
						m[y-1][x] = nextLevel
					}
				}
			}
		}
	}

	if found { // If we found a node at this level, call again with next level
		return getMapDistance(input, m, nextLevel)
	}
	return m
}

func isValidPosition(x, y, maxX, maxY int) bool {
	return x >= 0 && x < maxX && y >= 0 && y < maxY
}

// Takes a list of chars and returns a map of those chars with value true (for easy lookup)
func makeMap(chars ...byte) map[byte]bool {
	m := make(map[byte]bool)
	for _, c := range chars {
		m[c] = true
	}
	return m
}

// Array contains byte
func contains(arr []byte, b byte) bool {
	for _, a := range arr {
		if a == b {
			return true
		}
	}
	return false
}

// Check if the cell to the left is within the grid and unvisited
func canMoveLeft(x, y int, m [][]int) bool {
	return x > 0 && m[y][x-1] == -1
}

// Check if the cell to the right is within the grid and unvisited
func canMoveRight(x, y int, m [][]int) bool {
	return x < len(m[y])-1 && m[y][x+1] == -1
}

// Check if the cell above is within the grid and unvisited
func canMoveUp(x, y int, m [][]int) bool {
	return y > 0 && m[y-1][x] == -1
}

// Check if the cell below is within the grid and unvisited
func canMoveDown(x, y int, m [][]int) bool {
	return y < len(m)-1 && m[y+1][x] == -1
}
