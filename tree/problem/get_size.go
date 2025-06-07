package main

import (
	"data-structures-algorithms/tree"
	"fmt"
)

func getSize(root *tree.Node) int {
	if root == nil {
		return 0
	}

	l, r := getSize(root.Left), getSize(root.Right)

	return l + r + 1
}

func main() {
	root := &tree.Node{Value: 1}
	root.Left = &tree.Node{Value: 2}
	root.Right = &tree.Node{Value: 3}
	root.Left.Left = &tree.Node{Value: 4}
	root.Left.Right = &tree.Node{Value: 5}

	fmt.Println(getSize(root))
}
