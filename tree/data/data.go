package data

import "epi/tree"

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
