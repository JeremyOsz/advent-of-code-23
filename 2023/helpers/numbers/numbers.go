package number_helpers

// Check if a number is in a slice of numbers
func SliceContainsInt(numbers []int, number int) bool {
	for _, n := range numbers {
		if n == number {
			return true
		}
	}
	return false
}
