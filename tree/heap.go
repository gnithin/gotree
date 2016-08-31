package tree

// This can be a simple array
type nextLocation struct {
	nodePtr *Node
	dirn    string
}

type Heap struct {
	BaseTree
	nextLoc []*nextLocation
}

func makeHeap(b *BaseTree) *Heap {
	var nextLoc []*nextLocation
	return &Heap{
		*b,
		nextLoc,
	}
}

func Insert(interface{}) {
	panic("Not implemented yet!")
	// Bubble up
}

func Pop() (*interface{}, bool) {
	panic("Not implemented yet!")
	// Bubble down
}

// Keeping this for the interface purpose.
// TODO: Think about how to navigate through this problem. Or just let it be
func HasVal(*Node, interface{}) bool {
	panic("Heap does not understand the HasVal method. Get the root element")
}
func Remove(interface{}) bool {
	panic("Heap does not understand the Remove method. Use pop()")
}
