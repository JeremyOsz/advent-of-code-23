package main

import (
	"fmt"
	"testing"
)

func TestCalibrate(t *testing.T) {
	test1 := round2()

	// expect sum to be 467835
	if test1 != 467835 {
		t.Errorf("Expected 467835, got %d", test1)
	} else {
		fmt.Print(
			"\n\n!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n\n",
			"Expected 467835, got ", test1,
			"\n\nGreat job!",
			"\n\n!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n\n",
		)
	}
}
