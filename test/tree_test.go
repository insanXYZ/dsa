package test

import (
	"data-structures-algorithms/tree"
	"testing"
)

var bt *tree.BinaryTree

func init() {
	arr := []int{1, 2, 3}
	bt = tree.NewBinaryTree()

	for _, v := range arr {
		bt.Insert(v)
	}

}

func TestPostorder(t *testing.T) {
	tree.PostorderTraversal(bt.Node)
}

func TestPreorder(t *testing.T) {
	tree.PreorderTraversal(bt.Node)
}

func TestInorder(t *testing.T) {
	tree.InorderTraversal(bt.Node)
}

func TestDelete(t *testing.T) {
	bt.Delete(3)
	tree.PostorderTraversal(bt.Node)
}
