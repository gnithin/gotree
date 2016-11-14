package tree

import (
	"github.com/nu7hatch/gouuid"
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
		expectedKey := strings.Trim(strings.ToLower(key), "")
		debug(expectedKey)

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
		BaseTree:          *CreateTreeWithComparator(&funcPtr),
		matchSubstring:    expectedMap["partial_match"],
		caseInsensitive:   expectedMap["case_insensitive"],
		stripPunctuations: expectedMap["strip_punctuations"],
		stripStopWords:    expectedMap["strip_stopwords"],
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
		BaseTree:          *CreateTreeWithComparator(&funcPtr),
		matchSubstring:    supportSubstring,
		caseInsensitive:   caseInsensitive,
		stripPunctuations: TRIE_DEFAULT_STRIP_PUNCTUATIONS,
		stripStopWords:    TRIE_DEFAULT_STRIP_STOP_WORDS,
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
