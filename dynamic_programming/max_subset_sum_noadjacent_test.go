package dynamic_programming

import (
	"testing"
)

func TestMaxSubsetSumNoAdjacent_inputWithRepeatedItems(t *testing.T) {
	input := []int{75, 105, 120, 75, 90, 135}
	expected := 330

	result := MaxSubsetSumNoAdjacent(input)

	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}

}

func TestMaxSubsetSumNoAdjacent_inputArrayTwoDistinctItems(t *testing.T) {
	input := []int{1, 2}
	expected := 2

	result := MaxSubsetSumNoAdjacent(input)

	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestMaxSubsetSumNoAdjacent_inputArrayThreeDistinctItems(t *testing.T) {
	input := []int{1, 15, 3}
	expected := 15

	result := MaxSubsetSumNoAdjacent(input)

	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
