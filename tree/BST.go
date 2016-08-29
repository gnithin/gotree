package tree

import (
	"fmt"
)

// Insert into BST
func (self *Tree) insertBST(parent *Node, newNode *Node) {
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
		self.insertBST(childNode, newNode)
	} else {
		// It needs to be inserted here
		parent.AddChild(dirn, newNode)
	}
}

// Searching a BST
func (self *Tree) hasValueBST(node *Node, key int) bool {
	if node == nil {
		return false
	}

	if node.data.Num == key {
		return true
	}

	dirn := "left"
	if node.data.Num < key {
		dirn = "right"
	}

	return self.hasValueBST(node.link[dirn], key)
}
