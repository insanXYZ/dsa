package tree

import (
	"data-structures-algorithms/queue"
)

type Node struct {
	Left  *Node
	Value int
	Right *Node
}

type BinaryTree struct {
	Node *Node
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (b *BinaryTree) Insert(v int) {

	newNode := &Node{
		Value: v,
	}

	if b.Node == nil {
		b.Node = newNode
		return
	}

	queue := queue.NewQueue[*Node]()
	queue.Enqueue(b.Node)

	for !queue.IsEmpty() {
		curr := queue.Dequeue()

		if curr.Left != nil {
			queue.Enqueue(curr.Left)
		} else {
			curr.Left = newNode
			return
		}

		if curr.Right != nil {
			queue.Enqueue(curr.Right)
		} else {
			curr.Right = newNode
			return
		}

	}

}

func (b *BinaryTree) deepestDelete(keyNode *Node) {

	queue := queue.NewQueue[*Node]()
	queue.Enqueue(b.Node)

	for !queue.IsEmpty() {
		curr := queue.Dequeue()

		if curr.Left != nil {
			if curr.Left == keyNode {
				curr.Left = nil
				return
			} else {
				queue.Enqueue(curr.Left)
			}
		}

		if curr.Right != nil {
			if curr.Right == keyNode {
				curr.Right = nil
				return
			} else {
				queue.Enqueue(curr.Right)
			}
		}

	}

}

func (b *BinaryTree) Delete(v int) {
	if b.Node == nil {
		return
	}

	if b.Node.Left == nil && b.Node.Right == nil && b.Node.Value == v {
		b.Node = nil
		return
	}

	var curr, keyNode *Node
	queue := queue.NewQueue[*Node]()

	queue.Enqueue(b.Node)

	for !queue.IsEmpty() {
		curr = queue.Dequeue()

		if curr.Value == v {
			keyNode = curr
		}

		if curr.Left != nil {
			queue.Enqueue(curr.Left)
		}

		if curr.Right != nil {
			queue.Enqueue(curr.Right)
		}
	}

	if keyNode != nil {
		currVal := curr.Value

		keyNode.Value = currVal

		b.deepestDelete(curr)
	}

}
