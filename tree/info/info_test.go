package info

import (
	"testing"

	"epi/tree/data"

	"epi/tree"
)

func TestCanDetectBinarySearchTree(t *testing.T) {
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

func TestCanDetermineTreeHeight(t *testing.T) {
	// TODO - Expand tests once we can generate trees from arrays.
	treeHeight := TreeHeight(data.Tree)
	if treeHeight != 4 {
		t.Errorf("expected height of %d, got %d", 4, treeHeight)
	}
}
