package testSuite

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gotree/tree"
	"sort"
	"testing"
)

func compareIntSlices(arr1, arr2 []int) bool {
	for index, _ := range arr1 {
		if arr1[index] != arr2[index] {
			fmt.Println("Not matching", arr1[index], arr2[index])
			return false
		}
	}
	return true
}

func TestHeap_integer(t *testing.T) {
	assert := assert.New(t)
	maxHeapObj := tree.CreateMaxHeap()
	minHeapObj := tree.CreateMinHeap()

	ipArr := []int{
		10001, 22, 1002, 101, 11,
		32, 48, 54,
	}
	for _, val := range ipArr {
		minHeapObj.Insert(val)
		maxHeapObj.Insert(val)
	}
	opMinHeapArr := getHeapOp(minHeapObj)
	opMaxHeapArr := getHeapOp(maxHeapObj)

	// Compare the popped value and sorted Value
	sort.Ints(ipArr)
	assert.True(compareIntSlices(ipArr, opMinHeapArr))

	sort.Sort(sort.Reverse(sort.IntSlice(ipArr)))
	assert.True(compareIntSlices(ipArr, opMaxHeapArr))
}

func getHeapOp(heapVal *tree.Heap) []int {
	var opArr []int
	var poppedVal int
	heapLen := heapVal.GetHeapLen()
	for i := 0; i < heapLen; i++ {
		poppedTempVal, popFlag := heapVal.Pop()
		if poppedTempVal != nil && popFlag {
			poppedVal = (*poppedTempVal).(int)
			opArr = append(
				opArr,
				poppedVal,
			)
		}
	}
	return opArr
}
