package testSuite

// TODO: Benchmark stuff as well

import (
	//"fmt"
	"github.com/stretchr/testify/assert"
	"gotree/tree"
	"testing"
)

func TestBST_integer(t *testing.T) {
	assert := assert.New(t)

	// Checking insertion
	numberTree := tree.CreateBST()
	assert.True(numberTree.Insert(3, 5, 1))
	assert.True(numberTree.InsertOne(10))

	assert.False(numberTree.Insert())
	assert.False(numberTree.Insert(5))
	assert.False(numberTree.InsertOne(10))

	// Checking for HasVal
	insertList := []int{1, 3, 5, 10}
	for _, val := range insertList {
		assert.True(numberTree.HasVal(val))
		assert.False(numberTree.HasVal(val + 1))
	}

	// Checking remove()
	deleteItems := []interface{}{1, 3}
	assert.True(numberTree.Remove(deleteItems...))
	assert.False(numberTree.Remove(11111))

	// Pop till you can't pop anymore
	popStatus := true
	numPops := 0
	for popStatus {
		_, popStatus = numberTree.Pop()
		numPops += 1
	}
	assert.EqualValues(3, numPops)
}

func TestBST_string(t *testing.T) {
	assert := assert.New(t)

	// Checking insertion
	stringTree := tree.CreateBST()
	assert.True(stringTree.Insert(
		"hello",
		"hey",
		"wazzap",
	))
	assert.True(stringTree.InsertOne("Ribbick"))

	assert.False(stringTree.Insert())
	assert.False(stringTree.Insert("hey"))
	assert.False(stringTree.InsertOne("Ribbick"))

	// Checking for HasVal
	insertList := []string{
		"hello",
		"hey",
		"wazzap",
	}
	for _, val := range insertList {
		assert.True(stringTree.HasVal(val))
		assert.False(stringTree.HasVal(val + "s"))
	}

	// Checking remove()
	deleteItems := []interface{}{"hello", "hey"}
	assert.True(stringTree.Remove(deleteItems...))
	assert.False(stringTree.Remove("RHCP"))

	// Pop till you can't pop anymore
	popStatus := true
	numPops := 0
	for popStatus {
		_, popStatus = stringTree.Pop()
		numPops += 1
	}
	assert.EqualValues(3, numPops)
}

type myObject struct {
	name string
	age  int
	sal  float64
}

func TestBST_customObj(t *testing.T) {
	assert := assert.New(t)

	comparatorFunc := func(obj1, obj2 *interface{}) int {
		new_obj1 := (*obj1).(myObject)
		new_obj2 := (*obj2).(myObject)

		// This needn't be simple
		metric1 := new_obj1.sal + float64(20*(new_obj1.age))
		metric2 := new_obj2.sal + float64(20*(new_obj2.age))

		if metric1 < metric2 {
			return -1
		} else if metric1 > metric2 {
			return 1
		} else {
			return 0
		}
	}

	// Testing out an invalid custom tree
	invalidCustomObjTree := tree.CreateBST()
	assert.False(invalidCustomObjTree.Insert(
		myObject{name: "invalid", age: 199, sal: 100.00},
	))

	// Adding a comparator function
	customObjTree := tree.CreateBSTWithComparator(&comparatorFunc)
	assert.True(customObjTree.Insert(
		myObject{name: "james", age: 51, sal: 230.000},
		myObject{name: "mustaine", age: 55, sal: 140.000},
		myObject{name: "tom", age: 20, sal: 1240.000},
		myObject{name: "jerry", age: 11, sal: 1140.000},
	))

	assert.True(customObjTree.InsertOne(
		myObject{name: "pluto", age: 10, sal: 1200.00},
	))

	// check hasVal
	assert.True(customObjTree.HasVal(
		myObject{name: "pluto", age: 10, sal: 1200.00},
	))

	// Remember, even equality undergoes the test
	// from the comparator
	assert.False(customObjTree.HasVal(
		myObject{name: "pluto", age: 10, sal: 12000.00},
	))

	// TODO: Test pop and remove as well
}
