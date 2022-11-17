package fourteen

import "testing"

var (
	bst = &Node{
		L:   &Node{Val: 1},
		R:   &Node{Val: 3},
		Val: 2,
	}

	notBstOne = &Node{
		L:   &Node{Val: 2}, // Higher than the root.
		R:   &Node{Val: 3},
		Val: 1,
	}

	notBstTwo = &Node{
		L:   &Node{Val: 1},
		R:   &Node{Val: 2}, // Lower than the root.
		Val: 3,
	}

	notBstThree = &Node{
		L: &Node{
			L:   &Node{Val: 1},
			R:   &Node{Val: 1}, // Higher than the root.
			Val: 2,
		},
		R:   &Node{Val: 5},
		Val: 4,
	}
)

func Test(t *testing.T) {
	if !IsBstBruteForce(bst) {
		t.Errorf("evaluated BST as non-BST")
	}
	if IsBstBruteForce(notBstOne) {
		t.Errorf("evaluated non-BST as BST")
	}
	if IsBstBruteForce(notBstTwo) {
		t.Errorf("evaluated non-BST as BST")
	}
	if IsBstBruteForce(notBstThree) {
		t.Errorf("evaluated non-BST as BST")
	}
}

// IsBstBruteForce gets each node, and checks it against all its right and all its left children.
// There is no attempt to cache whether individual subtrees are BSTs.
func IsBstBruteForce(node *Node) bool {
	nodesToProcess := []*Node{node}

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
		var leftChildNodes []*Node
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
		var rightChildNodes []*Node
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
