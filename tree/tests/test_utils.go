package testSuite

import (
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
