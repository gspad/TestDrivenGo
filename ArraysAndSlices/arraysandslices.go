package arraysandslices

func Sum(numbers []int) int {
	totalSum := 0
	for _, number := range numbers {
		totalSum += number
	}
	return totalSum
}
