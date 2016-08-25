package main

import (
	"fmt"
	tree "tree_prac/tree"
)

func main() {
	t := tree.MakeNode(tree.CreateTreeData(22))
	c := tree.MakeNode(tree.CreateTreeData(2211))
	t.AddChild("left", c)
	fmt.Println(t)
}
