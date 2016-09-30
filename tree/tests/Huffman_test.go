package testSuite

import (
	"github.com/stretchr/testify/assert"
	"gotree/tree"
	"testing"
)

func TestHuffman_basic(t *testing.T) {
	assert := assert.New(t)

	// Insert map
	freqMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}
	// Create huffman tree
	huffmanTree := tree.CreateHuffmanTree(freqMap)

	original_string := "abbcccddddeeeee"
	expectedResult := "000001001010101101010101111111111"

	// Encode string
	encodedStr := huffmanTree.EncodeStr(original_string)

	assert.Equal(encodedStr, expectedResult)

	// Decode string
	decodedStr := huffmanTree.DecodeStr(encodedStr)

	assert.Equal(original_string, decodedStr)
}
