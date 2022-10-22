package queue

type node struct {
	element interface{}
	next    *node
	prev    *node
}

type LinkedListQueue struct {
	first    *node
	last     *node
	capacity int
	counter  int
}

func NewLinkedListQueue(capacity int) (*LinkedListQueue, error) {
	if capacity < 0 {
		return nil, ErrUnderflow
	}
	return &LinkedListQueue{
		first:    nil,
		last:     nil,
		capacity: capacity,
		counter:  0,
	}, nil
}

func (q LinkedListQueue) IsFull() bool {
	return q.capacity == q.counter
}

func (q LinkedListQueue) IsEmpty() bool {
	return q.counter == 0
}

func (q LinkedListQueue) Size() int {
	return q.counter
}

func (q *LinkedListQueue) Enqueue(element interface{}) error {
	newNode := &node{
		element: element,
	}

	if q.first == q.last && q.first == nil {
		if q.capacity == 0 {
			return ErrOverflow
		}
		q.first = newNode
		q.last = newNode
		q.counter++
		return nil
	}

	if q.capacity == q.counter {
		return ErrOverflow
	}

	newNode.prev = q.last
	q.last = newNode
	q.counter++
	return nil
}

func (q *LinkedListQueue) Dequeue() (interface{}, error) {
	if q.first == q.last && q.first == nil {
		return nil, ErrUnderflow
	}

	temp := q.first.element
	if q.first == q.last {
		// only one element
		q.first = nil
		q.last = nil
		q.counter--
		return temp, nil
	}

	q.first = q.first.next
	return temp, nil
}
