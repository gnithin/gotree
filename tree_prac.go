package main

import (
	"fmt"
	"gotree/helpers"
	tree "gotree/tree"
	"io/ioutil"
	"os"
)

type myObject struct {
	name string
	age  int
	sal  float64
}

func main() {
	heapObj := tree.CreateHeap()
	heapObj.Insert(10)
	fmt.Println(heapObj)

	os.Exit(1)

	// Testing custom objects
	///*
	comparatorFunc := func(obj1, obj2 *interface{}) int {
		new_obj1 := (*obj1).(myObject)
		new_obj2 := (*obj2).(myObject)

		// This needn't be simple
		metric1 := new_obj1.sal + float64(20*(new_obj1.age))
		metric2 := new_obj2.sal + float64(20*(new_obj2.age))

		if metric1 < metric2 {
			return -1
		} else if metric1 > metric2 {
			return 1
		} else {
			return 0
		}
	}

	//customObjTree := tree.CreateTree()
	///*
	customObjTree := tree.CreateBSTWithComparator(&comparatorFunc)
	customObjTree.Insert(myObject{name: "james", age: 51, sal: 230.000})
	customObjTree.Insert(myObject{name: "mustaine", age: 55, sal: 140.000})
	customObjTree.Insert(myObject{name: "tom", age: 20, sal: 1240.000})
	customObjTree.Insert(myObject{name: "jerry", age: 11, sal: 1140.000})

	poppedObj, isPoppedExists := customObjTree.Pop()
	if isPoppedExists {
		fmt.Println((*poppedObj).(myObject))
	} else {
		fmt.Println("Pop failed!")
	}

	//*/

	// Testing the strings
	///*
	stringTree := tree.CreateBST()
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
	intTree := tree.CreateBST()
	intTree.Insert(3)
	intTree.Insert(5)
	intTree.Insert(1)
	intTree.Insert(10)
	intTree.Insert(7)
	intTree.Insert(6)
	intTree.Insert(12)
	intTree.Insert(4)

	intTree.Remove(10)
	intTree.Remove(1)
	intTree.Insert(2)
	intTree.Insert(8)

	intTree.Remove(5)
	intTree.Insert(22)
	//*/

	///*
	checkHasVal := []int{
		10, 1, 5, 22, 6,
	}

	for _, val := range checkHasVal {
		fmt.Printf("Checking if %d is in the tree - %v\n", val, intTree.HasVal(val))
	}

	///*
	jsonResp := intTree.GetJSONTree()
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
	//*/
}
