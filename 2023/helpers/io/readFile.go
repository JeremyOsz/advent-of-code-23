package io_helpers

import (
	"log"
	"os"
	"strings"
)

func ReadFileLines(filename string) []string {
	// read input from filename and return as []string
	input, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Split the input into lines
	return strings.Split(string(input), "\n")
}

func ReadFileString(filename string) string {
	input, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return string(input)
}
