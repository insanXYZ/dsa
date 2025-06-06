package main

import (
	"fmt"
	"slices"
)

type Stack struct {
	arr []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) push(v int) {
	s.arr = slices.Concat([]int{v}, s.arr)
}

func (s *Stack) pop() int {
	f := s.arr[0]
	s.arr = s.arr[1:]

	return f
}

func main() {
	arr := []int{1, 2, 3}
	s := NewStack()

	for _, v := range arr {
		s.push(v)
	}

	for len(s.arr) != 0 {
		fmt.Println(s.pop())
	}
}
