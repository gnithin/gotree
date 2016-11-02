package tree

import (
	"strings"
)

// This is basically all that's required.
type Trie struct {
	BaseTree
	matchSubstring  bool
	caseInsensitive bool
}

func (self *Trie) Insert(valSlice ...interface{}) bool {
	if len(valSlice) <= 0 {
		debug("Nothing to insert in a trie")
		return false
	}

	finalResp := true
	for _, val := range valSlice {
		intermediate_resp := self.InsertOne(val)
		finalResp = finalResp && intermediate_resp
	}

	return true
}

func (self *Trie) InsertOne(ipObj interface{}) bool {
	debug(ipObj)
	ipStr := ipObj.(string)
	ipStr = strings.Trim(ipStr, "")

	if len(ipStr) == 0 {
		debug("Trying to insert empty string")
		return false
	}

	debug("Inserting ->", ipStr, "<-")

	currentNode := self.root
	finalIndex := len(ipStr) - 1

	// TODO: When creating a radix tree, this loop will have
	// to change.
	// Check if the existing nodes list has the value.
	for currIndex, char := range ipStr {
		mapVal, isPresent := currentNode.link[string(char)]

		if isPresent {
			currentNode = mapVal
		} else {
			// Making the new node
			newNode := self.createTrieNode(byte(char))
			currentNode.link[string(char)] = newNode
			currentNode = newNode
		}

		// Set the replenished current node value to support the final value
		if self.matchSubstring && currIndex == finalIndex {
			currentNode.link[TRIE_FINAL_NODE_KEY] = nil
		}
	}
	return true
}

func (self *Trie) createTrieNode(charEntry byte) *Node {
	var nodeVal interface{}
	nodeVal = charEntry
	return CreateNode(&nodeVal, map[string]*Node{})
}

func (self *Trie) HasVal(needle string) bool {
	currentNode := self.root
	if currentNode == nil {
		debug("Searching with nil current node in Trie")
		return false
	}

	return false
}
