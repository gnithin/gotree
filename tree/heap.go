package tree

import (
	"fmt"
)

const (
	MAX_SIZE           = 500
	DEFAULT_ROOT_INDEX = 0
)

// Public interface function
func MakeHeap(b *BaseTree, isMaxHeap bool, heapSize int) *Heap {
	if heapSize <= 0 {
		heapSize = MAX_SIZE
	}

	return &Heap{
		BaseSequentialTree{
			*b,
			make([]*Node, heapSize),
			heapSize,
		},
		DEFAULT_ROOT_INDEX,
		isMaxHeap,
	}
}

// HEAP
type Heap struct {
	BaseSequentialTree
	nextInsertIndex int
	isMaxHeap       bool
}

func (self *Heap) GetHeapLen() int {
	return self.len
}

func (self *Heap) String() string {
	heapType := "MinHeap"
	if self.isMaxHeap {
		heapType = "MaxHeap"
	}
	return fmt.Sprintf("Heap Size: %d\nHeap len: %d\nHeap Type: %s\nTree:\n%v\n",
		self.maxSize,
		self.len,
		heapType,
		&self.BaseSequentialTree)
}

// Do not mess arround with the link pointers
// Just change the data section inside the pointers
// This is so that the links to the pointers are not modified
func (self *Heap) swapData(node1, node2 *Node) {
	container := node1.data
	node1.data = node2.data
	node2.data = container
}

// This exchanges the pointer values
func (self *Heap) swapPtr(node1, node2 *Node) {
	container := node1
	node1 = node2
	node2 = container
}

// Used as a wrapper for changing the comparator logic
// depending on the type of heap used.
// Returns true, if after comparing obj1 and obj2, it turns out
// that obj1 is appropriately placed in the present type of heap
// compared to obj2.
// Naming logic ;{p
// isSizer -> (For maxheap) isBigger, (For minHeap) isSmaller
// Sorry, couldn't think of a better name
func (self *Heap) isSizer(obj1, obj2 *interface{}) bool {
	if self.comparator != nil {
		compareResp := (*self.comparator)(obj1, obj2)
		if self.isMaxHeap {
			// obj1 >= obj2
			if compareResp >= 0 {
				return true
			}
		} else {
			if compareResp <= 0 {
				return true
			}
		}
	}
	return false
}

func (self *Heap) reheapUpSeq(node *Node, currIndex int) {
	if self.len <= 0 {
		return
	}

	// Just trying out to see if the overhead of computing the parent
	// arithmatically is faster than directyl linking the parent
	//parentNode := node.link["parent"]
	parentNodeInd := self.getParentIndex(currIndex)
	if parentNodeInd >= 0 {
		parentNode := self.nodeArr[parentNodeInd]

		for parentNode != nil && !self.isSizer(parentNode.data, node.data) {
			// TODO: Don't know how bad this is.
			// On first glance, looks pretty bad.

			// Reason : It's actually pretty simple to visualize the data being swapped.
			// The links remain the same. Probably will have to check how badly
			// the performance degrades if the maps are being swapped.
			// Testing this on a life-sized structure should be more appropriate.

			// Dude, data swapping is basically the data pointers being swapped.
			// No harm in it. Asshole
			self.swapData(parentNode, node)

			node = parentNode
			parentNode = nil
			if parentNodeInd > 0 {
				parentNodeInd = self.getParentIndex(parentNodeInd)
				if parentNodeInd >= 0 {
					parentNode = self.nodeArr[parentNodeInd]
				}
			}
		}
	}
}

func (self *Heap) reheapUp(node *Node) {
	if self.len <= 0 {
		return
	}

	parentNode := node.link["parent"]

	for parentNode != nil && !self.isSizer(parentNode.data, node.data) {
		self.swapData(parentNode, node)
		node = parentNode
		parentNode = node.link["parent"]
	}
}

func (self *Heap) reheapDown() {
	if self.len <= 0 {
		return
	}

	parentNode := self.root
	needToCompareFlag := true

	for needToCompareFlag {
		rightChild := parentNode.link["right"]
		leftChild := parentNode.link["left"]

		needToCompareFlag = false

		if rightChild != nil || leftChild != nil {
			heavyChild := leftChild

			if rightChild != nil && leftChild != nil {
				if self.isSizer(rightChild.data, leftChild.data) {
					heavyChild = rightChild
				}
			} else if rightChild != nil {
				heavyChild = rightChild
			}

			// Compare parent with the heavy child
			if !self.isSizer(parentNode.data, heavyChild.data) {
				self.swapData(parentNode, heavyChild)
				parentNode = heavyChild

				// Need to repeat the loop because swapping happened
				needToCompareFlag = true
			}
		}
	}
}

/*
	Exact same code is available in BST as well.
	Think about making it common somehow.
	(Macros come to mind, by golang does not have it)
*/
func (self *Heap) Insert(valSlice ...interface{}) bool {
	if len(valSlice) == 0 {
		return false
	}

	insertResp := true
	for _, val := range valSlice {
		insertResp = insertResp && self.InsertOne(val)
	}
	return insertResp
}

func (self *Heap) InsertOne(newVal interface{}) bool {
	if self.IsFull() {
		debug("Heap size limit reached")
		return false
	}

	newNode := CreateTreeNode(&newVal)
	if self.root == nil {
		isValid := self.checkTypeForComparator(newNode)
		if !isValid {
			return false
		}
		self.root = newNode
	}

	// Inserting into the node arr
	self.nodeArr[self.nextInsertIndex] = newNode
	if self.nextInsertIndex != 0 {
		parentNode := self.nodeArr[self.getParentIndex(self.nextInsertIndex)]
		parentDirn := "right"
		if self.isLeftChild(self.nextInsertIndex) {
			parentDirn = "left"
		}

		// Updating the link properties
		parentNode.link[parentDirn] = newNode
		newNode.link["parent"] = parentNode

		//self.reheapUp(newNode)
		self.reheapUpSeq(newNode, self.nextInsertIndex)
	}

	self.nextInsertIndex += 1
	self.len += 1
	debug("Inserting", newVal)
	return true
}

func (self *Heap) Pop() (*interface{}, bool) {
	if self.len == 0 {
		return nil, false
	}

	lastElementIndex := self.nextInsertIndex - 1
	// Copy data for returning purposes
	respData := *self.root.data

	// Remove the last element
	lastElement := self.nodeArr[lastElementIndex]

	self.swapData(lastElement, self.nodeArr[DEFAULT_ROOT_INDEX])

	self.nodeArr[lastElementIndex] = nil

	// Clean up the node references as well.
	lastElement.link["left"] = nil
	lastElement.link["right"] = nil

	lastElementParent, isParentExists := lastElement.link["parent"]
	if isParentExists && lastElementParent != nil {
		parDirn := "right"
		if self.isLeftChild(lastElementIndex) {
			parDirn = "left"
		}
		lastElementParent.link[parDirn] = nil
	}

	self.len -= 1
	self.nextInsertIndex -= 1

	if self.len == 0 {
		self.root = nil
	}
	// Reheap down
	self.reheapDown()

	debug("Popping - ", respData)
	return &respData, true
}

// Keeping this for the interface purpose.
// TODO: Think about how to navigate through this problem. Or just let it be
func (self *Heap) HasVal(*Node, interface{}) bool {
	panic("Heap does not understand the HasVal method. Get the root element")
}

func (self *Heap) Remove(...interface{}) bool {
	panic("Heap does not understand the Remove method. Use pop()")
}
