package testSuite

import (
	//"fmt"
	"github.com/stretchr/testify/assert"
	"gotree/tree"
	//"sort"
	"strings"
	"testing"
)

var bigFilePath string = "resources/shakespeare_works.txt"
var passSearchKey string = "Swinstead"
var failSearchKey string = "Nihtin"

func TestTrie_development(t *testing.T) {
	assert := assert.New(t)

	// Create a Trie
	trieObj := tree.CreateTrie()
	assert.NotEmpty(trieObj)

	// Insertion
	assert.True(trieObj.Insert("basic Str"))

	// Searching
	assert.True(trieObj.HasVal("basic Str"))
	assert.True(trieObj.HasVal("BASIC STR"))
	assert.True(trieObj.HasVal("basic"))
	assert.False(trieObj.HasVal("avengers"))

	// Create a trie with custom options
	// No substring matching. Only complete words.
	// Case sensitive
	trieObjOpt := tree.CreateTrieWithOptions(false, false)
	assert.NotEmpty(trieObjOpt)

	// Insertion
	assert.True(trieObjOpt.Insert("New String"))

	// Searching
	assert.True(trieObjOpt.HasVal("New String"))
	assert.False(trieObjOpt.HasVal("BASIC STRING"))
	assert.False(trieObjOpt.HasVal("New"))
}

func TestTrie_multiple(t *testing.T) {
	assert := assert.New(t)

	// Create a simple trie
	trieObj := tree.CreateTrie()

	// Inserting
	insertStatus := trieObj.Insert(
		"Letme",
		"tell",
		"you",
		"a",
		"story",
		"to",
		"chill",
		"your",
		"bones",
		"About",
		"a",
		"thing",
		"that",
		"I",
		"saw",
	)
	assert.True(insertStatus)

	assert.True(trieObj.HasVal("tel"))
	assert.True(trieObj.HasVal("let"))
	assert.True(trieObj.HasVal("a"))
	assert.True(trieObj.HasVal("i"))
	assert.True(trieObj.HasVal("chil"))
	assert.True(trieObj.HasVal("bones"))

	assert.False(trieObj.HasVal("tella"))
	assert.False(trieObj.HasVal(""))
	assert.False(trieObj.HasVal("wooot"))
	assert.False(trieObj.HasVal("chiller"))
}

func TestTrie_withOptionsMap(t *testing.T) {
	assert := assert.New(t)

	// Creating a tree with options
	options := map[string]bool{
		"partial_match":      false,
		"case_insensitive":   false,
		"strip_stopwords":    false,
		"strip_punctuations": true,
	}

	trieObj := tree.CreateTrieWithOptionsMap(options)
	trieObj.Insert(
		"Orion", "is", "a", "freaking", "masterpiece!",
	)

	assert.False(trieObj.HasVal("orion"))
	assert.False(trieObj.HasVal("freak"))
	assert.True(trieObj.HasVal("masterpiece"))
}

func TestTrie_stopWords(t *testing.T) {
	assert := assert.New(t)

	options := map[string]bool{
		"partial_match":      false,
		"case_insensitive":   false,
		"strip_stopwords":    true,
		"strip_punctuations": true,
	}

	trieObj := tree.CreateTrieWithOptionsMap(options)

	insertStatus := trieObj.InsertStr(
		`Darkness, Imprisoning me.
		 All that I see, Absolute horror! 
		 I cannot live, 
		 I cannot die,
		 Trapped in myself, 
		 Body my holding cell!`,
	)

	assert.True(insertStatus)
	assert.False(trieObj.HasVal("i"))
	assert.False(trieObj.HasVal("my"))
	assert.False(trieObj.HasVal("absolute"))
	assert.True(trieObj.HasVal("Absolute"))
	assert.False(trieObj.HasVal("body"))
}

func TestTrie_Insertion(t *testing.T) {
	assert := assert.New(t)

	fileContents := getFileContentsAsString(bigFilePath)
	trieObj := tree.CreateTrie()

	insStatus := trieObj.InsertStr(fileContents)
	t.Log(trieObj.GetLen())

	assert.True(insStatus)
}

// Let's Benchmark
func Benchmark_trieInsertion(b *testing.B) {
	options := map[string]bool{
		"partial_match":      false,
		"case_insensitive":   false,
		"strip_stopwords":    false,
		"strip_punctuations": false,
	}

	// Read a big file
	fileContents := getFileContentsAsString(bigFilePath)

	for i := 0; i < b.N; i++ {
		trieObj := tree.CreateTrieWithOptionsMap(options)
		trieObj.InsertStr(fileContents)
	}
}

func Benchmark_triePassSearch(b *testing.B) {
	trieObj := createTrieWithBigFile()

	for i := 0; i < b.N; i++ {
		trieObj.HasVal(passSearchKey)
	}
}

func Benchmark_strPassSearch(b *testing.B) {
	fileContents := getFileContentsAsString(bigFilePath)
	for i := 0; i < b.N; i++ {
		strings.Index(fileContents, passSearchKey)
	}
}

func Benchmark_trieFailSearch(b *testing.B) {
	trieObj := createTrieWithBigFile()

	// Just assert the expected response once before the benchmark
	assert := assert.New(b)
	assert.False(trieObj.HasVal(failSearchKey))

	for i := 0; i < b.N; i++ {
		trieObj.HasVal(failSearchKey)
	}
}

func Benchmark_strFailSearch(b *testing.B) {
	assert := assert.New(b)

	// Just assert the expected response once before the benchmark
	fileContents := getFileContentsAsString(bigFilePath)
	assert.Equal(strings.Index(fileContents, failSearchKey), -1)

	for i := 0; i < b.N; i++ {
		strings.Index(fileContents, failSearchKey)
	}
}
