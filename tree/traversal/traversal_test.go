package traversal

import (
	"reflect"
	"testing"

	"epi/tree/generator"
)

func TestInOrderTraversal(t *testing.T) {
	treeValues := []int{1, 3, 2, 5, 4, 11, 6, 10, 7, 9, 8}
	testTree := generator.GenerateTree(treeValues)

	traversal := InOrderTraversal(testTree)
	if !reflect.DeepEqual(traversal, treeValues) {
		t.Errorf("in-order traversal order was incorrect")
	}
}
