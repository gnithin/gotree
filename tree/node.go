package tree

import (
	//"bytes"
	"fmt"
)

type customData struct {
	Num int
}
type customMap map[string]*Node

type Node struct {
	data *customData
	link customMap
}

// TODO: Stackoverflow when using cycles, as when using parent key
func (n *Node) String() string {
	return fmt.Sprintf("Data: %d\nMap: {\n%s\n}\n", n.data, n.link)
}

func (n *Node) AddChild(key string, childPtr *Node) {
	n.link[key] = childPtr
	/*
		_, isExists := childPtr.link["parent"]
		if !isExists {
			childPtr.link["parent"] = n
		}
	*/
}

/*
Creates a Tree node that can be added to the tree
TODO; This abstraction is needed now because, not sure
how to  make this more general. Probably use something like
an interface{}
*/
func CreateTreeNode(n int) *Node {
	customData := createTreeData(n)
	return makeNode(customData)
}

// CreateCustomData creates Tree data populated from the argument and
// returns a reference to the customData
func createTreeData(n int) *customData {
	return &customData{n}
}

// Creates a node
func makeNode(data *customData) *Node {
	// TODO: Needs to be changed whenever the Node is changed
	return &Node{
		data: data,
		link: make(map[string]*Node),
	}
}
