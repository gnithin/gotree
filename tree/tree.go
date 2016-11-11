package tree

import (
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"strconv"
	"strings"
)

// Public interface funtions
func CreateTrieWithOptionsMap(ipMap map[string]bool) *Trie {
	expectedMap := map[string]bool{
		"partial_match":      TRIE_DEFAULT_SUBSTRING_MATCH,
		"case_insensitive":   TRIE_DEFAULT_CASE_INSENSITIVE,
		"strip_stopwords":    TRIE_DEFAULT_STRIP_STOP_WORDS,
		"strip_punctuations": TRIE_DEFAULT_STRIP_PUNCTUATIONS,
		"word_separator":     TRIE_DEFAULT_WORD_SEPARATOR,
	}

	for key, value := range ipMap {
		_, keyExists := expectedMap[key]
		if keyExists {
			expectedMap[key] = value
		} else {
			debug("Incorrect key for creating a Trie", key)
		}
	}

	// By default adding a string comparator
	funcPtr := stringComparator
	trieObj := &Trie{
		*CreateTreeWithComparator(&funcPtr),
		expectedMap["partial_match"],
		expectedMap["case_insensitive"],
	}

	// Creating a base element. It's the default start
	var defaultVal interface{}
	defaultVal = TRIE_DEFAULT_VALUE
	baseElement := CreateNode(&defaultVal, map[string]*Node{})

	// Assigning the base element to the root
	trieObj.root = baseElement

	return trieObj
}
func CreateTrieWithOptions(supportSubstring, caseInsensitive bool) *Trie {
	// By default adding a string comparator
	funcPtr := stringComparator
	trieObj := &Trie{
		*CreateTreeWithComparator(&funcPtr),
		supportSubstring,
		caseInsensitive,
	}

	// Creating a base element. It's the default start
	var defaultVal interface{}
	defaultVal = TRIE_DEFAULT_VALUE
	baseElement := CreateNode(&defaultVal, map[string]*Node{})

	// Assigning the base element to the root
	trieObj.root = baseElement

	return trieObj
}

func CreateTrie() *Trie {
	return CreateTrieWithOptions(
		TRIE_DEFAULT_SUBSTRING_MATCH,
		TRIE_DEFAULT_CASE_INSENSITIVE,
	)
}

func CreateTree() *BaseTree {
	return CreateTreeWithComparator(nil)
}

func CreateTreeWithComparator(comparator *func(obj1, obj2 *interface{}) int) *BaseTree {
	uuid, err := uuid.NewV4()
	if err != nil {
		panic("Error generating a new UUID.")
	}

	nodesArr := []map[string]interface{}{}
	edgesArr := []map[string]interface{}{}

	tMap := map[string]interface{}{
		"nodes": nodesArr,
		"edges": edgesArr,
	}

	return &BaseTree{
		root:        nil,
		len:         0,
		leavesLen:   0,
		id:          uuid.String(),
		treeDispMap: tMap,
		comparator:  comparator,
	}
}

func CreateBST() *BST {
	return CreateBSTWithComparator(nil)
}

func CreateBSTWithComparator(comparator *func(obj1, obj2 *interface{}) int) *BST {
	return &BST{*(CreateTreeWithComparator(comparator))}
}

func CreateHeap() *Heap {
	return CreateMaxHeap()
}

func CreateMaxHeap() *Heap {
	return CreateMaxHeapWithSize(0)
}

func CreateMaxHeapWithSize(size int) *Heap {
	return CreateHeapWithComparator(nil, true, size)
}

func CreateMinHeap() *Heap {
	return CreateMinHeapWithSize(0)
}

func CreateMinHeapWithSize(size int) *Heap {
	return CreateHeapWithComparator(nil, false, size)
}

func CreateHeapWithComparator(comparator *func(obj1, obj2 *interface{}) int, isMaxHeap bool, heapSize int) *Heap {
	return MakeHeap(CreateTreeWithComparator(comparator), isMaxHeap, heapSize)
}

// Default integer comparator
func intComparator(obj1, obj2 *interface{}) int {
	new_obj1 := (*obj1).(int)
	new_obj2 := (*obj2).(int)
	if new_obj1 < new_obj2 {
		return -1
	} else if new_obj1 > new_obj2 {
		return 1
	} else {
		return 0
	}
}

// Default string comparator
func stringComparator(obj1, obj2 *interface{}) int {
	new_obj1 := (*obj1).(string)
	new_obj2 := (*obj2).(string)
	if new_obj1 < new_obj2 {
		return -1
	} else if new_obj1 > new_obj2 {
		return 1
	} else {
		return 0
	}
}

// Bucket of common functionality available across the different trees
type Tree interface {
	Insert(...interface{}) bool
	Pop() (*interface{}, bool)
}

type BaseTree struct {
	id          string
	root        *Node
	len         int
	leavesLen   int
	treeDispMap map[string]interface{}
	comparator  *func(obj1, obj2 *interface{}) int
}

func (self *BaseTree) IsEmpty() bool {
	return self.len == 0
}

func (self *BaseTree) checkTypeForComparator(node *Node) bool {
	// Just check if there's a comparator specified
	// Find the type. If the type is either a string or an int,
	// add the default comparator. Else raise error
	// Type assertion only provided for some things
	switch (*node.data).(type) {
	case int:
		temp := intComparator
		self.comparator = &temp
	case string:
		temp := stringComparator
		self.comparator = &temp
	default:
		if self.comparator == nil {
			//debug("City on Fire... City on Fire... Mischief!! Mischief!!")
			//panic("Need to specify comparator if the type is not string or int")
			return false
		}
	}
	return true
}

// Creates a JSON output for the current tree as specified by alchemy
func (self *BaseTree) GetJSONTree() []byte {
	self.postOrderTraverse(self.root)
	debug(self.treeDispMap)

	var treeJson []byte
	var jsonErr error

	if PRETTY_PRINT_TREE {
		treeJson, jsonErr = json.MarshalIndent(self.treeDispMap, "", "    ")
	} else {
		treeJson, jsonErr = json.Marshal(self.treeDispMap)
	}

	if jsonErr != nil {
		debug("Error marshalling the tree to json")
		debug(jsonErr)
	} else {
		return treeJson
	}

	return nil
}

/*
- Traverse each node
- Add it self
- Traverse in post order, when visiting every child, add the node.
- When visiting every root, add the edge
*/
func (self *BaseTree) postOrderTraverse(root *Node) (string, bool) {
	/*
		The resulting structure needs to be of this format -
		{
			"nodes": [
				{
					"id" : <int>,
					"caption" : "",
					"type" : ""
				},
				...
			],
			"edges" : [
				{
					"source" : <int>,
					"target" : <int>,
					"caption" :  ""
				},
				...
			]
		}
	*/
	if root == nil {
		return "", false
	}

	treeRepNodesArr := self.treeDispMap["nodes"].([]map[string]interface{})
	// Adding the root node
	treeRepNodesArr = append(
		treeRepNodesArr,
		map[string]interface{}{
			"id":      root.id,
			"caption": root.GetInfoString(),
			"type":    "",
		},
	)
	// This is a bit crazy, but hey, I didn't make the rules
	// http://stackoverflow.com/questions/28054913/modify-array-of-interface-golang
	//self.treeDispMap["nodes"] = treeRepNodesArr.(interface{})
	self.treeDispMap["nodes"] = treeRepNodesArr

	// Go left
	luuid, lExists := self.postOrderTraverse(root.link["left"])

	// Go right
	ruuid, rExists := self.postOrderTraverse(root.link["right"])

	// Adding the edges
	if lExists || rExists {
		treeRepEdgesArr := self.treeDispMap["edges"].([]map[string]interface{})
		if lExists {
			treeRepEdgesArr = append(
				treeRepEdgesArr,
				map[string]interface{}{
					"source":  root.id,
					"target":  luuid,
					"caption": "left",
				},
			)
		}

		if rExists {
			treeRepEdgesArr = append(
				treeRepEdgesArr,
				map[string]interface{}{
					"source":  root.id,
					"target":  ruuid,
					"caption": "right",
				},
			)
		}

		//self.treeDispMap["edges"] = treeRepEdgesArr.(interface{})
		self.treeDispMap["edges"] = treeRepEdgesArr
	}

	return root.id, true
}

type BaseSequentialTree struct {
	BaseTree
	nodeArr []*Node
	maxSize int
}

func (self *BaseSequentialTree) String() string {
	// Only print the non empty values
	// TODO: this can be made super fast
	var respList []string
	for i := 0; i < self.maxSize; i++ {
		if self.nodeArr[i] != nil && self.nodeArr[i].data != nil {
			respList = append(
				respList,
				strconv.Itoa(i)+":"+self.nodeArr[i].GetInfoString(),
			)
		}
	}
	resp := fmt.Sprintf(strings.Join(respList, "\n"))
	return resp
}

func (self *BaseSequentialTree) getParentIndex(childIndex int) int {
	if childIndex < 0 {
		panic("Tree index cannot be < 0")
	}

	return (childIndex - 1) / 2
}

func (self *BaseSequentialTree) getChildIndex(parentIndex int, isLeft bool) int {
	if parentIndex < 0 {
		panic("Tree index cannot be < 0")
	}

	inc := 1
	if !isLeft {
		inc = 2
	}
	return (2 * parentIndex) + inc
}

func (self *BaseSequentialTree) isLeftChild(childIndex int) bool {
	if childIndex <= 0 {
		panic("Tree index cannot be <= 0")
	}

	return (childIndex % 2) != 0
}

func (self *BaseSequentialTree) IsFull() bool {
	return self.len >= self.maxSize
}
