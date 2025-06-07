package test

import (
	"data-structures-algorithms/queue"
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	arr := []int{1, 2, 3}
	q := queue.NewQueue()

	for _, v := range arr {
		q.Enqueue(v)
	}

	for !q.IsEmpty() {
		fmt.Println(q.Dequeue())
	}

}
