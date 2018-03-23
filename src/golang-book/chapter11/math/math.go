package math

import "math"

// Average returns the average of a series of numbers
func Average(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	var total float64
	for _, x := range xs {
		total += x
	}

	return total / float64(len(xs))
}

// Min returns the minimal value of a series of numbers
func Min(xs []float64) float64 {
	min := math.Inf(1)
	for _, x := range xs {
		if x < min {
			min = x
		}
	}

	return min
}

// Max returns maximal value of a series of numbers
func Max(xs []float64) float64 {
	max := math.Inf(-1)
	for _, x := range xs {
		if x > max {
			max = x
		}
	}

	return max
}
