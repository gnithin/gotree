package tree

import (
//"fmt"
)

type BST struct {
	BaseTree
}

/*
	Exact same code is available in BST as well.
	Think about making it common somehow.
	(Macros come to mind, by golang does not have it)
*/
func (self *BST) Insert(valSlice ...interface{}) bool {
	if len(valSlice) <= 0 {
		// fmt.Println("No value provided for insertion")
		return false
	}

	insertStatus := true
	for _, val := range valSlice {
		insertStatus = insertStatus && self.InsertOne(val)
	}
	return insertStatus
}

func (self *BST) InsertOne(newVal interface{}) bool {
	newNode := CreateTreeNode(&newVal)
	var insertStatus bool
	if self.root == nil {
		isValid := self.checkTypeForComparator(newNode)
		if !isValid {
			return false
		}
		self.root = newNode
		insertStatus = true
	} else {
		insertStatus = self.insertBST(self.root, newNode)
	}
	self.len += 1

	return insertStatus
}

// Insert into BST
func (self *BST) insertBST(parent *Node, newNode *Node) bool {
	if parent == nil {
		panic("This shouldn't happen")
	}

	compareVal := (*self.comparator)(parent.data, newNode.data)

	if compareVal == 0 {
		// There's no need to do anything
		//fmt.Println("Already found that value. Doing nothing")
		return false
	}

	dirn := "left"
	if compareVal == -1 {
		dirn = "right"
	}

	childNode, isExists := parent.link[dirn]
	// isExists needn't be checked. It can be safely removed
	if isExists && childNode != nil {
		//fmt.Println("Going ", dirn)
		return self.insertBST(childNode, newNode)
	} else {
		//fmt.Println("Inserting at ", dirn)
		// It needs to be inserted here
		parent.link[dirn] = newNode

		// Adding a parent
		newNode.link["parent"] = parent
		return true
	}
}

// Searching a BST
func (self *BST) HasVal(key interface{}) bool {
	searchResp := self.getNodeBST(self.root, key)
	return searchResp != nil
}

func (self *BST) getNodeBST(node *Node, key interface{}) *Node {
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

func (self *BST) Pop() (*interface{}, bool) {
	// Popping the root element
	if self.root != nil {
		rootData := self.root.data
		if self.Remove(*rootData) {
			return rootData, true
		}
	}
	return nil, false
}

func (self *BST) Remove(keySlice ...interface{}) bool {
	respStatus := true
	for _, key := range keySlice {
		respStatus = respStatus && self.RemoveOne(key)
	}
	return respStatus
}

// Removing an element from a BST
func (self *BST) RemoveOne(key interface{}) bool {
	nodeResp := self.getNodeBST(self.root, key)
	if nodeResp == nil {
		//fmt.Println("Cannot find the required key to remove")
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
