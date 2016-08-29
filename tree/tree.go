package tree

import (
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"strconv"
)

type Tree struct {
	root        *Node
	len         int
	leavesLen   int
	treeType    int
	id          string
	treeDispMap map[string]interface{}
}

// Module level function
// TODO: Optional arguments for the treeType
func CreateTree() *Tree {
	uuid, err := uuid.NewV4()
	if err != nil {
		panic("Error generating a new UUID.")
	}

	nodesArr := []map[string]interface{}{}
	edgesArr := []map[string]interface{}{}

	tMap := map[string]interface{}{
		"nodes": nodesArr,
		"edges": edgesArr,
	}

	return &Tree{
		root:        nil,
		len:         0,
		leavesLen:   0,
		treeType:    TREE_TYPE_BST,
		id:          uuid.String(),
		treeDispMap: tMap,
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
		self.insertBST(self.root, newNode)
	}

	// Increment stuff
	self.len += 1
}

// Creates a JSON output for the current tree as specified by alchemy
func (self *Tree) GetJSONTree() []byte {
	self.postOrderTraverse(self.root)
	fmt.Println(self.treeDispMap)

	var treeJson []byte
	var jsonErr error

	if PRETTY_PRINT_TREE {
		treeJson, jsonErr = json.MarshalIndent(self.treeDispMap, "", "    ")
	} else {
		treeJson, jsonErr = json.Marshal(self.treeDispMap)
	}

	if jsonErr != nil {
		fmt.Println("Error marshalling the tree to json")
		fmt.Println(jsonErr)
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
func (self *Tree) postOrderTraverse(root *Node) (string, bool) {
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
			"caption": strconv.Itoa(root.data.Num),
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
