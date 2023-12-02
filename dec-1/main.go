package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// ReadInput reads calibrationValues from ./input.txt
func ReadInput() ([]string, error) {
	calibrationValues := []string{}
	file, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		calibrationValues = append(calibrationValues, scanner.Text())
	}
	return calibrationValues, scanner.Err()
}

func main() {
	input, err := ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	firstAndLastNumbers := []int{}
	for _, line := range input {
		number, err := getFirstAndLastNumber(line)
		if err != nil {
			log.Fatal(err)
		}
		firstAndLastNumbers = append(firstAndLastNumbers, number)
	}

	sum := 0
	for _, number := range firstAndLastNumbers {
		sum += number
	}

	fmt.Println(sum)
}

// parseNumber maps number words to numbers
func parseNumber(number string) string {
	numbersToString := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	if _, err := strconv.Atoi(number); err != nil {
		number = numbersToString[number]
	}
	return number
}

// reverse reverses a string
func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

// getFirstAndLastNumber searches a string for a number character from the start and end of the string
func getFirstAndLastNumber(s string) (int, error) {
	re := regexp.MustCompile(`[0-9]|one|two|three|four|five|six|seven|eight|nine`)
	firstNumber := re.FindString(s)

	reverseS := reverse(s)
	reverseRe := regexp.MustCompile(`[0-9]|orez|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin`)
	lastNumber := reverseRe.FindString(reverseS)

	firstDigit := parseNumber(firstNumber)
	lastDigitRestored := reverse(lastNumber)
	lastDigit := parseNumber(lastDigitRestored)

	firstAndLastDigit := firstDigit + lastDigit
	return strconv.Atoi(firstAndLastDigit)
}
