package tree

// Node is a node in a tree data structure.
type Node struct {
	L   *Node
	R   *Node
	Val int
}
