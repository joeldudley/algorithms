package data

import "epi/tree"

// Various trees. Used in tests.
var (
	// Tree is a random tree with no particular properties.
	Tree = &tree.Node{
		L: &tree.Node{
			L: &tree.Node{
				L:   &tree.Node{Val: 1},
				R:   &tree.Node{Val: 2},
				Val: 3,
			},
			R:   &tree.Node{Val: 4},
			Val: 5,
		},
		R: &tree.Node{
			L: &tree.Node{Val: 6},
			R: &tree.Node{
				L:   &tree.Node{Val: 7},
				R:   &tree.Node{Val: 8},
				Val: 9,
			},
			Val: 10,
		},
		Val: 11,
	}

	// BstOne is a small binary search tree.
	BstOne = &tree.Node{
		L:   &tree.Node{Val: 1},
		R:   &tree.Node{Val: 3},
		Val: 2,
	}

	// BstTwo is a larger binary search tree.
	BstTwo = &tree.Node{
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

	// BstThree is a binary search tree where a value on the left-hand side is equal to the root.
	BstThree = &tree.Node{
		L:   &tree.Node{Val: 1},
		R:   &tree.Node{Val: 2},
		Val: 1,
	}

	// BstFour is a binary search tree where a value on the right-hand side is equal to the root.
	BstFour = &tree.Node{
		L:   &tree.Node{Val: 1},
		R:   &tree.Node{Val: 2},
		Val: 2,
	}

	NotBstOne = &tree.Node{
		L:   &tree.Node{Val: 2}, // Higher than the root.
		R:   &tree.Node{Val: 3},
		Val: 1,
	}

	NotBstTwo = &tree.Node{
		L:   &tree.Node{Val: 1},
		R:   &tree.Node{Val: 2}, // Lower than the root.
		Val: 3,
	}

	NotBstThree = &tree.Node{
		L: &tree.Node{
			L:   &tree.Node{Val: 1},
			R:   &tree.Node{Val: 1}, // Higher than the root.
			Val: 2,
		},
		R:   &tree.Node{Val: 5},
		Val: 4,
	}
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
