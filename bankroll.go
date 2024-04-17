// Copyright (c) Paul Schick
// SPDX-License-Identifier: MPL-2.0

package GoWagerBrain

func BasicKellyFromDecimalOdds(probability, odds, kellySize float64) float64 {
	if probability < 0 || probability > 1 || odds <= 1 {
		return 0
	}

	b := odds - 1
	q := 1 - probability
	fullKelly := (b*probability - q) / b
	return fullKelly * kellySize
}

func BasicKellyFromAmericanOdds(probability, odds, kellySize float64) float64 {
	odds = DecimalOddsFromAmerican(odds)
	return BasicKellyFromDecimalOdds(probability, odds, kellySize)
}
