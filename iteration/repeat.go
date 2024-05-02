package iteration

// Function that repeats a character iterator number of times
func Repeat(character string, iterator int) string {
	var repeated string

	for i := 0; i < iterator; i++ {
		repeated += character
	}

	return repeated
}