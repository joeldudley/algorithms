package info

import (
	"epi/tree"
	"epi/tree/traversal"
)

// IsBinarySearchTree checks if the provided node is the root of a binary search tree.
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

// TreeHeight returns the height of the tree whose root is the provided node.
func TreeHeight(node *tree.Node) int {
	maxHeight := 1

	nodesToProcess := []*nodeAndHeight{{node: node, height: 1}}
	for i := 0; ; i++ {
		if i >= len(nodesToProcess) {
			break
		}
		nH := nodesToProcess[i]

		if nH.height > maxHeight {
			maxHeight = nH.height
		}
		if nH.node.L != nil {
			childNH := &nodeAndHeight{node: nH.node.L, height: nH.height + 1}
			nodesToProcess = append(nodesToProcess, childNH)
		}
		if nH.node.R != nil {
			childNH := &nodeAndHeight{node: nH.node.R, height: nH.height + 1}
			nodesToProcess = append(nodesToProcess, childNH)
		}
	}

	return maxHeight
}

// Pairs a node with its height in some tree.
type nodeAndHeight struct {
	node   *tree.Node
	height int
}
