package tree

// This is basically all that's required.
type Trie struct {
	BaseTree
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
	for _, char := range ipStr {
		debug(char)
		//TODO:
		// Steps
		//- Iterate through each element
		//- Create nodes for em
		//- Add them in the right location

	}
	return true
}
