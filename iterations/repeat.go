package iterations

// Repeat the character 5 times.
func Repeat(character string, repeat int) string {
	var repeated string
	// repeated := ""

	// classic for loop
	// for i := 0; i < 5; i++ {
	// 	repeated += character
	// }

	// i := 1
	// for i <= 5 {
	// 	repeated += character
	// 	i++
	// }

	for range repeat {
		repeated += character
	}

	return repeated
}
