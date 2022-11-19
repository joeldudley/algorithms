package data

import "epi/tree"

// Various trees. Used in tests.
var (
	// Tree is a random tree with no particular properties.
	Tree = GenerateTree([]int{1, 3, 2, 5, 4, 11, 6, 10, 7, 9, 8})

	// BstOne is a small binary search tree.
	BstOne = GenerateTree([]int{1, 2, 3})

	// BstTwo is a larger binary search tree.
	BstTwo = GenerateTree([]int{1, 2, 3, 4, 5, 6, 7})

	// BstThree is a binary search tree where a value on the left-hand side is equal to the root.
	BstThree = GenerateTree([]int{1, 1, 2})

	// BstFour is a binary search tree where a value on the right-hand side is equal to the root.
	BstFour = GenerateTree([]int{1, 2, 2})

	// NotBstOne is a binary search tree where the left-hand side has a value higher than the root.
	NotBstOne = GenerateTree([]int{2, 1, 3})

	// NotBstTwo is a binary search tree where the right-hand side has a value lower than the root.
	NotBstTwo = GenerateTree([]int{1, 3, 2})

	// NotBstThree is a binary search tree where the right-hand side has a value lower than the
	// root, but nested a level further down.
	NotBstThree = GenerateTree([]int{1, 2, 3, 4, 1, 5, 6})
)

// GenerateTree takes an array of values, uses the middle value as the root, then recursively
// builds a tree for the left-hand side using the half of the array before the middle value, and
// for the right-hand side using the half of the array after the middle value. A sorted array will
// produce a binary search tree.
func GenerateTree(values []int) *tree.Node {
	if len(values) == 0 {
		return nil
	}

	midPoint := len(values) / 2
	rootValue := values[len(values)/2]
	leftNode := GenerateTree(values[0:midPoint])
	rightNode := GenerateTree(values[midPoint+1:])

	return &tree.Node{
		Val: rootValue,
		L:   leftNode,
		R:   rightNode,
	}
}
