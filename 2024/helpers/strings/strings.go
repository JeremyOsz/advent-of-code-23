package string_helpers

import (
	"fmt"
	"log"
	"strconv"
)

func ConvertToInt(number string) int {
	numberInt, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("Error converting string to int")
		log.Fatal(err)
	}
	return numberInt
}

func ConvertToUInt(s string) uint {
	i, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		panic(err)
	}
	return uint(i)
}

func ConvertSliceToInts(slice []string) []int {
	intSlice := make([]int, 0, len(slice))
	for _, number := range slice {
		if number == "" {
			continue
		}
		numberInt, err := strconv.Atoi(number)
		if err != nil {
			log.Fatal(err)
		}
		intSlice = append(intSlice, numberInt)
	}
	return intSlice
}
