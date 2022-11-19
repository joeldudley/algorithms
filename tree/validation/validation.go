package validation

import (
	"epi/tree"
	"epi/tree/traversal"
)

// IsBinarySearchTree checks if the node is the root of a binary search tree.
func IsBinarySearchTree(node *tree.Node) bool {
	// We perform an in-order traversal, and check that it increases monotonically.
	inOrderTraversal := traversal.InOrderTraversal(node)
	for i := 1; i < len(inOrderTraversal); i++ {
		if inOrderTraversal[i] < inOrderTraversal[i-1] {
			return false
		}
	}
	return true
}
