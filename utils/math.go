package utils

import (
	"math"
)

func Min(numbers ...int) int {
	min := math.MaxInt
	for _, number := range numbers {
		if number < min {
			min = number
		}
	}
	return min
}
