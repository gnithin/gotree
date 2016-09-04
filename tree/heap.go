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
	if self.nextInsertIndex >= self.maxSize {
		//fmt.Println("Heap size limit reached")
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

	//fmt.Println("Index - ", self.nextInsertIndex)
	//fmt.Println("Value - ", newVal)

	// Inserting into the node arr
	self.nodeArr[self.nextInsertIndex] = newNode
	if self.nextInsertIndex != 0 {
		parentNode := self.nodeArr[self.getParentIndex(self.nextInsertIndex)]
		parentDirn := "right"
		if self.isLeftChild(self.nextInsertIndex) {
			parentDirn = "left"
		}
		parentNode.link[parentDirn] = newNode
		newNode.link["parent"] = parentNode

		// Reheap up from here
		self.reheapUp(newNode)

		//fmt.Println("Dirn - ", parentDirn)
	}
	//fmt.Println("*********************")
	self.len += 1
	self.nextInsertIndex += 1
	return true
}

func (self *Heap) reheapUp(node *Node) {
	// Compare the present node with it's parent
	parentNode := node.link["parent"]
	for parentNode != nil {
		if !self.isSizer(parentNode.data, node.data) {
			self.swapData(parentNode, node)
			node = parentNode
			parentNode = node.link["parent"]
		} else {
			break
		}
	}
}

func (self *Heap) swapData(node1, node2 *Node) {
	// Do not mess arround with the link pointers
	// Just change the data section
	container := node1.data
	node1.data = node2.data
	node2.data = container
}

func (self *Heap) swapPtr(node1, node2 *Node) {
	container := node1
	node1 = node2
	node2 = container
}

// Used as a wrapper for the comparator for channelling in different heaps
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

func (self *Heap) Pop() (*interface{}, bool) {
	if self.len == 0 {
		return nil, false
	}
	lastElementIndex := self.nextInsertIndex - 1
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

	// Reheap down
	self.reheapDown()

	return &respData, true
}

func (self *Heap) reheapDown() {
	if self.len <= 1 {
		return
	}

	parentNode := self.root
	needToCompareFlag := true

	for needToCompareFlag {
		rightChild := parentNode.link["right"]
		leftChild := parentNode.link["left"]

		if rightChild != nil || leftChild != nil {
			heavyChild := leftChild
			if rightChild != nil && leftChild != nil {
				//Compare left/right child (Gasp!)
				if self.isSizer(rightChild.data, leftChild.data) {
					heavyChild = rightChild
				}
			} else if rightChild != nil && leftChild == nil {
				heavyChild = rightChild
			}

			// Compare parent with child
			if !self.isSizer(parentNode.data, heavyChild.data) {
				self.swapData(parentNode, heavyChild)
				parentNode = heavyChild
			} else {
				needToCompareFlag = false
			}
		} else {
			needToCompareFlag = false
		}
	}
}

// Keeping this for the interface purpose.
// TODO: Think about how to navigate through this problem. Or just let it be
func (self *Heap) HasVal(*Node, interface{}) bool {
	panic("Heap does not understand the HasVal method. Get the root element")
}

func (self *Heap) Remove(...interface{}) bool {
	panic("Heap does not understand the Remove method. Use pop()")
}
