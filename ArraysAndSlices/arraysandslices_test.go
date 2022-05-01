package arraysandslices

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	assertEquals := func(expected []int, received []int) {
		if !reflect.DeepEqual(expected, received) {
			t.Errorf("got %d but want %d given", received, expected)
		}
	}

	t.Run("given two slices, when I sum numbers in each slice, I return a slice containing the sums for the slices", func(t *testing.T) {
		numbers1 := []int{7, 8, 9}
		numbers2 := []int{1, 2}

		got := SumAll(numbers1, numbers2)
		want := []int{24, 3}
		assertEquals(got, want)
	})

	t.Run("given three slices, when I sum numbers in each slice, I return a slice containing the sums for the slices", func(t *testing.T) {
		numbers1 := []int{7, 8, 9}
		numbers2 := []int{1, 2}
		numbers3 := []int{1}

		got := SumAll(numbers1, numbers2, numbers3)
		want := []int{24, 3, 1}
		assertEquals(got, want)
	})
}
