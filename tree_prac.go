package main

import (
	"fmt"
	//"gotree/helpers"
	tree "gotree/tree"
)

func main() {
	t := tree.MakeNode(tree.CreateTreeData(22))
	c := tree.MakeNode(tree.CreateTreeData(2211))
	t.AddChild("left", c)
	resp := t.GetJSONTree()
	fmt.Println("Format - ", resp)

	// Needed to display the graph
	// TODO: Uncomment please
	//helpers.CreateServer()
}
