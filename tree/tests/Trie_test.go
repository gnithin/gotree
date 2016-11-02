package testSuite

import (
	//"fmt"
	"github.com/stretchr/testify/assert"
	"gotree/tree"
	//"sort"
	"testing"
)

func compareIntSlices(arr1, arr2 []int) bool {
	for index, _ := range arr1 {
		if arr1[index] != arr2[index] {
			return false
		}
	}
	return true
}

func TestTrie_development(t *testing.T) {
	assert := assert.New(t)
	trieObj := tree.CreateTrie()
	t.Log(trieObj)

	assert.NotEmpty(trieObj)

	// Insertion
	insStatus := trieObj.Insert("basic Str")
	t.Log(insStatus)
	assert.True(insStatus)

	// Searching
}
