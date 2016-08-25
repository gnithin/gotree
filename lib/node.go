package lib

import "fmt"

type Node struct {
	// TODO: The data needs to be a general container.
	// It should hold anything. Could be anything
	data int

	// TODO: Think about what type the keys of the map needs to be.
	// Making it a string seems to be a simple cast. Myabe there are
	// situations where it shouldn't be a string
	next map[string]*Node
}

func MakeNode(data int) *Node {
	// TODO: Needs to be changed whenever the Node is changed
	return &Node{
		data: 0,
		next: make(map[string]*Node),
	}
}

func (n *Node) String() string {
	return fmt.Sprintf("Data: %d\nMap: %s\n", n.data, n.next)
}
