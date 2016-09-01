package tree

// HEAP
type Heap struct {
	BaseTree
	nextInsertIndex int
	isMaxHeap       bool
}

func makeHeap(b *BaseSequentialTree, isMaxHeap bool) *Heap {
	return &Heap{
		*b,
		0,
		isMaxHeap,
	}
}

func (self *Heap) Insert(newVal interface{}) {
	newNode := CreateTreeNode(&newVal)
	if self.root == nil {
		self.checkTypeForComparator(newNode)
		self.root = newNode
	} else {
		// Pop from the nextLocation
		if self.locQueue.isEmpty() {
			panic("The queue is empty. Something is wrong!")
		}

		positionVal := self.locQueue.pop()
		parentNode := positionVal.nodePtr
		parentDirn := positionVal.dirn

		parentNode.link[parentDirn] = newNode
		newNode.link["parent"] = parentNode

		// Reheap up from here
		self.reheapUp(newNode)
	}
	// Insert newNode into next Location
	self.locQueue.push(newNode, "left")
	self.locQueue.push(newNode, "right")
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
	tempData := node1.data
	node1.data = node2.data
	node2.data = tempData
}

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
	panic("Not implemented yet!")
	// Bubble down
}

// Keeping this for the interface purpose.
// TODO: Think about how to navigate through this problem. Or just let it be
func (self *Heap) HasVal(*Node, interface{}) bool {
	panic("Heap does not understand the HasVal method. Get the root element")
}

func (self *Heap) Remove(interface{}) bool {
	panic("Heap does not understand the Remove method. Use pop()")
}
