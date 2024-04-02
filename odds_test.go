package GoWagerBrain

import (
	"testing"
)

func TestAmericanOddsFromDecimal(t *testing.T) {
	expectedAndOddsMap := map[float64]float64{
		-10000: 1.01,
		-400:   1.25,
		-500:   1.20,
		-455:   1.22,
		-250:   1.40,
		-120:   1.83,
		-125:   1.80,
		140:    2.40,
		180:    2.80,
	}

	run := func(expected float64, odds float64) bool {
		actual := Round(AmericanOddsFromDecimal(odds), 0)
		exp := Round(expected, 0)
		if actual != exp {
			return false
		}
		return true
	}

	for expected, odds := range expectedAndOddsMap {
		if !run(expected, odds) {
			t.Errorf("expected %f, got %f", expected, AmericanOddsFromDecimal(odds))
		}
	}
}

func TestAmericanOddsFromFractional(t *testing.T) {
	runList := map[string]float64{
		"1/100": -10000,
		"4/11":  -275,
		"23/20": 115,
		"6/4":   150,
		"8/11":  -137.5,
		"11/8":  137.5,
		"10/3":  333.33,
	}
	run := func(in string, exp float64) bool {
		actual, err := AmericanOddsFromFractional(in)
		actualRound := Round(actual, 2)
		expectedRound := Round(exp, 2)
		if err != nil {
			return false
		}
		if actualRound != expectedRound {
			return false
		}
		return true
	}

	for in, exp := range runList {
		if !run(in, exp) {
			result, _ := AmericanOddsFromFractional(in)
			t.Errorf("expected %f, got %f", exp, result)
		}
	}
}

func TestDecimalOddsFromAmerican(t *testing.T) {
	testParams := map[float64]float64{
		333.33:  4.33,
		275:     3.75,
		-333.33: 1.30,
		-250:    1.40,
		-187.5:  1.53,
		1800:    19.0,
	}
	run := func(american, expectedDecimal float64) bool {
		actual := Round(DecimalOddsFromAmerican(american), 2)
		expected := Round(expectedDecimal, 2)
		if actual != expected {
			return false
		}
		return true
	}

	for american, expectedDecimal := range testParams {
		if !run(american, expectedDecimal) {
			t.Errorf("expected %f, got %f", expectedDecimal, DecimalOddsFromAmerican(american))
		}
	}
}

func TestDecimalOddsFromFractional(t *testing.T) {
	params := map[string]float64{
		"11/1": 12,
		"8/13": 1.62,
		"1/1":  2,
		"4/7":  1.57,
		"8/11": 1.73,
	}
	run := func(toTest string, expected float64) bool {
		actual := Round(DecimalOddsFromFractional(toTest), 2)
		expectedRound := Round(expected, 2)
		if actual != expectedRound {
			return false
		}
		return true
	}

	for toTest, expected := range params {
		if !run(toTest, expected) {
			t.Errorf("expected %f, got %f", expected, DecimalOddsFromFractional(toTest))
		}
	}
}

func TestParlayOddsFromDecimal(t *testing.T) {
	testVals := []float64{2.1, 1.5, 2.2, 1.2}
	expected := 8.32
	result := Round(ParlayOddsFromDecimal(testVals), 2)
	if result != expected {
		t.Errorf("expected %f, got %f", expected, result)
	}
}

func TestParlayOddsFromAmerican(t *testing.T) {
	odds := []float64{110, -200, 120, -500}
	expected := 8.32
	result := Round(ParlayOddsFromAmerican(odds), 2)
	if result != expected {
		t.Errorf("expected %f, got %f", expected, result)
	}
}
