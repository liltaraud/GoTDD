package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {

	t.Run("Test with two collections", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSum := func(t *testing.T, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Test with two collections", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{5, 10, 2})
		want := []int{5, 12}

		checkSum(t, got, want)
	})

	t.Run("Test with an empty collection", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{5, 10, 2})
		want := []int{0, 12}

		checkSum(t, got, want)
	})
}
