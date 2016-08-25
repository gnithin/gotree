package tree

import (
	//"bytes"
	"fmt"
)

type customMap map[string]*Node

type TreeData struct {
	Num int
}

// CreateTreeData creates Tree data populated from the argument and
// returns a reference to the TreeData
func CreateTreeData(n int) *TreeData {
	return &TreeData{n}
}

type Node struct {
	data *TreeData
	link customMap
}

func MakeNode(data *TreeData) *Node {
	// TODO: Needs to be changed whenever the Node is changed
	return &Node{
		data: data,
		link: make(map[string]*Node),
	}
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
