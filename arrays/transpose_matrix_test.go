package arrays

import (
	"reflect"
	"testing"
)

func TestTransposeMaxtrixArray(t *testing.T) {
	expected := [][]int{{1}, {2}}
	input := [][]int{{1, 2}}

	result := TransposeMatrix(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}

}
