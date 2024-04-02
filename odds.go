package GoWagerBrain

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func AmericanOddsFromDecimal(odds float64) float64 {
	if odds > 2.0 {
		return (odds - 1) * 100
	}
	return -100 / (odds - 1)
}

func AmericanOddsFromFractional(odds string) (float64, error) {
	if !strings.Contains(odds, "/") {
		return 0, fmt.Errorf("invalid fractional odds: %s", odds)
	}
	values := strings.Split(odds, "/")
	if len(values) != 2 {
		return 0, fmt.Errorf("invalid fractional odds: %s", odds)
	}

	numerator := values[0]
	n, err := strconv.Atoi(numerator)
	if err != nil {
		return 0, fmt.Errorf("invalid fractional odds %s: %w", odds, err)
	}
	denominator := values[1]
	d, err := strconv.Atoi(denominator)
	if err != nil {
		return 0, fmt.Errorf("invalid fractional odds %s: %w", odds, err)
	}
	nf := float64(n)
	df := float64(d)

	if n > d {
		return nf / df * 100, nil
	}
	return -100.0 / (nf / df), nil
}

func DecimalOddsFromAmerican(odds float64) float64 {
	if odds >= 100 {
		return math.Abs(1 + (odds / 100))
	}
	return 100/math.Abs(odds) + 1
}

func DecimalOddsFromFractional(odds string) float64 {
	if !strings.Contains(odds, "/") {
		return 0
	}
	values := strings.Split(odds, "/")
	if len(values) != 2 {
		return 0
	}

	numerator := values[0]
	n, err := strconv.Atoi(numerator)
	if err != nil {
		return 0
	}
	denominator := values[1]
	d, err := strconv.Atoi(denominator)
	if err != nil {
		return 0
	}
	nf := float64(n)
	df := float64(d)

	return 1 + nf/df
}

func ParlayOddsFromDecimal(odds []float64) float64 {
	oddsProduct := 1.0
	for _, odd := range odds {
		oddsProduct *= odd
	}
	return oddsProduct
}

func ParlayOddsFromAmerican(odds []float64) float64 {
	oddsProduct := 1.0
	for _, odd := range odds {
		oddsProduct *= DecimalOddsFromAmerican(odd)
	}
	return oddsProduct
}
