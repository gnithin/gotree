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
			return false
		}
	}
	return true
}

func BenchmarkMaxHeap_ins_50000(b *testing.B) {
	benchmarkHeap_insertion_k(50000, b, true)
}

func BenchmarkMaxHeap_ins_90000(b *testing.B) {
	benchmarkHeap_insertion_k(90000, b, true)
}

func BenchmarkMaxHeap_ins_100000(b *testing.B) {
	benchmarkHeap_insertion_k(100000, b, true)
}

func BenchmarkMaxHeap_ins_200000(b *testing.B) {
	benchmarkHeap_insertion_k(200000, b, true)
}

func BenchmarkMaxHeap_ins_300000(b *testing.B) {
	benchmarkHeap_insertion_k(300000, b, true)
}

func BenchmarkMaxHeap_ins_400000(b *testing.B) {
	benchmarkHeap_insertion_k(400000, b, true)
}

func BenchmarkMaxHeap_ins_500000(b *testing.B) {
	benchmarkHeap_insertion_k(500000, b, true)
}

func BenchmarkMinHeap_ins_50000(b *testing.B) {
	benchmarkHeap_insertion_k(50000, b, false)
}

func BenchmarkMinHeap_ins_90000(b *testing.B) {
	benchmarkHeap_insertion_k(90000, b, false)
}

func BenchmarkMinHeap_ins_100000(b *testing.B) {
	benchmarkHeap_insertion_k(100000, b, false)
}

func BenchmarkMinHeap_ins_200000(b *testing.B) {
	benchmarkHeap_insertion_k(200000, b, false)
}

func BenchmarkMinHeap_ins_300000(b *testing.B) {
	benchmarkHeap_insertion_k(300000, b, false)
}

func BenchmarkMinHeap_ins_400000(b *testing.B) {
	benchmarkHeap_insertion_k(400000, b, false)
}

func BenchmarkMinHeap_ins_500000(b *testing.B) {
	benchmarkHeap_insertion_k(500000, b, false)
}

func benchmarkHeap_insertion_k(k int, b *testing.B, isMaxHeap bool) {
	maxCount := k + 1

	var heapObj *tree.Heap

	if isMaxHeap {
		heapObj = tree.CreateMaxHeapWithSize(maxCount)
	} else {
		heapObj = tree.CreateMinHeapWithSize(maxCount)
	}

	for i := 0; i < k; i++ {
		heapObj.Insert(i)
	}

	fmt.Println(b.N)
	for i := 0; i <= b.N; i++ {
		heapObj.Insert(k)
	}
}

func BenchmarkMaxHeap_pop5000(b *testing.B) {
	benchmarkHeap_pop_k(5000, b, true)
}

func BenchmarkMaxHeap_pop20000(b *testing.B) {
	benchmarkHeap_pop_k(20000, b, true)
}

func BenchmarkMaxHeap_pop50000(b *testing.B) {
	benchmarkHeap_pop_k(50000, b, true)
}

func BenchmarkMaxHeap_pop90000(b *testing.B) {
	benchmarkHeap_pop_k(90000, b, true)
}

func BenchmarkMaxHeap_pop100000(b *testing.B) {
	benchmarkHeap_pop_k(100000, b, true)
}

func BenchmarkMaxHeap_pop200000(b *testing.B) {
	benchmarkHeap_pop_k(200000, b, true)
}

func BenchmarkMaxHeap_pop300000(b *testing.B) {
	benchmarkHeap_pop_k(300000, b, true)
}

func BenchmarkMaxHeap_pop400000(b *testing.B) {
	benchmarkHeap_pop_k(400000, b, true)
}

func BenchmarkMaxHeap_pop500000(b *testing.B) {
	benchmarkHeap_pop_k(500000, b, true)
}

// This method pops all
func benchmarkHeap_pop_k(k int, b *testing.B, isMaxHeap bool) {
	maxCount := k

	// First filling the heap to brim
	var heapObj *tree.Heap

	if isMaxHeap {
		heapObj = tree.CreateMaxHeapWithSize(maxCount)
	} else {
		heapObj = tree.CreateMinHeapWithSize(maxCount)
	}

	for i := 0; i < k; i++ {
		heapObj.Insert(i)
	}

	for i := 0; i < b.N; i++ {
		for j := maxCount; j > 0; j-- {
			heapObj.Pop()
		}
	}
}

func TestHeap_loadTesting(t *testing.T) {
	assert := assert.New(t)

	maxCount := 500
	// This is 5 million.
	// Crazily enough, this seems to work.
	// In 47.9 secs. Don't know how fast that is.
	// Any increase in maxCount seems to be linear.
	// No errors, nothing.

	// Pretty cool considering the whole project is
	// implemented using the recursive strategies of tree.

	// Procrastinating the custom stack implementation was a
	// good idea :p
	i := 0

	maxHeapObj := tree.CreateMaxHeapWithSize(maxCount)

	for i = 0; i < maxCount; i++ {
		insert_status := maxHeapObj.InsertOne(i)
		assert.True(insert_status)
	}

	for i > 0 {
		i -= 1
		poppedVal, _ := maxHeapObj.Pop()

		assert.Equal((*poppedVal).(int), i)
	}
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
