package GoWagerBrain

import "math"

type OddsType string

const (
	Decimal  OddsType = "Decimal"
	American OddsType = "American"
)

func AmericanPayouts(stake, odds float64) float64 {
	if odds > 0 {
		return (stake * (odds / 100)) + stake
	}
	val := math.Abs(stake / (odds / 100))
	return val + stake
}

func DecimalPayouts(stake, odds float64) float64 {
	return stake * odds
}

func AmericanProfit(stake, odds float64) float64 {
	if odds > 0 {
		return stake * (odds / 100)
	}
	return math.Abs(stake / (odds / 100))
}

func DecimalProfit(stake, odds float64) float64 {
	return stake * (odds - 1)
}

func GetPayout(stake, odds float64, oddsType OddsType) float64 {
	switch oddsType {
	case Decimal:
		return DecimalPayouts(stake, odds)
	case American:
		return AmericanPayouts(stake, odds)
	}
	return 0
}

func GetProfit(stake, odds float64, oddsType OddsType) float64 {
	switch oddsType {
	case Decimal:
		return DecimalProfit(stake, odds)
	case American:
		return AmericanProfit(stake, odds)
	}
	return 0
}
