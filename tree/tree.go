package tree

import (
	"encoding/json"
	"fmt"
)

type Tree struct {
	root      *Node
	len       int
	leavesLen int
	treeType  int
}

// Module level function
// TODO: Optional arguments for the treeType
func CreateTree() *Tree {
	return &Tree{
		root:      nil,
		len:       0,
		leavesLen: 0,
		treeType:  TREE_TYPE_BST,
	}
}

func (self *Tree) Insert(newVal int) {
	fmt.Println("************")
	fmt.Println("Adding - ", newVal)
	newNode := CreateTreeNode(newVal)
	self.addNode(newNode)
	fmt.Println("************")
}

func (self *Tree) addNode(newNode *Node) {
	if newNode == nil {
		panic("Cant handle empty nodes")
	}

	switch self.treeType {
	case TREE_TYPE_BST:
		self.addNodeBST(newNode)
	default:
		panic("Not impletemented this add node")
	}
}

func (self *Tree) addNodeBST(newNode *Node) {
	if self.root == nil {
		fmt.Println("Adding root")
		self.root = newNode
	} else {
		insertBST(self.root, newNode)
	}

	// Increment stuff
	self.len += 1
}

// Insert into BST
func insertBST(parent *Node, newNode *Node) {
	if parent == nil {
		panic("This shouldn't happen")
	}

	parentData := parent.data.Num
	nodeData := newNode.data.Num

	if parentData == nodeData {
		// There's no need to do anything
		return
	}

	dirn := "left"
	if parentData < nodeData {
		dirn = "right"
	}

	childNode, isExists := parent.link[dirn]
	if isExists {
		fmt.Println("Going ", dirn)
		insertBST(childNode, newNode)
	} else {
		// It needs to be inserted here
		parent.AddChild(dirn, newNode)
	}
}

// Creates a JSON output for the current tree as specified by alchemy
func (self *Tree) GetJSONTree() []byte {
	treeRep := make(map[string]interface{})

	var nodesArr []map[string]interface{}
	var edgesArr []map[string]interface{}
	treeRep["nodes"] = nodesArr
	treeRep["edges"] = edgesArr

	self.traverseAndPopulate(treeRep)
	fmt.Println(treeRep)

	treeJson, err := json.Marshal(treeRep)
	if err != nil {
		fmt.Println("Error marshalling the tree to json")
		fmt.Println(err)
	} else {
		return treeJson
	}

	return nil
}

func (self *Tree) traverseAndPopulate(treeRep map[string]interface{}) {
	postOrderTraverse(self.root, treeRep)
}

/*
- Traverse each node
- Add it self
- Traverse in post order, when visiting every child, add the node.
- When visiting every root, add the edge
*/
func postOrderTraverse(root *Node, treeRep map[string]interface{}) (string, bool) {
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

	treeRepNodesArrV := treeRep["nodes"].([]map[string]interface{})
	treeRepEdgesArrV := treeRep["edges"].([]map[string]interface{})

	treeRepNodesArr := &treeRepNodesArrV
	treeRepEdgesArr := &treeRepEdgesArrV

	// Adding the root node
	*treeRepNodesArr = append(
		*treeRepNodesArr,
		map[string]interface{}{
			"id":      root.id,
			"caption": "",
			"type":    "",
		},
	)

	// Go left
	luuid, lExists := postOrderTraverse(root.link["left"], treeRep)
	if lExists {
		*treeRepEdgesArr = append(
			*treeRepEdgesArr,
			map[string]interface{}{
				"source":  root.id,
				"target":  luuid,
				"caption": "",
			},
		)
	}

	// Go right
	ruuid, rExists := postOrderTraverse(root.link["right"], treeRep)
	if rExists {
		*treeRepEdgesArr = append(
			*treeRepEdgesArr,
			map[string]interface{}{
				"source":  root.id,
				"target":  ruuid,
				"caption": "",
			},
		)
	}

	fmt.Println("here")
	return root.id, true
}
