package deque

import (
	"fmt"
	"slices"
)

type Deque struct {
	arr []int
}

func NewDeque() *Deque {
	return &Deque{}
}

func (d *Deque) IsEmpty() bool {
	return len(d.arr) == 0
}

func (d *Deque) String() string {
	return fmt.Sprintf("%v", d.arr)
}

func (d *Deque) InsertFront(v int) {
	d.arr = slices.Concat(d.arr, []int{v})
}

func (d *Deque) InsertRear(v int) {
	d.arr = slices.Concat([]int{v}, d.arr)
}

func (d *Deque) DeleteFront() int {
	f := d.arr[len(d.arr)-1]
	d.arr = d.arr[:len(d.arr)-1]

	return f
}

func (d *Deque) DeleteRear() int {

	r := d.arr[0]
	d.arr = d.arr[1:]

	return r
}

func (d *Deque) FrontEle() int {
	return d.arr[len(d.arr)-1]
}

func (d *Deque) RearEle() int {
	return d.arr[0]
}
