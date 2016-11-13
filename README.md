# goTree

This repo consists of implementations of Trees in Go.
The Tree can store built-in types and custom types.

Currently covered are -
- BST
- Heap (Priority Queue)
- Trie

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

## Trie
Here's a basic example for using a Trie

```go
// Create a trieObject. It also allows some parameters
trieObj := tree.CreateTrie()

// Insert as many values as you want
trieObj.Insert("I", "am", "gonna", "love", "you", "till", "the", "heaven", "starts", "to", "rain")

// Search
trieObj.HasVal("gonna")  // true:  Exact match
trieObj.HasVal("GONna")  // true:  Case insensitive match
trieObj.HasVal("heav")   // true:  Partial match
trieObj.HasVal("absent") // false: Mismatch
```
The Trie is by default case insensitive and allows partial substring searches.
The case insensitivity and partial matching ability can be controlled by passing 
the valid parameters to the `CreateTrieWithOptions`

```go
// These two are true by default
partialMatch := false
caseInsensitive := false
trieOptObj := tree.CreateTrieWithOptions(partialMatch, caseInsensitive)
trieOptObj.Insert("Wherever", "I", "may", "roam", "where", "I", "lay", "my", "head", "is", "home")

trieOptObj.HasVal("wherever")  // false:  Case sensitive match 
trieOptObj.HasVal("Wherever")  // true
trieOptObj.HasVal("hea")       // false:  It's a partial match
trieOptObj.HasVal("head")      // true
```

You can also use the `CreateTrieWithOptionsMap` to try out all the flags used for a Trie.
The below example represents all the options available. 
This primarily showcases how to remove the stopwords before insertion, using the `strip_stopwords`.

```go
options := map[string]bool{
    "case_insensitive":   false,
    "partial_match":      false,
    "strip_punctuations": true,
    "strip_stopwords":    true,
}

trieObj := tree.CreateTrieWithOptionsMap(options)

trieObj.InsertStr(
    `Darkness, Imprisoning me.
     All that I see, Absolute horror! 
     I cannot live, 
     I cannot die,
     Trapped in myself, 
     Body my holding cell!`,
)

trieObj.HasVal("i")         // False: Stopword
trieObj.HasVal("my")        // False: Stopword
trieObj.HasVal("absolute")  // False: Case-Insensitive
trieObj.HasVal("Body")      // True 
trieObj.HasVal("Absolute")  // True
```
