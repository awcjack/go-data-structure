package stack

import "errors"

var (
	ErrUnderflow = errors.New("stack underflow")
	ErrOverflow  = errors.New("stack overflow")
)

type Stack struct {
	capacity   int
	pointer    int
	stackArray []interface{}
}

func New(capacity int) (*Stack, error) {
	if capacity < 0 {
		return nil, ErrUnderflow
	}
	return &Stack{
		capacity:   capacity,
		pointer:    -1,
		stackArray: make([]interface{}, capacity),
	}, nil
}

func (s Stack) IsFull() bool {
	return s.pointer+1 == s.capacity
}

func (s Stack) IsEmpty() bool {
	return s.pointer == -1
}

func (s Stack) Size() int {
	return s.pointer + 1
}

func (s *Stack) Push(element interface{}) error {
	if s.IsFull() {
		return ErrOverflow
	}
	s.pointer++
	s.stackArray[s.pointer] = element
	return nil
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrUnderflow
	}
	s.pointer--
	return s.stackArray[s.pointer+1], nil
}

func (s Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrUnderflow
	}
	return s.stackArray[s.pointer], nil
}
