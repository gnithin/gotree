package tree

import (
//"fmt"
)

type huffmanData struct {
	data string
	freq int
}

type HuffmanTree struct {
	BaseTree
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

	return &HuffmanTree{
		*CreateTree(),
		CreateHeapWithComparator(&comparatorFunc, true, 500),
		freqMap,
	}
}

func (self *HuffmanTree) buildTree() bool {
	if len(self.freqMap) == 0 {
		return false
	}

	respStatus := true
	// Put everything inside the priority queue
	for keyStr, freq := range self.freqMap {
		respStatus = respStatus &&
			self.priorityQueue.Insert(huffmanData{
				keyStr,
				freq,
			})
	}

	// TODO:
	// Pop 2 elements at a time.
	// Create a new node - Need to change huffmanData
	// Then add it back

	return true
}

func (self *HuffmanTree) EncodeStr(ipStr string) string {
	return "Not yet implemented"
}
