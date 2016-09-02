package tree

// TODO: Benchmark stuff as well

import (
	"github.com/stretchr/testify/assert"
	"gotree/tree"
	"testing"
)

func TestIntegerBST(t *testing.T) {
	assert := assert.New(t)

	numberTree := tree.CreateBST()
	assert.True(numberTree.Insert(3, 5, 1))
	assert.True(numberTree.InsertOne(10))
	assert.False(numberTree.Insert(5))

	// Compare the structure as well
}

func TestStringBST(t *testing.T) {
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
