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
}
