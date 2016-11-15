package tree

import (
	"strings"
	"unicode"
)

// This is basically all that's required.
type Trie struct {
	BaseTree
	matchSubstring    bool
	caseInsensitive   bool
	stripPunctuations bool
	stripStopWords    bool
}

func (self *Trie) InsertStr(ipStr string) bool {
	// I know that strings.Fields does precisely this.
	// But letting it open for allowing custom seperators
	// in the future

	//ipList = strings.FieldsFunc(
	//ipStr,
	//func(c rune) bool {
	//return unicode.IsSpace(c)
	//},
	//)

	finalResp := false
	ipList := strings.Fields(ipStr)
	if len(ipList) > 0 {
		/*
			Why am I doing this?
			That's because golang does not allow me to
			send []string... to an interface{}...
			Do NOT use cutesy ideas of making another function
			here to prevent DRY. That will also need the parameter
			type.
		*/
		finalResp = true
		for _, val := range ipList {
			intermediate_resp := self.InsertOne(val)
			finalResp = finalResp && intermediate_resp
		}
	}
	return finalResp
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

	return finalResp
}

func (self *Trie) InsertOne(ipObj interface{}) bool {
	ipStr := ipObj.(string)
	ipStr = strings.Trim(ipStr, "")

	if self.stripPunctuations {
		ipStr = strings.Map(
			func(r rune) rune {
				if unicode.IsPunct(r) {
					return -1
				}
				return r
			},
			ipStr,
		)
	}

	if self.caseInsensitive {
		ipStr = strings.ToLower(ipStr)
	}

	if len(ipStr) == 0 {
		debug("Trying to insert empty string")
		return false
	}

	if self.stripStopWords {
		// Create the stop word's Trie - This is awesomely meta
		stopWordsTrie := getTrieForStopWords()
		if stopWordsTrie.HasVal(ipStr) {
			debug("(stop word) Skipping", ipStr)

			// returning true, otherwise the whole thing evaluates to false
			return true
		}
	}

	debug("Inserting ->", ipStr, "<-")

	currentNode := self.root
	finalIndex := len(ipStr) - 1

	// TODO: When creating a radix tree, this loop will have to change.
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
		if currIndex == finalIndex {
			currentNode.link[TRIE_FINAL_NODE_KEY] = nil
		}
	}

	self.len += 1
	return true
}

func (self *Trie) createTrieNode(charEntry byte) *Node {
	var nodeVal interface{}
	nodeVal = charEntry
	return CreateNode(&nodeVal, map[string]*Node{})
}

func (self *Trie) HasVal(needle string) bool {
	needle = strings.Trim(needle, "")
	if self.caseInsensitive {
		needle = strings.ToLower(needle)
	}

	if len(needle) <= 0 {
		debug("Sent an empty string to search")
		return false
	}

	//debug("Searching -", needle)

	currentNode := self.root
	if currentNode == nil {
		debug("Searching with nil current node in Trie")
		return false
	}

	for _, char := range needle {
		nextNodeVal, isPresent := currentNode.link[string(char)]
		if !isPresent {
			return false
		}
		currentNode = nextNodeVal
	}

	if self.matchSubstring {
		return true
	} else {
		// Needs to match the whole thing, then the last node must
		// have a final_node_key in it's link.
		_, isFinalKeyPresent := currentNode.link[TRIE_FINAL_NODE_KEY]
		return isFinalKeyPresent
	}
}
