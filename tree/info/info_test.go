package info

import (
	"math"
	"testing"

	"epi/tree/data"

	"epi/tree"
)

var (
	// A small binary search tree.
	bstOne = data.GenerateTree([]int{1, 2, 3})

	// A larger binary search tree.
	bstTwo = data.GenerateTree([]int{1, 2, 3, 4, 5, 6, 7})

	// A binary search tree where a value on the left-hand side is equal to the root.
	bstThree = data.GenerateTree([]int{1, 1, 2})

	// A binary search tree where a value on the right-hand side is equal to the root.
	bstFour = data.GenerateTree([]int{1, 2, 2})

	// A binary search tree where the left-hand side has a value higher than the root.
	notBstOne = data.GenerateTree([]int{2, 1, 3})

	// A binary search tree where the right-hand side has a value lower than the root.
	notBstTwo = data.GenerateTree([]int{1, 3, 2})

	// A binary search tree where the right-hand side has a value lower than the root, but nested
	// a level further down.
	notBstThree = data.GenerateTree([]int{1, 2, 3, 4, 1, 5, 6})
)

func TestCanDetectBinarySearchTree(t *testing.T) {
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

func TestCanDetermineTreeHeight(t *testing.T) {
	treeValues := []int{1, 3, 2, 5, 4, 11, 6, 10, 7, 9, 8}
	testTree := data.GenerateTree(treeValues)

	treeHeight := TreeHeight(testTree)
	expectedTreeHeight := treeHeightBalanced(len(treeValues))
	if treeHeight != expectedTreeHeight {
		t.Errorf("expected height of %d, got %d", expectedTreeHeight, treeHeight)
	}
}

// Calculates the height of a balanced tree with the provided number of nodes.
func treeHeightBalanced(numNodes int) int {
	return int(math.Ceil(math.Log2(float64(numNodes + 1))))
}
