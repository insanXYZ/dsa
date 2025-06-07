package queue

type Queue struct {
	arr []int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(v int) {
	q.arr = append(q.arr, v)
}

func (q *Queue) Dequeue() int {
	if len(q.arr) == 0 {
		return -1
	}

	d := q.arr[0]

	q.arr = q.arr[1:]
	return d
}

func (q *Queue) GetFront() int {
	if len(q.arr) == 0 {
		return -1
	}

	return q.arr[0]
}

func (q *Queue) GetRear() int {
	if q.Size() == 0 {
		return -1
	}

	return q.arr[len(q.arr)-1]
}

func (q *Queue) IsEmpty() bool {
	return len(q.arr) == 0
}

func (q *Queue) Size() int {
	return len(q.arr)
}
