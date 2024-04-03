package GoWagerBrain

import "math"

func Round(value float64, decimals int) float64 {
	if value == 0 {
		return 0
	}

	if decimals < 0 {
		return value
	}

	if decimals > 0 && value == math.Trunc(value) {
		return value
	}

	shift := 1.0
	for i := 0; i < decimals; i++ {
		shift *= 10.0
	}

	return math.Round(value*shift) / shift
}
