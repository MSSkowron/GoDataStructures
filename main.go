package main

import (
	binarysearchtree "github.com/MSSkowron/GoDataStructuresAndAlgorithms/binary_search_tree"
)

func main() {
	bst := binarysearchtree.NewNode[int](8)
	bst.Insert(3)
	bst.Insert(1)
	bst.Insert(6)
	bst.Insert(4)
	bst.Insert(7)
	bst.Insert(10)
	bst.Insert(14)
	bst.Insert(13)

	bst.InorderTraversal()
}
