package GoWagerBrain

import "testing"

func TestRound(t *testing.T) {
	expected := 1.23
	result := Round(1.23456, 2)
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}

	testNoDecimal := Round(1.555, 0)
	if testNoDecimal != 2 {
		t.Errorf("Expected 2, got %f", testNoDecimal)
	}

	testZeroValue := Round(0, 2)
	if testZeroValue != 0 {
		t.Errorf("Expected 0, got %f", testZeroValue)
	}

	testNegativeDecimals := Round(1.23456, -1)
	if testNegativeDecimals != 1.23456 {
		t.Errorf("Expected 1.23456, got %f", testNegativeDecimals)
	}
}
