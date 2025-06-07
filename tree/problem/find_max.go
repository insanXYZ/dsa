package main

import (
	"data-structures-algorithms/tree"
	"fmt"
)

func findMax(node *tree.Node) int {

	if node == nil {
		return 0
	}

	l := findMax(node.Left)
	r := findMax(node.Right)

	return max(node.Value, l, r)
}

func main() {
	arr := []int{25, 43, 400, 99, 231}

	bt := tree.NewBinaryTree()

	for _, v := range arr {
		bt.Insert(v)
	}

	fmt.Println(findMax(bt.Node))
}
