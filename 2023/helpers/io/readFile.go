package io_helpers

import (
	"bufio"
	"log"
	"os"
)

func ReadFileLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func ReadFileString(filename string) string {
	input, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return string(input)
}
