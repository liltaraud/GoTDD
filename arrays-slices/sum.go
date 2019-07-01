package main

// Sum returns the sum of the elements of the array given as a parameter
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll returns a slice containing the sums of all the slices
// given as a parameter
func SumAll(slices ...[]int) (sums []int) {
	for _, slice := range slices {
		sums = append(sums, Sum(slice))
	}
	return
}

// SumAllTails returns a slice containing the sums of all tails
// of the slices given as parameters
func SumAllTails(slices ...[]int) (sums []int) {
	for _, slice := range slices {
		if len(slice) == 0 {
			sums = append(sums, 0)
		} else {
			tail := slice[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}
