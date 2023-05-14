package greedy

import (
	"testing"
)

func TestMinimumWaitTime(t *testing.T) {
	expected := 17
	arrayIn := []int{3, 2, 1, 2, 6}

	result := MinimumWaitingTime(arrayIn)

	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
