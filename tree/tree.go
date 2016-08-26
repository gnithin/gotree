package tree

import (
	"encoding/json"
	"fmt"
)

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

// Creates a JSON output for the current tree as specified by alchemy
func (t *Tree) GetJSONTree() []byte {
	/*
		The resulting structure needs to be of this format -
		{
			"nodes": [
				{
					"id" : <int>,
					"caption" : "",
					"type" : ""
				},
				...
			],
			"edges" : [
				{
					"source" : <int>,
					"target" : <int>,
					"caption" :  ""
				},
				...
			]
		}
	*/
	treeRep := make(map[string]interface{})

	var nodesArr []map[string]interface{}
	var edgesArr []map[string]interface{}
	treeRep["nodes"] = nodesArr
	treeRep["edges"] = edgesArr

	popStatus := t.traverseAndPopulate(&treeRep)

	if !popStatus {
		fmt.Println("Error populating the tree representation")
	} else {
		treeJson, err := json.Marshal(treeRep)
		if err != nil {
			fmt.Println("Error marshalling the tree to json")
			fmt.Println(err)
		} else {
			return treeJson
		}
	}

	return nil
}

func (t *Tree) traverseAndPopulate(treeRep *map[string]interface{}) bool {
	return false
}
