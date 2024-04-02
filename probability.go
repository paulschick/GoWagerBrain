package GoWagerBrain

func DecimalImpliedWinProb(odds float64, decimals int) float64 {
	v := 1 / odds
	return Round(v, decimals)
}

func AmericanImpliedWinProb(odds float64, decimals int) float64 {
	dOdds := DecimalOddsFromAmerican(odds)
	return DecimalImpliedWinProb(dOdds, decimals)
}
