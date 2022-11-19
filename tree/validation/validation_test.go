package validation

import (
	"testing"

	"epi/tree"
)

var (
	bst = &tree.Node{
		L:   &tree.Node{Val: 1},
		R:   &tree.Node{Val: 3},
		Val: 2,
	}

	notBstOne = &tree.Node{
		L:   &tree.Node{Val: 2}, // Higher than the root.
		R:   &tree.Node{Val: 3},
		Val: 1,
	}

	notBstTwo = &tree.Node{
		L:   &tree.Node{Val: 1},
		R:   &tree.Node{Val: 2}, // Lower than the root.
		Val: 3,
	}

	notBstThree = &tree.Node{
		L: &tree.Node{
			L:   &tree.Node{Val: 1},
			R:   &tree.Node{Val: 1}, // Higher than the root.
			Val: 2,
		},
		R:   &tree.Node{Val: 5},
		Val: 4,
	}
)

func Test(t *testing.T) {
	if !IsBinarySearchTree(bst) {
		t.Errorf("incorrectly found BST to be non-BST")
	}
	if IsBinarySearchTree(notBstOne) {
		t.Errorf("incorrectly found non-BST to be BST")
	}
	if IsBinarySearchTree(notBstTwo) {
		t.Errorf("incorrectly found non-BST to be BST")
	}
	if IsBinarySearchTree(notBstThree) {
		t.Errorf("incorrectly found non-BST to be BST")
	}
}
