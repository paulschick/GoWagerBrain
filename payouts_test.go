// Copyright (c) Paul Schick
// SPDX-License-Identifier: MPL-2.0

package GoWagerBrain

import "testing"

type payoutsTestParams struct {
	stake, odds, expected float64
}

func TestAmericanPayouts(t *testing.T) {
	stake := 100.0
	tests := []payoutsTestParams{
		{stake, 100, 200},
		{stake, -100, 200},
		{stake, 110, 210},
		{stake, -110, 190.91},
		{stake, 120, 220},
		{stake, -120, 183.33},
	}
	for _, test := range tests {
		actual := Round(AmericanPayouts(test.stake, test.odds), 2)
		expected := Round(test.expected, 2)
		if actual != expected {
			t.Errorf("AmericanPayouts(%f, %f) = %f; want %f", test.stake, test.odds, actual, expected)
		}
	}
}

func TestDecimalPayouts(t *testing.T) {
	stake := 100.0
	tests := []payoutsTestParams{
		{stake, 2.0, 200},
		{stake, 1.0, 100},
		{stake, 1.1, 110},
		{stake, 1.91, 191},
		{stake, 2.2, 220},
		{stake, 1.83, 183},
	}
	for _, test := range tests {
		actual := Round(DecimalPayouts(test.stake, test.odds), 2)
		expected := Round(test.expected, 2)
		if actual != expected {
			t.Errorf("DecimalPayouts(%f, %f) = %f; want %f", test.stake, test.odds, actual, expected)
		}
	}
}

func TestAmericanProfit(t *testing.T) {
	stake := 100.0
	tests := []payoutsTestParams{
		{stake, 100, 100},
		{stake, -100, 100},
		{stake, 110, 110},
		{stake, -110, 90.91},
		{stake, 120, 120},
		{stake, -120, 83.33},
	}
	for _, test := range tests {
		actual := Round(AmericanProfit(test.stake, test.odds), 2)
		expected := Round(test.expected, 2)
		if actual != expected {
			t.Errorf("AmericanProfit(%f, %f) = %f; want %f", test.stake, test.odds, actual, expected)
		}
	}
}

func TestDecimalProfit(t *testing.T) {
	stake := 100.0
	tests := []payoutsTestParams{
		{stake, 2.0, 100},
		{stake, 1.0, 0},
		{stake, 1.1, 10},
		{stake, 1.91, 91},
		{stake, 2.2, 120},
		{stake, 1.83, 83},
	}
	for _, test := range tests {
		actual := Round(DecimalProfit(test.stake, test.odds), 2)
		expected := Round(test.expected, 2)
		if actual != expected {
			t.Errorf("DecimalProfit(%f, %f) = %f; want %f", test.stake, test.odds, actual, expected)
		}
	}
}
