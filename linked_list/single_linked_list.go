package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	Node *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) insert(v int) {
	newNode := &Node{
		data: v,
	}

	curr := l.Node

	if curr != nil {

		for curr.next != nil {
			curr = curr.next
		}

		curr.next = newNode

	} else {
		l.Node = newNode
	}
}

func (l *LinkedList) search(key int) bool {

	for curr := l.Node; curr != nil; curr = curr.next {
		if curr.data == key {
			return true
		}
	}

	return false
}

func (l *LinkedList) print() {
	for curr := l.Node; curr != nil; curr = curr.next {
		fmt.Println(curr.data)
	}
}

func (l *LinkedList) length() int {
	var length int
	for curr := l.Node; curr != nil; curr = curr.next {
		length++
	}
	return length
}

func main() {
	arr := []int{1, 2, 3}

	ll := NewLinkedList()

	for _, v := range arr {
		ll.insert(v)
	}

	fmt.Println("exist", ll.search(9))
	fmt.Println("length", ll.length())
	ll.print()

}
