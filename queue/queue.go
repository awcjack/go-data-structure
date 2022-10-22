package queue

import "errors"

var (
	ErrUnderflow = errors.New("stack underflow")
	ErrOverflow  = errors.New("stack overflow")
)

type Queue struct {
	array    []interface{}
	capacity int
}

func New(capacity int) (*Queue, error) {
	if capacity < 0 {
		return nil, ErrUnderflow
	}
	return &Queue{
		array:    make([]interface{}, 0, capacity),
		capacity: capacity,
	}, nil
}

func (q Queue) IsFull() bool {
	return len(q.array) == q.capacity
}

func (q Queue) IsEmpty() bool {
	return len(q.array) == 0
}

func (q Queue) Size() int {
	return len(q.array)
}

func (q *Queue) Enqueue(element interface{}) error {
	if len(q.array) >= q.capacity {
		return ErrOverflow
	}
	q.array = append(q.array, element)
	return nil
}

func (q *Queue) Dequeue() (interface{}, error) {
	if len(q.array) == 0 {
		return nil, ErrUnderflow
	}
	temp := q.array[0]
	for i := 0; i < len(q.array)-1; i++ {
		q.array[i] = q.array[i+1]
	}
	q.array = q.array[:len(q.array)-1]
	return temp, nil
}
