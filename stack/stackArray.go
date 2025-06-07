package stack

import "fmt"

type StackArr[T any] struct {
	empty T
	val   []T
}

func NewStackArr[T any]() *StackArr[T] {
	return &StackArr[T]{}
}

func (s *StackArr[T]) Push(v T) {
	s.val = append(s.val, v)
}

func (s *StackArr[T]) Length() int {
	return len(s.val)
}

func (s *StackArr[T]) Pop() T {
	if len(s.val) == 0 {
		return s.empty
	}

	l := s.val[len(s.val)-1]
	s.val = s.val[:len(s.val)-1]
	return l
}

func (s *StackArr[T]) Top() T {
	if len(s.val) == 0 {
		return s.empty
	}

	return s.val[len(s.val)-1]
}

func (s *StackArr[T]) String() string {
	return fmt.Sprintf("%v", s.val)
}

func (s *StackArr[T]) IsEmpty() bool {
	return len(s.val) == 0
}
