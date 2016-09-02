package testSuite

// TODO: Benchmark stuff as well

import (
	"github.com/stretchr/testify/assert"
	"gotree/tree"
	"testing"
)

func TestBST(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(1, 1)
	numberTree := tree.CreateBST()
	numberTree.Insert(3, 5, 1)

}
