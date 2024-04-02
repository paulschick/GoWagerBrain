package GoWagerBrain

// BreakEvenPercent
// stake: the amount of money you are willing to bet
// payout: the amount of money you will win if you win the bet (stake + profit)
// for example, if your odds are -110, then you risk $110 to win $100
// in this case, the stake is $110 and the payout is $210
// the break-even percentage is 110 / 210 = 0.5238
func BreakEvenPercent(stake, payout float64) float64 {
	return stake / payout
}

// Vig
// favStake: the amount of money you are willing to bet on the favorite
// favPayout: the amount of money you will win if the favorite wins the bet (favStake + profit)
// underStake: the amount of money you are willing to bet on the underdog
// underPayout: the amount of money you will win if the underdog wins the bet (underStake + profit)
// return the vig percentage paid to the bookmaker
func Vig(favStake, favPayout, underStake, underPayout float64) float64 {
	favBreakEven := BreakEvenPercent(favStake, favPayout)
	underBreakEven := BreakEvenPercent(underStake, underPayout)
	return (favBreakEven + underBreakEven) - 1
}
