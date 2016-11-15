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

// This is a function returns the trie containing all the stop words
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
