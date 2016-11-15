package testSuite

import (
	"gotree/tree"
	"io/ioutil"
)

func compareIntSlices(arr1, arr2 []int) bool {
	for index, _ := range arr1 {
		if arr1[index] != arr2[index] {
			return false
		}
	}
	return true
}

func checkForPanicTime(e error) {
	if e != nil {
		panic(e)
	}
}

func getFileContentsAsString(filePath string) string {
	data, err := ioutil.ReadFile(filePath)
	checkForPanicTime(err)
	return string(data)
}

func createTrieWithBigFile() *tree.Trie {
	options := map[string]bool{
		"partial_match":      false,
		"case_insensitive":   false,
		"strip_stopwords":    false,
		"strip_punctuations": false,
	}

	// Read a big file
	fileContents := getFileContentsAsString(bigFilePath)
	trieObj := tree.CreateTrieWithOptionsMap(options)
	trieObj.InsertStr(fileContents)
	return trieObj
}
