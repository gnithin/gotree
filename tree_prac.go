package main

import (
	//"fmt"
	//"gotree/helpers"
	tree "gotree/tree"
	//"io/ioutil"
)

func main() {
	//tree.CreateTreeNode(1)
	treeObj := tree.CreateTree()
	treeObj.Insert(3)
	treeObj.Insert(5)
	treeObj.Insert(1)
	treeObj.Insert(10)
	treeObj.Insert(7)
	treeObj.Insert(6)
	treeObj.Insert(12)
	treeObj.Insert(4)

	/*
		checkHasVal := []int{
			11, 121, 3, 4, 5,
		}

		for _, val := range checkHasVal {
			fmt.Printf("Checking if %d is in the tree - %v\n", val, treeObj.HasVal(val))
		}

			jsonResp := treeObj.GetJSONTree()
			fmt.Println("Format - ", string(jsonResp))

			// write json to file
			destFileName := "autogen.json"
			destFilePath := "assets/data/" + destFileName
			writeErr := ioutil.WriteFile(destFilePath, jsonResp, 0644)
			if writeErr != nil {
				fmt.Println(writeErr)
				panic("Error writing json to file")
			} else {
				fmt.Println("Written to ", destFilePath)
			}

			// Needed to display the graph
			// TODO: Uncomment please
			helpers.CreateServer()
	*/
}
