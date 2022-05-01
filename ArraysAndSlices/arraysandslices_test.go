package arraysandslices

import "testing"

func TestSum(t *testing.T) {
	assertEquals := func(expected int, received int) {
		if expected != received {
			t.Errorf("got %d but want %d given", received, expected)
		}
	}

	t.Run("given an array of numbers, result should be the sum of the numbers", func(t *testing.T) {
		numbersToSum := []int{1, 2, 3, 4, 5}
		got := Sum(numbersToSum)
		want := 15
		assertEquals(got, want)
	})

	t.Run("given an empty array, result should be 0", func(t *testing.T) {
		var numbers []int
		got := Sum(numbers)
		want := 0
		assertEquals(got, want)
	})
}
