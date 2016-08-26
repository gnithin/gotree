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

	fmt.Println("Aas")

	//resp := t.GetJSONTree()
	//fmt.Println("Format - ", resp)

	// Needed to display the graph
	// TODO: Uncomment please
	//helpers.CreateServer()
}
