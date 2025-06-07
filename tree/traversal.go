package tree

import "fmt"

func InorderTraversal(root *Node) {
	if root == nil {
		return
	}

	InorderTraversal(root.Left)
	fmt.Println(root.Value)
	InorderTraversal(root.Right)
}

func PreorderTraversal(root *Node) {
	if root == nil {
		return
	}

	fmt.Println(root.Value)
	PreorderTraversal(root.Left)
	PreorderTraversal(root.Right)
}

func PostorderTraversal(root *Node) {
	if root == nil {
		return
	}

	PostorderTraversal(root.Left)
	PostorderTraversal(root.Right)
	fmt.Println(root.Value)
}
