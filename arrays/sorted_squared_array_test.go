package arrays

import (
	"reflect"
	"testing"
)

func TestSquareSortedArray(t *testing.T) {
	expected := []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}
	arrayIn := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := SquareSortedArray(arrayIn)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}

}
