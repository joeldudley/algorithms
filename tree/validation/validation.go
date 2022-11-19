package validation

import (
	"epi/tree"
)

// IsBinarySearchTree gets each node, and checks it against all its right and all its left children.
// There is no attempt to cache whether individual subtrees are BSTs - for each node, we recheck
// the entire subtree.
// TODO - Improve on this brute-force approach.
func IsBinarySearchTree(node *tree.Node) bool {
	nodesToProcess := []*tree.Node{node}

	// We build a list of all the nodes.
	for i := 0; i < len(nodesToProcess); i++ {
		currentNode := nodesToProcess[i]

		if currentNode.L != nil {
			nodesToProcess = append(nodesToProcess, currentNode.L)
		}
		if currentNode.R != nil {
			nodesToProcess = append(nodesToProcess, currentNode.R)
		}
	}

	// For each node...
	for i := 0; i < len(nodesToProcess); i++ {
		currentNode := nodesToProcess[i]

		// ...we build a list of its left-hand nodes, and check none is larger...
		var leftChildNodes []*tree.Node
		if currentNode.L != nil {
			leftChildNodes = append(leftChildNodes, currentNode.L)
		}

		for j := 0; j < len(leftChildNodes); j++ {
			childNode := leftChildNodes[j]
			if childNode.Val > currentNode.Val {
				return false
			}
			if childNode.L != nil {
				leftChildNodes = append(leftChildNodes, childNode.L)
			}
			if childNode.R != nil {
				leftChildNodes = append(leftChildNodes, childNode.R)
			}
		}

		// ...we build a list of its right-hand nodes, and check none is smaller...
		var rightChildNodes []*tree.Node
		if currentNode.R != nil {
			rightChildNodes = append(rightChildNodes, currentNode.R)
		}

		for j := 0; j < len(rightChildNodes); j++ {
			childNode := rightChildNodes[j]
			if childNode.Val < currentNode.Val {
				return false
			}
			if childNode.L != nil {
				rightChildNodes = append(rightChildNodes, childNode.L)
			}
			if childNode.R != nil {
				rightChildNodes = append(rightChildNodes, childNode.R)
			}
		}
	}

	return true
}
