package tree

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"strings"
)

type customData struct {
	Num int
}
type customMap map[string]*Node

type Node struct {
	id   string
	data *interface{}
	link customMap
}

func (n *Node) String() string {
	// only display the keys which have values
	keysList := []string{
		"left", "right", "parent",
	}
	var validKeyList []string

	for _, key := range keysList {
		_, isExists := n.link[key]
		if isExists {
			validKeyList = append(
				validKeyList,
				key,
			)
		}
	}

	mapString := "No Valid maps"
	validKeyStr := strings.Join(validKeyList, " ")
	if validKeyStr != "" {
		mapString = "Valid maps - " + validKeyStr
	}

	return fmt.Sprintf("Data: %d\nMap: \n%s\n", n.data, mapString)
}

func (n *Node) GetInfoString() string {
	return self.String()
}

func (n *Node) AddChild(key string, childPtr *Node) {
	n.link[key] = childPtr
}

/*
Creates a Tree node that can be added to the tree
TODO; This abstraction is needed now because, not sure
how to  make this more general. Probably use something like
an interface{}
*/
func CreateTreeNode(n interface{}) *Node {
	//customData := createTreeData(n)
	//return makeNode(customData)
	return makeNode(&n)
}

// TODO: This is not needed. Remove this -
// CreateCustomData creates Tree data populated from the argument and
// returns a reference to the customData
func createTreeData(n int) *customData {
	return &customData{n}
}

// Creates a node
func makeNode(data *interface{}) *Node {
	uuid, err := uuid.NewV4()
	if err != nil {
		panic("Error generating a new UUID.")
	}
	return &Node{
		id:   uuid.String(),
		data: data,
		link: make(map[string]*Node),
	}
}
