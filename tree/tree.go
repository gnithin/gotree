package tree

type Tree struct {
	root      *Node
	len       int
	leavesLen int
}

// Module level function
func createTree() *Tree {
	return &Tree{
		root:      nil,
		len:       0,
		leavesLen: 0,
	}
}

func (t *Tree) addNode(n *Node) {
	if n == nil {
		panic("Cant handle empty nodes")
	}
	// Add a nodeN
	if t.root != nil {
		t.root = n
	}
	// Where to put this thing?
	panic("Not implemented this yet!!! :p")
}
