package utils

import "testing"

func TestMin(t *testing.T) {
	min := Min(0, -2, 10, 50, -4, 1)

	if min != -4 {
		t.Errorf("Min should be -4, got %d", min)
	}
}
