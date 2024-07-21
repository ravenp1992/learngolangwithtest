package arrayandslices

// Arrays have a fixed capacity which you define when you declare the variable. We can
// initialize an array in two ways:
// [N]type{value1, value2, ..., valueN} e.g. numbers := [5]int{1, 2, 3, 4, 5}
// [...]type{value1, value2, ..., valueN} e.g. numbers := [...]int{1, 2, 3, 4, 5}

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)

	// for i, numbers := range numbersToSum {
	// 	sums[i] = Sum(numbers)
	// }

	// return sums
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(slicesToSum ...[]int) []int {
	var sums []int
	for _, numbers := range slicesToSum {
		// var tail []int
		if len(numbers) == 0 {
			// tail = append(tail, make([]int, 5)...)
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
			// tail = append(tail, numbers[1:]...)
		}
	}

	return sums
}
