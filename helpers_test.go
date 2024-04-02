package GoWagerBrain

import "testing"

func TestRound(t *testing.T) {
	expected := 1.23
	result := Round(1.23456, 2)
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}
