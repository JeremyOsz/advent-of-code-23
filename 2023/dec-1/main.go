package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
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

	firstAndLastNumbers, errors := processLines(input)

	// Start a goroutine to receive from the errors channel
	go func() {
		for err := range errors {
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	sum := collectResults(firstAndLastNumbers)

	fmt.Println(sum)
}

func processLines(input []string) (chan int, chan error) {
	firstAndLastNumbers := make(chan int)
	errors := make(chan error)

	var wg sync.WaitGroup

	// Start a goroutine for each line in the input
	for _, line := range input {
		wg.Add(1)
		go func(line string) {
			defer wg.Done()
			number, err := getFirstAndLastNumber(line)
			if err != nil {
				errors <- err
				return
			}
			firstAndLastNumbers <- number
		}(line)
	}

	// Start a goroutine to close the channels after all the other goroutines have finished
	go func() {
		wg.Wait()
		close(firstAndLastNumbers)
		close(errors)
	}()

	return firstAndLastNumbers, errors
}

func collectResults(firstAndLastNumbers chan int) int {
	sum := 0
	for number := range firstAndLastNumbers {
		sum += number
	}
	return sum
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
	n := len(s)
	return strings.Map(func(r rune) rune {
		n--
		return rune(s[n])
	}, s)
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
