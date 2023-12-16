package main

import (
	"fmt"
	io_helpers "jeremyosz/go-advent-2023/2023/helpers/io"
	"regexp"
	"strings"
)

func main() {
	input := readInput("./input.txt")
	farthest := getTilesEnclosed(input)
	fmt.Println(farthest)
}

func readInput(filename string) []string {
	// read input from filename and return as []string
	return io_helpers.ReadFileLines(filename)

}

func getTilesEnclosed(s []string) int {
	var sy, sx int
	m := make([][]int, len(s))

	// Make map matrix of pipes, empty space, and starting position
	for y := 0; y < len(s); y++ {
		m[y] = make([]int, len(s[y]))
		for x := 0; x < len(s[y]); x++ {
			if s[y][x] == 'S' {
				m[y][x] = 0
				sy = y
				sx = x
			} else if s[y][x] == '.' {
				m[y][x] = -2
			} else {
				m[y][x] = -1
			}
		}
	}

	// Get map distances
	m = getMapDistance(s, m, 0)

	s[sy] = strings.Replace(s[sy], "S", string(getStartSymbol(s, sy, sx)), -1) // Replace S with correct symbol

	// Make all non visited cells empty
	for y := 0; y < len(s); y++ {
		newS := make([]byte, len(s[y]))
		for x := 0; x < len(s[y]); x++ {
			if m[y][x] == -1 {
				newS[x] = '.'
			} else {
				newS[x] = s[y][x]
			}
		}
		s[y] = string(newS)
	}

	// Check if any cells are enclosed
	noWall := regexp.MustCompile(`F-*7|L-*J`)
	wall := regexp.MustCompile(`F-*J|L-*7`)

	// for each row, replace all enclosed cells with empty space
	// And all walls with pipes
	for i, ss := range s {
		s1 := noWall.ReplaceAllString(ss, " ")
		s2 := wall.ReplaceAllString(s1, "|")
		s[i] = s2
	}

	for _, row := range s {
		for _, value := range row {
			fmt.Printf("%s ", string(value))
		}
		fmt.Println()
	}

	// Sum all enclosed cells
	// We work this out by seeing if number of walls are odd
	// If odd, there is one more ahead - so we are enclosed
	var (
		parity int
		count  int = 0
	)
	for y := 0; y < len(s); y++ {
		parity = 0
		for x := 0; x < len(s[y]); x++ {
			if s[y][x] == '|' {
				parity++
			}
			if s[y][x] == '.' && parity%2 == 1 {
				count++
			}
		}
	}
	return count
}

// Takes coordinates of S and returns the symbol that should be there
func getStartSymbol(s []string, sy, sx int) uint8 {

	// Initialise as empty
	directions := map[string]uint8{
		"up":    '.',
		"down":  '.',
		"left":  '.',
		"right": '.',
	}

	// Check if we can go in each direction (will be . if out of bounds)
	if sy > 0 {
		directions["up"] = s[sy-1][sx]
	}
	if sy < len(s)-1 {
		directions["down"] = s[sy+1][sx]
	}
	if sx > 0 {
		directions["left"] = s[sy][sx-1]
	}
	if sx < len(s[0])-1 {
		directions["right"] = s[sy][sx+1]
	}

	// Compare directions to possilbe symbols
	canGo := map[string]bool{
		"up":    directions["up"] == '7' || directions["up"] == 'F' || directions["up"] == '|',
		"down":  directions["down"] == 'J' || directions["down"] == 'L' || directions["down"] == '|',
		"left":  directions["left"] == 'F' || directions["left"] == 'L' || directions["left"] == '-',
		"right": directions["right"] == 'J' || directions["right"] == '7' || directions["right"] == '-',
	}

	// Determine symbol by 2 possible directions
	switch {
	case canGo["up"] && canGo["down"]:
		return '|'
	case canGo["left"] && canGo["right"]:
		return '-'
	case canGo["left"] && canGo["down"]:
		return '7'
	case canGo["right"] && canGo["down"]:
		return 'F'
	case canGo["right"] && canGo["up"]:
		return 'L'
	case canGo["left"] && canGo["up"]:
		return 'J'
	default:
		return 'S'
	}
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
