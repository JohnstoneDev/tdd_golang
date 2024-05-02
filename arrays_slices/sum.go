package arraysslices

// Function that sums app the items in an integer array
// and returns the value
func Sum(numbers []int) int {
	sum := 0

	for _, num := range numbers {
		sum += num
	}

	return sum
}