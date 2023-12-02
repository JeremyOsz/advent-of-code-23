package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Round struct {
	red   int
	green int
	blue  int
}

type Game struct {
	ID     string
	Rounds []Round
}

func main() {
	// read input from ./input.txt - can be any length
	input := ReadInput()

	cubeQty := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	games := readGames(input)

	getPossibleGames(games, cubeQty)

	checkGamePower(games)

}

func checkGamePower(input []Game) {

	totalCubePower := 0

	for _, game := range input {
		fmt.Println("***********************")
		fmt.Println(" ")
		minimumCubes := getMinimumCubes(game)
		fmt.Println("Minimum Cubes for Game", game.ID, ":", minimumCubes)

		// Cube power is the product of the minimum cubes for each color
		cubePower := minimumCubes.red * minimumCubes.green * minimumCubes.blue
		totalCubePower += cubePower
		fmt.Println("Cube Power for Game", game.ID, ":", cubePower)

		fmt.Println(" ")
	}

	fmt.Printf(`
	***********************

	Total Cube Power: %v
	
	***********************`,
		totalCubePower)

}

func getMinimumCubes(game Game) Round {
	// Get the minimum number of cubes for each round
	// and add them together
	minimumCubes := Round{
		red:   0,
		green: 0,
		blue:  0,
	}
	for _, round := range game.Rounds {
		if round.red > minimumCubes.red {
			minimumCubes.red = round.red
		}
		if round.green > minimumCubes.green {
			minimumCubes.green = round.green
		}
		if round.blue > minimumCubes.blue {
			minimumCubes.blue = round.blue
		}
	}
	return minimumCubes
}

func getPossibleGames(input []Game, cubeQty map[string]int) {
	possibleGames := []string{}
	impossibleGames := []string{}

	for _, game := range input {

		// Check if the game is valid for each round
		if checkGame(game, cubeQty) {
			possibleGames = append(possibleGames, game.ID)
		} else {
			impossibleGames = append(impossibleGames, game.ID)
		}
	}
	possibleGamesSum := sumGameIds(possibleGames)
	fmt.Println("***********************")
	fmt.Println(" ")
	fmt.Println("With this cube configuration:")
	fmt.Println("Red:", cubeQty["red"])
	fmt.Println("Green:", cubeQty["green"])
	fmt.Println("Blue:", cubeQty["blue"])
	fmt.Println(" ")
	fmt.Println("There are", len(possibleGames), "possible games and", len(impossibleGames), "impossible games.")
	fmt.Println("Possible Games Sum:", possibleGamesSum)
	fmt.Println(" ")
	fmt.Println("***********************")

}

func sumGameIds(gameIds []string) int {
	gameIdsSum := 0
	for _, gameId := range gameIds {
		gameIdInt, err := strconv.Atoi(gameId)
		if err != nil {
			log.Fatal(err)
		}
		gameIdsSum = gameIdsSum + gameIdInt
	}
	return gameIdsSum
}

func checkGame(game Game, cubeQty map[string]int) bool {
	for _, round := range game.Rounds {
		if !checkRound(round, cubeQty) {
			return false
		}
	}
	return true
}

func ReadInput() []string {
	// read input from ./input.txt - can be any length
	input := []string{}
	// read file ./input.txt
	// file, err := os.Open("./calibrate.txt")
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write each line of ./input.txt to input as a string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

func readGames(input []string) []Game {
	games := []Game{}
	for _, game := range input {
		currentGame := readGame(game)
		prettyPrint(currentGame)
		games = append(games, currentGame)
	}
	return games
}

func readGame(input string) Game {

	// 1. Read game ID - will follow the pattern "Game {ID}:" using regex
	// This will take the number from {ID} between "Game " and ":"
	gameCheck := regexp.MustCompile(`Game (\d+):`)
	gameID := gameCheck.FindStringSubmatch(input)[1]
	// fmt.Println("Game ID", gameID)

	// Remove game ID from input
	input = gameCheck.ReplaceAllString(input, "")

	// 2. Each round will be separated by a ";"
	// 3. The number of cubes will be in the pattern {qty} {colour}, {qty} {colour},
	// eg. "12 red, 13 green, 14 blue"

	// Split input into rounds by ;
	rounds := regexp.MustCompile(`;`).Split(input, -1)

	parsedRounds := []Round{}

	for _, round := range rounds {
		parsedRound := Round{
			red:   0,
			green: 0,
			blue:  0,
		}

		// Find the qty of each color in the round
		parsedRound.red = parseColor("red", round)
		parsedRound.green = parseColor("green", round)
		parsedRound.blue = parseColor("blue", round)

		parsedRounds = append(parsedRounds, parsedRound)
	}

	// Add round and game ID to game struct
	game := Game{
		ID:     gameID,
		Rounds: parsedRounds,
	}

	return game

}

//  parse color and qty from string
//  string will be in the pattern {qty} {colour},
//  takes the color to look for and the string to look in
//  returns the qty of that color
func parseColor(color string, input string) int {
	// Find the color in the string using regex {qty} {colour}
	colorCheck := regexp.MustCompile(`(\d+) ` + color)

	// Find the qty of the color in the string
	colorQty := colorCheck.FindStringSubmatch(input)

	// If the color is not found, return 0
	if len(colorQty) == 0 {
		return 0
	}

	// Convert qty string to int using strconv.Atoi
	colorQtyInt, err := strconv.Atoi(colorQty[1])
	if err != nil {
		log.Fatal(err)
	}

	return colorQtyInt
}

// Pretty print game

func prettyPrint(game Game) {
	fmt.Println("***********************")
	fmt.Println("Game ID:", game.ID)
	for index, round := range game.Rounds {
		fmt.Println("=====================")
		fmt.Println("ROUND:", index+1)
		fmt.Println("Red: ", round.red)
		fmt.Println("Green: ", round.green)
		fmt.Println("Blue: ", round.blue)
		fmt.Println("=====================")
	}
	fmt.Println("***********************")
}

// Checks if the values in the round are valid based on current cubeQty
// Takes a round and a map of cubeQty
// Returns true if the round is valid, false if not
func checkRound(round Round, cubeQty map[string]int) bool {
	if round.red > cubeQty["red"] ||
		round.green > cubeQty["green"] ||
		round.blue > cubeQty["blue"] {
		return false
	}
	return true
}
