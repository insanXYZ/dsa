package main

import (
	"data-structures-algorithms/tree"
	"fmt"
)

func identicalTree(node1, node2 *tree.Node) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil {
		return false
	}

	return node1.Value == node2.Value && identicalTree(node1.Left, node2.Left) && identicalTree(node1.Right, node2.Right)
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	// arr2 := arr1
	arr2 := []int{1, 2, 3, 5, 4}

	bt1 := tree.NewBinaryTree()
	bt2 := tree.NewBinaryTree()

	for _, v := range arr1 {
		bt1.Insert(v)
	}

	for _, v := range arr2 {
		bt2.Insert(v)
	}

	fmt.Println(identicalTree(bt1.Node, bt2.Node))
}
