package validation

import (
	"testing"

	"epi/tree/data"

	"epi/tree"
)

func TestBinarySearchTreeValidation(t *testing.T) {
	for _, bst := range []*tree.Node{data.BstOne, data.BstTwo, data.BstThree, data.BstFour} {
		if !IsBinarySearchTree(bst) {
			t.Errorf("incorrectly found BST to be non-BST")
		}
	}

	for _, bst := range []*tree.Node{data.NotBstOne, data.NotBstTwo, data.NotBstThree} {
		if IsBinarySearchTree(bst) {
			t.Errorf("incorrectly found non-BST to be BST")
		}
	}
}
