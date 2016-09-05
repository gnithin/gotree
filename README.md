# goTree

This repo consists of implementations of Trees in Go.
The Tree can store built-in types and custom types.

Currently covered are -
- BST
- Heap (Priority Queue)

All the trees have three basic functionalities
- Insertion of new elements
- Pop the root

All the trees also have their own public methods, owing to their semantics and structures.

## BST

Here's a basic example of adding to a tree
```go
// Creating a new BST
numberTree := tree.CreateBST()

// Adding entries
// Supports integer and strings by default
numberTree.Insert(3, 5, 1, 10)

// Here's a BST from a string
stringTree := tree.CreateBST()
stringTree.Insert("Infinity", "and", "beyond!") 

// Other methods
numberTree.Remove(10)
stringTree.HasVal("Infinity")
```

This example shows how to add custom objects into the tree
```go
// Here is my custom object
type myObject struct {
	name string
	age  int
	sal  float64
}

// To create a BST with a custom object,
// a comparator is needed, according to which
// the BST will be structured

// A comparator function takes 2 args, compares them,
// and returns, -1, 0 and 1, for less than, equal to and greater
// than values. Ofcourse, you can tailor it to your needs.
// Here is a basic comparator function 
comparatorFunc := func(obj1, obj2 *interface{}) int {
    new_obj1 := (*obj1).(myObject)
    new_obj2 := (*obj2).(myObject)

    // This needn't be simple
    // Can fit any use case
    metric1 := new_obj1.sal + float64(20*(new_obj1.age))
    metric2 := new_obj2.sal + float64(20*(new_obj2.age))

    if metric1 < metric2 {
        return -1
    } else if metric1 > metric2 {
        return 1
    } 
    return 0
}

// Here's creating a tree with passing the comparator pointer
customObjTree := tree.CreateBSTWithComparator(&comparatorFunc)

// All the functions are available to this tree as well
customObjTree.Insert(
    myObject{name: "james", age: 51, sal: 230.000},
    myObject{name: "mustaine", age: 55, sal: 140.000},
    myObject{name: "tom", age: 20, sal: 1240.000},
)

poppedObj, isPopSuccess := customObjTree.Pop()
// Remember to type assert back your popped object
poppedMyObj := (*poppedObj).(myObject)

customObjTree.HasVal(
    myObject{name: "tom", age: 20, sal: 1240.000}
)
```
The above functionality for custom objects is uniform across 
different types of trees.

## Heap (Priority Queue)
Here's a basic set of functions in the Heap.

```go
heapObj := tree.CreateMinHeap() // tree.CreateMaxHeap() also exists
// Insertion
heapObj.Insert(10, 20, 1001, 1)

// Pop operation
minimumObj, isValExists := heapObj.Pop() // Returns 1
if isValExists {
    minimumNum := (*minimumObj).(int)
}
```
The heap works with strings and custom objects similar to how it was shown above using BST.

The [huffman tree](https://github.com/gnithin/gotree/blob/master/tree/huffman.go#L24-L38) internally uses a MinHeap by passing it's
custom structure and comparator.

### Things left to do
- Fix the UI part of creating a tree
- Change doc by adding more code and diagrams
- Create AVL trees
