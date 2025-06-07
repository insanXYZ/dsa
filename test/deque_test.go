package test

import (
	"data-structures-algorithms/deque"
	"fmt"
	"testing"
)

func TestDeque(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	d := deque.NewDeque()

	for _, v := range arr {
		if v%2 == 0 {
			d.InsertFront(v)
		} else {
			d.InsertRear(v)
		}
	}

	for !d.IsEmpty() {
		fmt.Println(d.DeleteFront())
	}
}
