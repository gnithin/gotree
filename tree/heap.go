package tree

type locationNode struct {
	nodePtr *Node
	dirn    string
	nextPtr *locationNode
}

type nextLocationQueue struct {
	front      *locationNode
	rear       *locationNode
	lastPopped *locationNode
}

func (q *nextLocationQueue) isEmpty() bool {
	return q.front == nil && q.front == q.rear
}

func (q *nextLocationQueue) insertFront(locPtr *locationNode) {
	if q.front != nil {
		fNode := q.front
		fNode.nextPtr = locPtr
	}
}

func (q *nextLocationQueue) push(nodePtr *Node, dirn string) {
	newNode := &(locationNode{
		nodePtr: nodePtr,
		dirn:    dirn,
		nextPtr: nil,
	})

	// Add to the rear of the queue
	if q.rear == nil {
		q.rear = newNode

		// It stands to reason that even q.front will be nil
		q.front = newNode
	} else {
		oldElem := q.rear
		oldElem.nextPtr = newNode
		q.rear = newNode
	}
}

func (q *nextLocationQueue) pop() *locationNode {
	if q.front == nil {
		return nil
	}

	nodeVal := q.front
	q.front = nodeVal.nextPtr
	return nodeVal
}

// HEAP
type Heap struct {
	BaseTree
	locQueue  *nextLocationQueue
	isMaxHeap bool
}

func makeHeap(b *BaseTree, isMaxHeap bool) *Heap {
	locQueue := &nextLocationQueue{
		front: nil,
		rear:  nil,
	}

	return &Heap{
		*b,
		locQueue,
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

func (self *Heap) isSizer(obj1, obj2 *interface{}) {
	compareResp := self.comparator(obj1, obj2)
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
