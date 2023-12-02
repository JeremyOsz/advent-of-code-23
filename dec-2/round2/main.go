package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"sync"
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

type CubePower struct {
	GameID       string
	MinimumCubes Round
	CubePower    int
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

func calculateGamePower(input []Game) (totalCubePower int, powers []CubePower) {
	powerChan := make(chan CubePower, len(input))
	var wg sync.WaitGroup

	for _, game := range input {
		wg.Add(1)
		go func(game Game) {
			defer wg.Done()
			minimumCubes := getMinimumCubes(game)
			cubePower := minimumCubes.red * minimumCubes.green * minimumCubes.blue
			powerChan <- CubePower{game.ID, minimumCubes, cubePower}
		}(game)
	}

	go func() {
		wg.Wait()
		close(powerChan)
	}()

	for power := range powerChan {
		totalCubePower += power.CubePower
		powers = append(powers, power)
	}

	return totalCubePower, powers
}

func printGamePower(totalCubePower int, powers []CubePower) {
	for _, power := range powers {
		fmt.Println("***********************")
		fmt.Println(" ")
		fmt.Println("Minimum Cubes for Game", power.GameID, ":", power.MinimumCubes)
		fmt.Println("Cube Power for Game", power.GameID, ":", power.CubePower)
		fmt.Println(" ")
	}
	fmt.Printf(`
	***********************

	Total Cube Power: %v
	
	***********************`,
		totalCubePower)
}

func checkGamePower(input []Game) {
	totalCubePower, powers := calculateGamePower(input)
	printGamePower(totalCubePower, powers)
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

func checkRound(round Round, cubeQty map[string]int) bool {
	return round.red <= cubeQty["red"] && round.green <= cubeQty["green"] && round.blue <= cubeQty["blue"]
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

func parseGameID(input string) string {
	gameCheck := regexp.MustCompile(`Game (\d+):`)
	return gameCheck.FindStringSubmatch(input)[1]
}

func removeGameID(input string) string {
	gameCheck := regexp.MustCompile(`Game (\d+):`)
	return gameCheck.ReplaceAllString(input, "")
}

func splitIntoRounds(input string) []string {
	return regexp.MustCompile(`;`).Split(input, -1)
}

func parseRound(round string) Round {
	return Round{
		red:   parseColor("red", round),
		green: parseColor("green", round),
		blue:  parseColor("blue", round),
	}
}

func parseRounds(rounds []string) []Round {
	parsedRounds := make([]Round, len(rounds))
	for i, round := range rounds {
		parsedRounds[i] = parseRound(round)
	}
	return parsedRounds
}

func readGame(input string) Game {
	gameID := parseGameID(input)
	input = removeGameID(input)
	rounds := splitIntoRounds(input)
	parsedRounds := parseRounds(rounds)

	return Game{
		ID:     gameID,
		Rounds: parsedRounds,
	}
}

//  parse color and qty from string
//  string will be in the pattern {qty} {colour},
//  takes the color to look for and the string to look in
//  returns the qty of that color
func parseColor(color, input string) int {
	// Use a regular expression to find the quantity of the color in the string.
	colorCheck := regexp.MustCompile(`(\d+)\s+` + regexp.QuoteMeta(color))

	matches := colorCheck.FindStringSubmatch(input)
	if matches == nil {
		// The color was not found in the string.
		return 0
	}

	// Convert the quantity string to an integer.
	qty, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0
	}

	return qty
}

// PrettyPrint displays game details in a formatted manner.
func prettyPrint(game Game) {
	fmt.Printf("***********************\nGame ID: %s\n", game.ID)
	for index, round := range game.Rounds {
		fmt.Printf("=====================\nROUND: %d\nRed: %d\nGreen: %d\nBlue: %d\n=====================\n",
			index+1, round.red, round.green, round.blue)
	}
	fmt.Println("***********************")
}
