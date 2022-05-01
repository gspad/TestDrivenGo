package arraysandslices

func Sum(numbers []int) int {
	totalSum := 0
	for _, number := range numbers {
		totalSum += number
	}
	return totalSum
}

func SumAll(numbersToSum ...[]int) (result []int) {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
