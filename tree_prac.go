package main

import (
	"fmt"
	//"gotree/helpers"
	tree "gotree/tree"
)

func main() {
	treeObj := tree.CreateTree()
	treeObj.Insert(3)
	treeObj.Insert(5)
	treeObj.Insert(1)
	treeObj.Insert(10)
	treeObj.Insert(7)
	treeObj.Insert(5)
	treeObj.Insert(6)

	resp := treeObj.GetJSONTree()
	fmt.Println("Format - ", string(resp))

	// Needed to display the graph
	// TODO: Uncomment please
	//helpers.CreateServer()
}
