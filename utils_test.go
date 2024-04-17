// Copyright (c) Paul Schick
// SPDX-License-Identifier: MPL-2.0

package GoWagerBrain

import (
	"testing"
)

func TestVig(t *testing.T) {
	// +117 stake 100 payout 217.00 -> fav
	// -154 stake 100 payout 164.94 -> under
	expected := Round(0.0671, 4)
	result := Round(Vig(100, 217.00, 100, 164.94), 4)
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}

	testZero := Vig(10000, 0, 0, 1000)
	if testZero != 0 {
		t.Errorf("Expected 0, got %f", testZero)
	}
}
