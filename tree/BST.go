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
	if node == nil {
		return false
	}

	compareVal := (*self.comparator)(node.data, &key)
	if compareVal == 0 {
		return true
	}

	dirn := "left"
	if compareVal == -1 {
		dirn = "right"
	}

	return self.hasValueBST(node.link[dirn], key)
}
