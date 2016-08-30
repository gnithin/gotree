package tree

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"strings"
)

type Node struct {
	id   string
	data *interface{}
	link map[string]*Node
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
	return n.String()
}

// Creates a node
func CreateTreeNode(data *interface{}) *Node {
	uuid, err := uuid.NewV4()
	if err != nil {
		panic("Error generating a new UUID.")
	}

	return &Node{
		id:   uuid.String(),
		data: data,
		link: map[string]*Node{
			"left":   nil,
			"right":  nil,
			"parent": nil,
		},
	}
}
