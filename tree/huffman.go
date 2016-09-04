package tree

import (
	"fmt"
)

type huffmanData struct {
	dataVal string
	freq    int
	leaf    bool
	link    map[string]*huffmanData
}

type HuffmanTree struct {
	root          *huffmanData
	priorityQueue *Heap
	freqMap       map[string]int
}

func CreateHuffmanTree(freqMap map[string]int) *HuffmanTree {
	comparatorFunc := func(obj1, obj2 *interface{}) int {
		new_obj1 := (*obj1).(huffmanData)
		new_obj2 := (*obj2).(huffmanData)

		// Never send 0 - Doesn't have any use
		if new_obj1.freq >= new_obj2.freq {
			return 1
		} else {
			return -1
		}
	}

	huffmanTree := &HuffmanTree{
		nil,
		CreateHeapWithComparator(&comparatorFunc, true, 500),
		freqMap,
	}

	huffmanTree.buildTree()

	return huffmanTree
}

func (self *HuffmanTree) buildTree() bool {
	if len(self.freqMap) == 0 {
		return false
	}

	fmt.Println("In buildtree")

	respStatus := true
	// Put everything inside the priority queue
	var newData interface{}
	for keyStr, freq := range self.freqMap {
		newData = huffmanData{
			keyStr,
			freq,
			true,
			map[string]*huffmanData{
				"0": nil,
				"1": nil,
			},
		}
		respStatus = respStatus &&
			self.priorityQueue.Insert(newData)
	}

	// Pop 2 elements at a time.
	for !self.priorityQueue.IsEmpty() {
		leftChildInt, isValidLChild := self.priorityQueue.Pop()
		rightChildInt, isValidRChild := self.priorityQueue.Pop()

		if isValidLChild {
			leftChild := (*leftChildInt).(huffmanData)
			if isValidRChild {
				rightChild := (*rightChildInt).(huffmanData)

				// Add the data from both the nodes
				newData := huffmanData{
					"",
					leftChild.freq + rightChild.freq,
					false,
					map[string]*huffmanData{
						"0": &rightChild,
						"1": &leftChild,
					},
				}
				respStatus = respStatus &&
					self.priorityQueue.Insert(newData)
			} else {
				// Only one child remains. Add it to the tree
				self.root = &leftChild
			}
		}
	}
	return true
}

func (self *HuffmanTree) EncodeStr(ipStr string) string {
	return "Not yet implemented"
}
