package iteration

func Repeat(character string, iterator int) string {
	var repeated string

	for i := 0; i < iterator; i++ {
		repeated += character
	}

	return repeated
}