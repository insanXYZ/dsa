package queue

type Queue[T any] struct {
	arr   []T
	empty T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(v T) {
	q.arr = append(q.arr, v)
}

func (q *Queue[T]) Dequeue() T {
	if len(q.arr) == 0 {
		return q.empty
	}

	d := q.arr[0]

	q.arr = q.arr[1:]
	return d
}

func (q *Queue[T]) GetFront() T {
	if len(q.arr) == 0 {
		return q.empty
	}

	return q.arr[0]
}

func (q *Queue[T]) GetRear() T {
	if q.Size() == 0 {
		return q.empty
	}

	return q.arr[len(q.arr)-1]
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.arr) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.arr)
}
