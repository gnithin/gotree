package testSuite

import (
	"github.com/stretchr/testify/assert"
	"gotree/tree"
	"sort"
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

func TestHeap_integer(t *testing.T) {
	assert := assert.New(t)
	maxHeapObj := tree.CreateMaxHeap()
	minHeapObj := tree.CreateMinHeap()

	ipArr := []int{
		10001, 22, 1002, 101, 11,
		32, 48, 54,
	}

	// Hack to make the []int unpack into an interface{}
	var ipInterfaceArr []interface{}
	var ipInterafaceVal interface{}

	for _, ipVal := range ipArr {
		ipInterafaceVal = ipVal
		ipInterfaceArr = append(
			ipInterfaceArr,
			ipInterafaceVal,
		)
	}

	minHeapObj.Insert(ipInterfaceArr...)
	maxHeapObj.Insert(ipInterfaceArr...)

	// Getting the heap
	opMinHeapArr := popAll(minHeapObj)
	opMaxHeapArr := popAll(maxHeapObj)

	// Compare the popped value and sorted Value
	sort.Ints(ipArr)
	assert.True(compareIntSlices(ipArr, opMinHeapArr))

	sort.Sort(sort.Reverse(sort.IntSlice(ipArr)))
	assert.True(compareIntSlices(ipArr, opMaxHeapArr))

	assert.False(maxHeapObj.IsFull())
}

func popAll(heapVal *tree.Heap) []int {
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
