package main

import (
	"fmt"
	"gotree/helpers"
	tree "gotree/tree"
)

func main() {
	t := tree.MakeNode(tree.CreateTreeData(22))
	c := tree.MakeNode(tree.CreateTreeData(2211))
	t.AddChild("left", c)
	fmt.Println(t)

	// Needed to display the graph
	helpers.CreateServer()
}
