package tree

import (
	"github.com/nu7hatch/gouuid"
)

type huffmanData struct {
	id      string
	dataVal string
	freq    int
	leaf    bool
	link    map[string]*huffmanData
}

type HuffmanTree struct {
	root          *huffmanData
	priorityQueue *Heap
	freqMap       map[string]int
	encodingMap   map[string]string
}

func CreateHuffmanTree(freqMap map[string]int) *HuffmanTree {
	// Huffmans' custom comparator for the heap
	// The heap uses the main comparision value as frequency
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
		CreateHeapWithComparator(&comparatorFunc, false, 500),
		freqMap,
		make(map[string]string),
	}

	if !huffmanTree.buildTree() {
		return nil
	}

	return huffmanTree
}

func (self *HuffmanTree) buildTree() bool {
	if len(self.freqMap) == 0 {
		return false
	}

	leavesMapPtr := make(map[string]*huffmanData)

	respStatus := true
	// Put everything inside the priority queue
	var interfaceData interface{}
	for keyStr, freq := range self.freqMap {
		uuid, _ := uuid.NewV4()
		newData := huffmanData{
			uuid.String(),
			keyStr,
			freq,
			true,
			map[string]*huffmanData{
				"0":      nil,
				"1":      nil,
				"parent": nil,
			},
		}
		interfaceData = newData

		currInsStatus := self.priorityQueue.Insert(interfaceData)
		respStatus = respStatus && currInsStatus

		leavesMapPtr[keyStr] = &newData
	}

	debug("******************")
	debug("Leaves!")
	debug(leavesMapPtr)
	debug(self.priorityQueue.GetHeapLen())
	debug("******************")

	// Pop 2 elements at a time.
	for !self.priorityQueue.IsEmpty() {

		debug("Heap len", self.priorityQueue.GetHeapLen())
		leftChildInt, isValidLChild := self.priorityQueue.Pop()
		rightChildInt, isValidRChild := self.priorityQueue.Pop()

		if isValidLChild {
			leftChild := (*leftChildInt).(huffmanData)
			debug("Heap State - left - ", leftChild)
			if isValidRChild {
				rightChild := (*rightChildInt).(huffmanData)
				debug("Heap State - right - ", rightChild)
				uuid, _ := uuid.NewV4()

				// Add the data from both the nodes
				newData := huffmanData{
					uuid.String(),
					"",
					leftChild.freq + rightChild.freq,
					false,
					map[string]*huffmanData{
						"1":      &rightChild,
						"0":      &leftChild,
						"parent": nil,
					},
				}
				rightChild.link["parent"] = &newData
				leftChild.link["parent"] = &newData

				// Insert it back into the priority queue
				debug("Inserting - ", newData)
				insState := self.priorityQueue.Insert(newData)
				respStatus = respStatus && insState
			} else {
				// Only one child remains. Add it to the tree
				self.root = &leftChild
			}
		} else {
			debug("Is INVALID LEFTIE")
		}
	}

	// Fill the encoding map
	// Traverse upwards till the root for every leaf
	for keyStr, valPtr := range leavesMapPtr {
		debug(keyStr)

		leafNode := valPtr
		codedStr := ""
		for leafNode != nil {
			debug(leafNode.freq)
			debug(leafNode.link)

			//Find out what child the current node is
			parentNode := leafNode.link["parent"]
			if parentNode != nil {
				selectedCode := "1"
				if parentNode.link["0"].id == leafNode.id {
					selectedCode = "0"
				}
				debug("selected Code - ", selectedCode)

				codedStr = selectedCode + codedStr

				leafNode = parentNode
			} else {
				break
			}
		}
		self.encodingMap[keyStr] = codedStr
	}
	return true
}

func (self *HuffmanTree) EncodeStr(ipStr string) string {
	debug("Encoding map - ", self.encodingMap)
	encodedStr := ""
	for _, r := range ipStr {
		c := string(r)
		encodedStr += self.encodingMap[c]
	}
	return encodedStr
}

func (self *HuffmanTree) DecodeStr(ipStr string) string {
	debug("****************")
	debug("String to decode")
	debug(ipStr)

	curr_elem := self.root
	debug("Curr element - ", curr_elem)
	op_str := ""

	for _, r := range ipStr {
		c := string(r)
		curr_elem = curr_elem.link[c]

		if curr_elem == nil {
			debug("This should not happen")
			panic("Reached an invalid state")
		}

		if curr_elem.leaf {
			op_str += string(curr_elem.dataVal)
			curr_elem = self.root
		}
	}

	return op_str
}

// Creates a freq map for given input
func CreateFreqMap(ipStr string) map[string]int {
	// TODO: Add a exclude chars support
	freqMap := make(map[string]int)

	for _, ipRune := range ipStr {
		ipChar := string(ipRune)

		old_value, has_value := freqMap[ipChar]
		if has_value {
			freqMap[ipChar] = old_value + 1
		} else {
			freqMap[ipChar] = 0
		}
	}
	return freqMap
}
