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

// Function that sums the 'tails' of arrays
// (items after the first element) in the collection
func SumAllTails(numbersToSum ...[]int) []int{
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}