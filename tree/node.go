package tree

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	//"strings"
)

type Node struct {
	id   string
	data *interface{}
	link map[string]*Node
}

// Setters and getters
func (n *Node) GetNodeData() *interface{} {
	return n.data
}

func (n *Node) GetNodeLinks() map[string]*Node {
	return n.link
}

func (n *Node) SetNodeData(data interface{}) {
	n.data = &data
}

func (n *Node) SetNodeLink(link map[string]*Node) {
	n.link = link
}

func (n *Node) isDirn(dirn string) bool {
	if n != nil {
		parentNode, isExists := n.link["parent"]
		if isExists && parentNode != nil {
			return parentNode.link[dirn] == n
		}
	}

	return false
}

func (n *Node) IsLeft() bool {
	return n.isDirn("left")
}

func (n *Node) IsRight() bool {
	return n.isDirn("right")
}

func (n *Node) String() string {
	// only display the keys which have values
	/*
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
	*/

	return fmt.Sprintf("Data: %v", *n.data)
}

func (n *Node) GetInfoString() string {
	return n.String()
}

// Creates a node
func CreateTreeNode(data *interface{}) *Node {
	return CreateNode(data, map[string]*Node{
		"left":   nil,
		"right":  nil,
		"parent": nil,
	})
}

func CreateNode(data *interface{}, linkMap map[string]*Node) *Node {
	uuid, err := uuid.NewV4()
	if err != nil {
		panic("Error generating a new UUID.")
	}

	return &Node{
		id:   uuid.String(),
		data: data,
		link: linkMap,
	}
}
