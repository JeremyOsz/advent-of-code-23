package main

import (
	"fmt"
	"testing"
)

func TestCalibrate(t *testing.T) {
	// Test code here

	input, err := readLines("./calibrate.txt")
	if err != nil {
		t.Error("Error reading file: ", err)
	}

	calories := calorieCount(input)

	if calories != 45000 {
		t.Error("Expected 45000, got ", calories)
	}

	fmt.Println("Calories: ", calories)

}

func TestInput(t *testing.T) {
	// Test code here

	input, err := readLines("./input.txt")
	if err != nil {
		t.Error("Error reading file: ", err)
	}

	calories := calorieCount(input)

	fmt.Println("Calories: ", calories)

}
