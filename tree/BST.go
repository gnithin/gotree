package tree

import (
	"fmt"
)

// Insert into BST
func (self *Tree) insertBST(parent *Node, newNode *Node) {
	if parent == nil {
		panic("This shouldn't happen")
	}

	compareVal := (*self.comparator)(parent.data, newNode.data)

	if compareVal == 0 {
		// There's no need to do anything
		fmt.Println("Already found that value. Doing nothing")
		return
	}

	dirn := "left"
	if compareVal == -1 {
		dirn = "right"
	}

	childNode, isExists := parent.link[dirn]
	if isExists {
		fmt.Println("Going ", dirn)
		self.insertBST(childNode, newNode)
	} else {
		fmt.Println("Inserting at ", dirn)
		// It needs to be inserted here
		parent.AddChild(dirn, newNode)
	}
}

// Searching a BST
func (self *Tree) hasValueBST(node *Node, key interface{}) bool {
	searchResp := self.getNodeBST(node, key)
	return searchResp != nil
}

func (self *Tree) getNodeBST(node *Node, key interface{}) *Node {
	if node == nil {
		return node
	}

	compareVal := (*self.comparator)(node.data, &key)
	if compareVal == 0 {
		return node
	}

	dirn := "left"
	if compareVal == -1 {
		dirn = "right"
	}

	return self.getNodeBST(node.link[dirn], key)
}
