package main

import (
	"data-structures-algorithms/tree"
	"fmt"
)

func sumTree(node *tree.Node) int {
	if node == nil {
		return 0
	}

	return node.Value + sumTree(node.Left) + sumTree(node.Right)
}

func isSumTree(node *tree.Node) bool {
	if node == nil || (node.Left == nil && node.Right == nil) {
		return true
	}

	l := sumTree(node.Left)
	r := sumTree(node.Right)

	return node.Value == l+r && isSumTree(node.Left) && isSumTree(node.Right)
}

func main() {
	arr := []int{26, 10, 3, 4, 6, 3}
	bt := tree.NewBinaryTree()
	for _, v := range arr {
		bt.Insert(v)
	}

	fmt.Println(isSumTree(bt.Node))
}
