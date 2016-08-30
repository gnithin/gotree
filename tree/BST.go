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
func (self *Tree) RemoveValBST(key interface{}) bool {
	nodeResp := self.getNodeBST(self.root, key)
	if nodeResp == nil {
		fmt.Println("Cannot find the required key to remove")
		return false
	}

	parentNode, IsParentKey := nodeResp.link["parent"]
	if !IsParentKey {
		panic("The parent key in the nodes must be enabled for the removal to work!!! The alternative has not been implemented!")
	}

	// NOTE: ParentDirn should only be used whenever there is a parentNode available
	// Enforce that
	parentDirn := ""
	if parentNode != nil {
		if parentNode.link["left"] == nodeResp {
			// it's the left kid
			parentDirn = "left"
		} else {
			parentDirn = "right"
		}
	}

	leftNode, IsLeftKey := nodeResp.link["left"]
	rightNode, IsRightKey := nodeResp.link["right"]

	if !IsLeftKey || !IsRightKey {
		panic("The left and right key in the nodes must be enabled for the removal to work!!!")
	}

	// The actual logic of the removal starts here
	if leftNode == nil && rightNode == nil {
		// If it's the root, then just nuke it
		if parentNode == nil {
			self.root = nil
		} else {
			parentNode.link[parentDirn] = nil
		}
	} else if leftNode == nil || rightNode == nil {
		nonEmptyNode := leftNode
		if leftNode == nil {
			nonEmptyNode = rightNode
		}

		// Just make the pointer of the non-empty side point to the parent
		if parentNode == nil {
			self.root = nonEmptyNode
			nonEmptyNode.link["parent"] = nil
		} else {
			parentNode.link[parentDirn] = nonEmptyNode
			nonEmptyNode.link["parent"] = parentNode
		}
	} else {
		// Both the children exist
		lrNode, lrExists := leftNode.link["right"]

		if !lrExists {
			panic("The left and right key in the nodes must be enabled for the removal to work!!!")
		}

		// Attach lr to the left of the leftmost child of right node of r
		if lrNode != nil {
			// Getting the leftMostRightNode
			leftMostRightNode := rightNode
			for leftMostRightNode.link["left"] != nil {
				leftMostRightNode = leftMostRightNode.link["left"]
			}

			leftMostRightNode.link["left"] = lrNode
			lrNode.link["parent"] = leftMostRightNode
		}

		leftNode.link["right"] = rightNode
		rightNode.link["parent"] = leftNode

		if parentNode == nil {
			self.root = leftNode
			leftNode.link["parent"] = nil
		} else {
			leftNode.link["parent"] = parentNode
			parentNode.link[parentDirn] = leftNode
		}
	}

	// Removing all references from the node to be deleted
	// This is just for safety purposes
	nodeResp.link["left"] = nil
	nodeResp.link["right"] = nil
	nodeResp.link["parent"] = nil

	return true
}
