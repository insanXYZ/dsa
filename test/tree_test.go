package test

import (
	"data-structures-algorithms/tree"
	"testing"
)

var root *tree.Node

func init() {
	root = &tree.Node{
		Value: 1,
	}

	root.Left = &tree.Node{
		Value: 2,
	}

	root.Right = &tree.Node{
		Value: 3,
	}
}

func TestPostorder(t *testing.T) {
	tree.PostorderTraversal(root)
}

func TestPreorder(t *testing.T) {
	tree.PreorderTraversal(root)
}

func TestInorder(t *testing.T) {
	tree.InorderTraversal(root)
}
