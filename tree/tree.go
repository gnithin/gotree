package tree

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Bucket of common functionality available across the different trees
type Tree interface {
	Insert(...interface{}) bool
	Pop() (*interface{}, bool)
}

type BaseTree struct {
	id          string
	root        *Node
	len         int
	leavesLen   int
	treeDispMap map[string]interface{}
	comparator  *func(obj1, obj2 *interface{}) int
}

func (self *BaseTree) IsEmpty() bool {
	return self.len == 0
}

func (self *BaseTree) GetLen() int {
	return self.len
}

func (self *BaseTree) checkTypeForComparator(node *Node) bool {
	// Just check if there's a comparator specified
	// Find the type. If the type is either a string or an int,
	// add the default comparator. Else raise error
	// Type assertion only provided for some things
	switch (*node.data).(type) {
	case int:
		temp := intComparator
		self.comparator = &temp
	case string:
		temp := stringComparator
		self.comparator = &temp
	default:
		if self.comparator == nil {
			//debug("City on Fire... City on Fire... Mischief!! Mischief!!")
			//panic("Need to specify comparator if the type is not string or int")
			return false
		}
	}
	return true
}

// Creates a JSON output for the current tree as specified by alchemy
func (self *BaseTree) GetJSONTree() []byte {
	self.postOrderTraverse(self.root)
	debug(self.treeDispMap)

	var treeJson []byte
	var jsonErr error

	if PRETTY_PRINT_TREE {
		treeJson, jsonErr = json.MarshalIndent(self.treeDispMap, "", "    ")
	} else {
		treeJson, jsonErr = json.Marshal(self.treeDispMap)
	}

	if jsonErr != nil {
		debug("Error marshalling the tree to json")
		debug(jsonErr)
	} else {
		return treeJson
	}

	return nil
}

/*
- Traverse each node
- Add it self
- Traverse in post order, when visiting every child, add the node.
- When visiting every root, add the edge
*/
func (self *BaseTree) postOrderTraverse(root *Node) (string, bool) {
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
	if root == nil {
		return "", false
	}

	treeRepNodesArr := self.treeDispMap["nodes"].([]map[string]interface{})
	// Adding the root node
	treeRepNodesArr = append(
		treeRepNodesArr,
		map[string]interface{}{
			"id":      root.id,
			"caption": root.GetInfoString(),
			"type":    "",
		},
	)
	// This is a bit crazy, but hey, I didn't make the rules
	// http://stackoverflow.com/questions/28054913/modify-array-of-interface-golang
	//self.treeDispMap["nodes"] = treeRepNodesArr.(interface{})
	self.treeDispMap["nodes"] = treeRepNodesArr

	// Go left
	luuid, lExists := self.postOrderTraverse(root.link["left"])

	// Go right
	ruuid, rExists := self.postOrderTraverse(root.link["right"])

	// Adding the edges
	if lExists || rExists {
		treeRepEdgesArr := self.treeDispMap["edges"].([]map[string]interface{})
		if lExists {
			treeRepEdgesArr = append(
				treeRepEdgesArr,
				map[string]interface{}{
					"source":  root.id,
					"target":  luuid,
					"caption": "left",
				},
			)
		}

		if rExists {
			treeRepEdgesArr = append(
				treeRepEdgesArr,
				map[string]interface{}{
					"source":  root.id,
					"target":  ruuid,
					"caption": "right",
				},
			)
		}

		//self.treeDispMap["edges"] = treeRepEdgesArr.(interface{})
		self.treeDispMap["edges"] = treeRepEdgesArr
	}

	return root.id, true
}

type BaseSequentialTree struct {
	BaseTree
	nodeArr []*Node
	maxSize int
}

func (self *BaseSequentialTree) String() string {
	// Only print the non empty values
	// TODO: this can be made super fast
	var respList []string
	for i := 0; i < self.maxSize; i++ {
		if self.nodeArr[i] != nil && self.nodeArr[i].data != nil {
			respList = append(
				respList,
				strconv.Itoa(i)+":"+self.nodeArr[i].GetInfoString(),
			)
		}
	}
	resp := fmt.Sprintf(strings.Join(respList, "\n"))
	return resp
}

func (self *BaseSequentialTree) getParentIndex(childIndex int) int {
	if childIndex < 0 {
		panic("Tree index cannot be < 0")
	}

	return (childIndex - 1) / 2
}

func (self *BaseSequentialTree) getChildIndex(parentIndex int, isLeft bool) int {
	if parentIndex < 0 {
		panic("Tree index cannot be < 0")
	}

	inc := 1
	if !isLeft {
		inc = 2
	}
	return (2 * parentIndex) + inc
}

func (self *BaseSequentialTree) isLeftChild(childIndex int) bool {
	if childIndex <= 0 {
		panic("Tree index cannot be <= 0")
	}

	return (childIndex % 2) != 0
}

func (self *BaseSequentialTree) IsFull() bool {
	return self.len >= self.maxSize
}
