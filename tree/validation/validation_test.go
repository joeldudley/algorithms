package validation

import (
	"testing"

	"epi/tree"
)

var (
	// A small binary search tree.
	bstOne = &tree.Node{
		L:   &tree.Node{Val: 1},
		R:   &tree.Node{Val: 3},
		Val: 2,
	}

	// A larger binary search tree.
	bstTwo = &tree.Node{
		L: &tree.Node{
			L:   &tree.Node{Val: 1},
			R:   &tree.Node{Val: 3},
			Val: 2,
		},
		R: &tree.Node{
			L:   &tree.Node{Val: 5},
			R:   &tree.Node{Val: 7},
			Val: 6,
		},
		Val: 4,
	}

	// A binary search tree where a value on the left-hand side is equal to the root.
	bstThree = &tree.Node{
		L:   &tree.Node{Val: 1},
		R:   &tree.Node{Val: 2},
		Val: 1,
	}

	// A binary search tree where a value on the right-hand side is equal to the root.
	bstFour = &tree.Node{
		L:   &tree.Node{Val: 1},
		R:   &tree.Node{Val: 2},
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
	for _, bst := range []*tree.Node{bstOne, bstTwo, bstThree, bstFour} {
		if !IsBinarySearchTree(bst) {
			t.Errorf("incorrectly found BST to be non-BST")
		}
	}

	for _, bst := range []*tree.Node{notBstOne, notBstTwo, notBstThree} {
		if IsBinarySearchTree(bst) {
			t.Errorf("incorrectly found non-BST to be BST")
		}
	}
}
