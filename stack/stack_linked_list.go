package stack

import linkedlist "data-structures-algorithms/linkedList"

type StackLinkedList struct {
	Node *linkedlist.Node
}

func NewStackLinkedList() *StackLinkedList {
	return &StackLinkedList{}
}

func (s *StackLinkedList) Push(v int) {
	newNode := &linkedlist.Node{
		Data: v,
		Next: s.Node,
	}

	s.Node = newNode
}

func (s *StackLinkedList) Pop() int {
	if s.Node == nil {
		return -1
	}

	curr := s.Node

	s.Node = curr.Next

	return curr.Data
}

func (s *StackLinkedList) IsEmpty() bool {
	return s.Node == nil
}
