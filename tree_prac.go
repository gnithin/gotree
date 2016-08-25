package main

import (
	"fmt"
	tree "tree_prac/tree"
)

func main() {
	t := tree.MakeNode(22)
	c := tree.MakeNode(221)
	t.AddChild("left", c)
}
