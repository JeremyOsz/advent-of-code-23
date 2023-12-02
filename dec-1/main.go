package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// hello world in go

func main() {

	// read calibrationValues from ./input.txt - can be any length
	calibrationValues := []string{}
	// read file ./input.txt
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write each line of ./input.txt to calibrationValues as a string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		calibrationValues = append(calibrationValues, scanner.Text())
	}

	// calibrationNumbers := [4]int{}
	// apply getFirstAndLastNumber to each string in calibrationValues and store in calibrationNumbers
	for i, calibrationValue := range calibrationValues {
		calibrationValues[i] = strconv.Itoa(getFirstAndLastNumber(calibrationValue))
	}

	// add calibrationNumbers together
	calibrationSum := 0
	for _, calibrationValue := range calibrationValues {
		calibrationValueInt, err := strconv.Atoi(calibrationValue)
		if err != nil {
			fmt.Println(err)
		}
		calibrationSum += calibrationValueInt
	}

	fmt.Println(calibrationSum)
}

func parseNumber(number string) string {
	// dictionary mapping of number words to numbers
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

	// reverseNumbersToString := map[string]string{
	// 	"orez":  "0",
	// 	"eno":   "1",
	// 	"owt":   "2",
	// 	"eerht": "3",
	// 	"ruof":  "4",
	// 	"evif":  "5",
	// 	"xis":   "6",
	// 	"neves": "7",
	// 	"thgie": "8",
	// 	"enin":  "9",
	// }

	if _, err := strconv.Atoi(number); err != nil {
		number = numbersToString[number]
	}
	return number
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		// swap the letters of the string
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns) // return the reversed string
}

// Function to search a string for a number charecter backwards from the end of the string
func getFirstAndLastNumber(s string) int {
	re := regexp.MustCompile(`[0-9]|one|two|three|four|five|six|seven|eight|nine`)
	// Find first number in string
	fmt.Println(s)
	firstNumber := re.FindString(s)

	// reverse s
	reverseS := reverse(s)

	// create reverse regex
	reverseRe := regexp.MustCompile(`[0-9]|orez|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin`)
	// Find last number in string using a lookbehind regex
	lastNumber := reverseRe.FindString(reverseS)

	// Make a slice of the last number

	//  Parse firstNumber into a digit
	fistDigit := parseNumber(firstNumber)

	//  restore lastNumber to original order
	lastDigitRestored := reverse(lastNumber)

	fmt.Println(firstNumber, lastDigitRestored)
	lastDigit := parseNumber(lastDigitRestored)

	// Make a string of the first and last number
	// check if numbers is length one if so then firstAndLastNumber is the only number
	// else firstAndLastNumber is the first and last number
	firstAndLastDigit := fistDigit + lastDigit
	fmt.Println(firstAndLastDigit)

	// convert firstAndLastNumber to int
	firstAndLastDigitInt, err := strconv.Atoi(firstAndLastDigit)

	if err != nil {
		fmt.Println(err)

	}
	return firstAndLastDigitInt
}
