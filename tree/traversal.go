package tree

import "fmt"

// left > root > right
func InorderTraversal(root *Node) {
	if root == nil {
		return
	}

	InorderTraversal(root.Left)
	fmt.Println(root.Value)
	InorderTraversal(root.Right)
}

// root > left > right
func PreorderTraversal(root *Node) {
	if root == nil {
		return
	}

	fmt.Println(root.Value)
	PreorderTraversal(root.Left)
	PreorderTraversal(root.Right)
}

// left > right > root
func PostorderTraversal(root *Node) {
	if root == nil {
		return
	}

	PostorderTraversal(root.Left)
	PostorderTraversal(root.Right)
	fmt.Println(root.Value)
}
