package dynamic_programming

import (
	"testing"
)

func TestKadane_EmptyArray(t *testing.T) {
	input := []int{}
	expected := 0

	result := KadanesAlgorithm(input)

	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}

}

func TestKadane_OneElementArray(t *testing.T) {
	input := []int{10}
	expected := 10

	result := KadanesAlgorithm(input)

	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}

}

func TestKadane(t *testing.T) {
	input := []int{3, 5, -9, 1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1, -5, 4}
	expected := 19

	result := KadanesAlgorithm(input)

	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}

}
