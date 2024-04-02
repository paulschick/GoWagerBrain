package GoWagerBrain

func BasicKellyFromDecimalOdds(probability, odds, kellySize float64) float64 {
	b := odds - 1
	q := 1 - probability
	fullKelly := (b*probability - q) / b
	return fullKelly * kellySize
}

func BasicKellyFromAmericanOdds(probability, odds, kellySize float64) float64 {
	odds = DecimalOddsFromAmerican(odds)
	return BasicKellyFromDecimalOdds(probability, odds, kellySize)
}
