package traversal

import "epi/tree"

// InOrderTraversal returns the sequence of values produced by traversing the tree in-order.
// TODO - Implement a non-recursive solution.
func InOrderTraversal(node *tree.Node) []int {
	var inOrderValues []int

	if node.L != nil {
		leftInOrderValues := InOrderTraversal(node.L)
		inOrderValues = append(inOrderValues, leftInOrderValues...)
	}
	inOrderValues = append(inOrderValues, node.Val)
	if node.R != nil {
		rightInOrderValues := InOrderTraversal(node.R)
		inOrderValues = append(inOrderValues, rightInOrderValues...)
	}

	return inOrderValues
}
