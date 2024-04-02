package helpers

import "math"

func Round(value float64, decimals int) float64 {
	shift := 1.0
	for i := 0; i < decimals; i++ {
		shift *= 10.0
	}
	return math.Round(value*shift) / shift
}
