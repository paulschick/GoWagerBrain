package GoWagerBrain

import (
	"testing"
)

type kellySetup struct {
	Odds      float64
	Prob      float64
	Exp       float64
	KellySize float64
}

func runHelper(setup kellySetup, testFn func(a, b, c float64) float64) (float64, bool) {
	actual := testFn(setup.Prob, setup.Odds, setup.KellySize)
	actual = Round(actual, 4)
	expected := Round(setup.Exp, 4)
	if actual != expected {
		return actual, false
	}
	return actual, true
}

func TestBasicKellyFromDecimalOdds(t *testing.T) {
	kellySetups := []kellySetup{
		{2.1, 0.55, 0.1409, 1.0},
		{2.1, 0.55, 0.0705, 0.5},
		{1.4, 0.75, 0.1250, 1.0},
		{1.4, 0.75, 0.0750, 0.6},
		{1, 0.5, 0, 1.0},
		{0, 0.5, 0, 1.0},
	}

	for _, setup := range kellySetups {
		received, ok := runHelper(setup, BasicKellyFromDecimalOdds)
		if !ok {
			t.Errorf("Expected %f, got %f", setup.Exp, received)
		}
	}
}

func TestBasicKellyFromAmericanOdds(t *testing.T) {
	kellySetups := []kellySetup{
		{110, 0.55, 0.1409, 1.0},
		{110, 0.55, 0.0705, 0.5},
		{-250, 0.75, 0.1250, 1.0},
		{-250, 0.75, 0.0750, 0.6},
	}

	for _, setup := range kellySetups {
		received, ok := runHelper(setup, BasicKellyFromAmericanOdds)
		if !ok {
			t.Errorf("Expected %f, got %f", setup.Exp, received)
		}
	}
}
