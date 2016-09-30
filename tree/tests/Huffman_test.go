package testSuite

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gotree/tree"
	"testing"
)

func TestHuffman_basic(t *testing.T) {
	assert := assert.New(t)

	// Insert map
	/*
		freqMap := map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
			"d": 4,
			"e": 5,
		}
	*/
	original_string := "abbcccddddeeeee"
	expectedResult := "000001001010101101010101111111111"

	// Create huffman tree
	freqMap := tree.CreateFreqMap(original_string)
	huffmanTree := tree.CreateHuffmanTree(freqMap)

	// Encode string
	encodedStr := huffmanTree.EncodeStr(original_string)

	assert.Equal(encodedStr, expectedResult)

	// Decode string
	decodedStr := huffmanTree.DecodeStr(encodedStr)

	assert.Equal(original_string, decodedStr)
}

func TestHuffman_lyrics(t *testing.T) {
	ipStr := "Say your prayers little one, don't forget my son, to include everyone."

	// Create a frequency map
	freqMap := tree.CreateFreqMap(ipStr)
	fmt.Println(freqMap)

	// Create huffman tree
	huffmanTree := tree.CreateHuffmanTree(freqMap)

	// Encode it.
	fmt.Println(huffmanTree.EncodeStr(ipStr))
}
