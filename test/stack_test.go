package test

import (
	"data-structures-algorithms/stack"
	"fmt"
	"testing"
)

func TestStackArray(t *testing.T) {
	arr := []int{1, 2, 3}

	s := stack.NewStackArr[int]()
	for _, v := range arr {
		s.Push(v)
	}

	fmt.Println(s)

	for !s.IsEmpty() {
		fmt.Println(s.Pop())
	}
}

func TestStackLinkedList(t *testing.T) {
	arr := []int{1, 2, 3}

	s := stack.NewStackLinkedList()

	for _, v := range arr {
		s.Push(v)
	}

	for !s.IsEmpty() {
		fmt.Println(s.Pop())
	}
}
