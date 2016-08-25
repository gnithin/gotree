package lib

import (
	//"bytes"
	"fmt"
)

type customMap map[string]*Node

type Node struct {
	// TODO: The data needs to be a general container.
	// It should hold anything. Could be anything
	data int

	// TODO: Think about what type the keys of the map needs to be.
	// Making it a string seems to be a simple cast. Myabe there are
	// situations where it shouldn't be a string
	next customMap
}

func MakeNode(data int) *Node {
	// TODO: Needs to be changed whenever the Node is changed
	return &Node{
		data: data,
		next: make(map[string]*Node),
	}
}

func (n *Node) String() string {
	return fmt.Sprintf("Data: %d\nMap: {\n%s\n}\n", n.data, n.next)
}

func (n *Node) AddChild(key string, childPtr *Node) {
	n.next[key] = childPtr
}

// TODO: This is a fucking mess
// Probably make it into a JSON
/*
func (m customMap) String() string {
	var buffer bytes.Buffer

	for k, v := range m {
		buffer.WriteString(fmt.Sprintf("{Key_%s : Value_{%s}}", k, v))
		fmt.Println("asdasdAS : ", k)
		fmt.Println("asdasd", v)
	}
	if buffer.Len() != 0 {
		return buffer.String()
	}
	return ""
}
*/
