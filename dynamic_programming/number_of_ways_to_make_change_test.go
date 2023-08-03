package dynamic_programming

import (
	"testing"
)

func TestNumberOfWaysToMakeChange(t *testing.T) {
	//fixture
	amount := 6
	coins := []int{1, 5}

	expected := 2

	// test
	result := NumberOfWaysToMakeChange(amount, coins)

	//assertion
	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}

}
