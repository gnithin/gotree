/*
Utils file for common stuff
*/
package tree

import (
	"fmt"
)

// This is a global var (gasp!).
// There, I said it.
// Shame me
var stopWordsTrie *Trie

/*
Helper function that's used to print out stuff when
debug is true
*/
func debug(ip ...interface{}) {
	if DEBUG {
		fmt.Println(ip...)
	}
}

func getTrieForStopWords() *Trie {
	if stopWordsTrie == nil {
		oldDebug := DEBUG
		DEBUG = false
		options := map[string]bool{
			"partial_match":      false,
			"case_insensitive":   false,
			"strip_stopwords":    false,
			"strip_punctuations": false,
		}
		stopWordsTrie = CreateTrieWithOptionsMap(options)

		// Insert list of stop words in this thing
		stopWordsTrie.Insert(stopWordsList...)

		DEBUG = oldDebug
	}

	return stopWordsTrie
}
