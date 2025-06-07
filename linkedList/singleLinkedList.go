package linkedlist

type SingleLinkedList struct {
	Node *Node
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{}
}

func (l *SingleLinkedList) Insert(v int) {
	newNode := &Node{
		Data: v,
	}

	curr := l.Node

	if curr != nil {

		for curr.Next != nil {
			curr = curr.Next
		}

		curr.Next = newNode

	} else {
		l.Node = newNode
	}
}

func (l *SingleLinkedList) Search(key int) bool {

	for curr := l.Node; curr != nil; curr = curr.Next {
		if curr.Data == key {
			return true
		}
	}

	return false
}

func (l *SingleLinkedList) Length() int {
	var length int
	for curr := l.Node; curr != nil; curr = curr.Next {
		length++
	}
	return length
}
