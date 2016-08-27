package tree

import (
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
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
			"caption": "",
			"type":    "",
		},
	)
	// This is a bit crazy, but hey, I didn't make the rules
	// http://stackoverflow.com/questions/28054913/modify-array-of-interface-golang
	//self.treeDispMap["nodes"] = treeRepNodesArr.(interface{})
	/*
		^ is needed if the present method is not a struct method(I
		don't know what else to call this thing).
		But if it's not a struct method, the above line raises and
		error and a simple assignment does the trick.
		TODO: Figure out what the fuck is this shit
	*/
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
					"caption": "",
				},
			)
		}

		if rExists {
			treeRepEdgesArr = append(
				treeRepEdgesArr,
				map[string]interface{}{
					"source":  root.id,
					"target":  ruuid,
					"caption": "",
				},
			)
		}

		//self.treeDispMap["edges"] = treeRepEdgesArr.(interface{})
		self.treeDispMap["edges"] = treeRepEdgesArr
	}

	return root.id, true
}
