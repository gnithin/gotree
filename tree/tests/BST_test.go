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

	numberTree := tree.CreateBST()
	assert.True(numberTree.Insert(3, 5, 1))
	assert.True(numberTree.InsertOne(10))
	assert.False(numberTree.Insert(5))

	// Compare the structure as well
}

func TestBST_string(t *testing.T) {
	assert := assert.New(t)

	numberTree := tree.CreateBST()
	assert.True(numberTree.Insert(
		"hello",
		"hey",
		"wazzap",
	))
	assert.True(numberTree.InsertOne("Ribbick"))
	assert.False(numberTree.Insert("hey"))

	// Compare the structure as well
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
	assert.False(invalidCustomObjTree.Insert(myObject{name: "invalid", age: 199, sal: 100.00}))

	customObjTree := tree.CreateBSTWithComparator(&comparatorFunc)
	assert.True(customObjTree.Insert(
		myObject{name: "james", age: 51, sal: 230.000},
		myObject{name: "mustaine", age: 55, sal: 140.000},
		myObject{name: "tom", age: 20, sal: 1240.000},
		myObject{name: "jerry", age: 11, sal: 1140.000},
	))
}
