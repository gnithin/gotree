package main

import (
	//"fmt"
	//"gotree/helpers"
	tree "gotree/tree"
	//"io/ioutil"
)

func main() {
	//tree.CreateTreeNode(1)
	stringTree := tree.CreateTree()
	// Testing the strings
	///*
	stringTree.Insert("hey")
	stringTree.Insert("Oh")
	stringTree.Insert("Listen")
	stringTree.Insert("what")
	stringTree.Insert("I")
	stringTree.Insert("say")
	stringTree.Insert("Oh")
	//*/

	// Testing the numbers
	///*
	intTree := tree.CreateTree()
	intTree.Insert(3)
	intTree.Insert(5)
	intTree.Insert(1)
	intTree.Insert(10)
	intTree.Insert(7)
	intTree.Insert(6)
	intTree.Insert(12)
	intTree.Insert(4)
	//*/

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
