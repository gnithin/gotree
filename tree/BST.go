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
	// isExists needn't be checked. It can be safely removed
	if isExists && childNode != nil {
		fmt.Println("Going ", dirn)
		self.insertBST(childNode, newNode)
	} else {
		fmt.Println("Inserting at ", dirn)
		// It needs to be inserted here
		parent.AddChild(dirn, newNode)

		// Adding a parent
		_, isExists := newNode.link["parent"]
		if !isExists {
			newNode.link["parent"] = parent
		}
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

// Removing an element from a BST
func (self *Tree) removeBST(node *Node, key interface{}) bool {
	nodeResp := self.getNodeBST(self.root, key)
	if nodeResp == nil {
		return false
	}

	// This will only work if there's a parent key
	parentNode, IsParentKey := nodeResp.link["parent"]
	if !IsParentKey {
		panic("The parent key in the nodes must be enabled for the removal to work!!! The alternative has not been implemented!")
	}
	var parentDirn string
	parentDirn = ""
	if parentNode != nil {
		if parentNode.link["left"] == node {
			// it's the left kid
			parentDirn = "left"
		} else {
			parentDirn = "right"
		}
	}

	leftNode, IsLeftKey := nodeResp.link["left"]
	rightNode, IsRightKey := nodeResp.link["right"]

	if !IsLeftKey || !IsRightKey {
		// TODO: Better error message
		panic("This shouldn't be happening!!!! Arrrrgh!!!!")
	}

	if leftNode == nil && rightNode == nil {
		// If it's the root, then just nuke it
		if parentNode == nil {
			self.root = nil
		} else {
			// Remove this thing from the parent.
			// Find out what child of the parent it is
			parentNode.link[parentDirn] = nil
		}
		// TODO: DEstroy node
	} else if leftNode == nil || rightNode == nil {
		nonEmptyNode := leftNode
		if leftNode == nil {
			nonEmptyNode = rightNode
		}
		// Either one are nil
		// Just make the pointer of the non-empty side point to the parent
		if parentNode == nil {
			// It's the root. Handle it differently
			self.root = nonEmptyNode
			nonEmptyNode.link["parent"] = nil
		} else {
			parentNode.link[parentDirn] = nonEmptyNode
			nonEmptyNode.link["parent"] = parentNode
		}
	} else {
		// Both are not nil.
		// TODO: this is a bit tricky
	}
	return false
}
