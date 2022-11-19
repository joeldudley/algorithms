package traversal

import (
	"reflect"
	"testing"

	"epi/tree/data"
)

func TestInOrderTraversal(t *testing.T) {
	calculatedTraversal := InOrderTraversal(data.Tree)
	expectedTraversal := []int{1, 3, 2, 5, 4, 11, 6, 10, 7, 9, 8}

	if !reflect.DeepEqual(calculatedTraversal, expectedTraversal) {
		t.Errorf("in-order traversal was incorrect")
	}
}
