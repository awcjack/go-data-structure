package queue_test

import (
	"testing"

	"github.com/awcjack/go-data-structure/queue"
)

func TestLinkedListQueue_New(t *testing.T) {
	type testcase struct {
		test          string
		size          int
		expectedError error
	}

	testcases := []testcase{
		{
			test:          "Valid queue",
			size:          1,
			expectedError: nil,
		},
		{
			test:          "Empty valid queue",
			size:          0,
			expectedError: nil,
		},
		{
			test:          "Invalid queue",
			size:          -1,
			expectedError: queue.ErrUnderflow,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := queue.NewLinkedListQueue(tc.size)
			if err != tc.expectedError {
				t.Errorf("Expected get %v, but got %v", tc.expectedError, err)
			}
		})
	}
}

func TestLinkedListQueue_Full(t *testing.T) {
	type testcase struct {
		test           string
		size           int
		enqueue        int
		expectedResult bool
	}

	testcases := []testcase{
		{
			test:           "Empty queue",
			size:           0,
			enqueue:        0,
			expectedResult: true,
		},
		{
			test:           "Queue filled element",
			size:           1,
			enqueue:        1,
			expectedResult: true,
		},
		{
			test:           "Queue with element but not filled",
			size:           2,
			enqueue:        1,
			expectedResult: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			q, err := queue.NewLinkedListQueue(tc.size)
			if err != nil {
				t.Fatal(err)
			}
			for i := 0; i < tc.enqueue; i++ {
				err := q.Enqueue(i)
				if err != nil {
					t.Fatal(err)
				}
			}
			result := q.IsFull()
			if result != tc.expectedResult {
				t.Errorf("Expected get %v, but got %v", tc.expectedResult, result)
			}
		})
	}
}

func TestLinkedListQueue_Empty(t *testing.T) {
	type testcase struct {
		test           string
		size           int
		enqueue        int
		expectedResult bool
	}

	testcases := []testcase{
		{
			test:           "Empty queue",
			size:           0,
			enqueue:        0,
			expectedResult: true,
		},
		{
			test:           "Queue filled element",
			size:           1,
			enqueue:        1,
			expectedResult: false,
		},
		{
			test:           "Queue with element but not filled",
			size:           2,
			enqueue:        1,
			expectedResult: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			q, err := queue.NewLinkedListQueue(tc.size)
			if err != nil {
				t.Fatal(err)
			}
			for i := 0; i < tc.enqueue; i++ {
				err := q.Enqueue(i)
				if err != nil {
					t.Fatal(err)
				}
			}
			result := q.IsEmpty()
			if result != tc.expectedResult {
				t.Errorf("Expected get %v, but got %v", tc.expectedResult, result)
			}
		})
	}
}

func TestLinkedListQueue_Size(t *testing.T) {
	type testcase struct {
		test           string
		size           int
		enqueue        int
		expectedResult int
	}

	testcases := []testcase{
		{
			test:           "Empty queue",
			size:           0,
			enqueue:        0,
			expectedResult: 0,
		},
		{
			test:           "Queue filled element",
			size:           1,
			enqueue:        1,
			expectedResult: 1,
		},
		{
			test:           "Queue with element but not filled",
			size:           2,
			enqueue:        1,
			expectedResult: 1,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			s, err := queue.NewLinkedListQueue(tc.size)
			if err != nil {
				t.Fatal(err)
			}
			for i := 0; i < tc.enqueue; i++ {
				err := s.Enqueue(i)
				if err != nil {
					t.Fatal(err)
				}
			}
			result := s.Size()
			if result != tc.expectedResult {
				t.Errorf("Expected get %v, but got %v", tc.expectedResult, result)
			}
		})
	}
}

func TestLinkedListQueue_Enqueue(t *testing.T) {
	type testcase struct {
		test          string
		size          int
		enqueue       int
		expectedError error
	}

	testcases := []testcase{
		{
			test:          "Empty stack",
			size:          0,
			enqueue:       1,
			expectedError: queue.ErrOverflow,
		},
		{
			test:          "Fill Stack",
			size:          1,
			enqueue:       1,
			expectedError: nil,
		},
		{
			test:          "Fill Stack too much",
			size:          2,
			enqueue:       4,
			expectedError: queue.ErrOverflow,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			s, err := queue.NewLinkedListQueue(tc.size)
			if err != nil {
				t.Fatal(err)
			}
			for i := 0; i < tc.enqueue; i++ {
				err = s.Enqueue(i)
				if tc.expectedError != nil && err != nil && tc.expectedError != err {
					t.Errorf("Expected get %v, but got %v", tc.expectedError, err)
				}
				if tc.expectedError == nil && err != nil {
					t.Fatal(err)
				}
			}
			if tc.expectedError != nil && err != nil && tc.expectedError != err {
				t.Errorf("Expected get %v, but got %v", tc.expectedError, err)
			}
		})
	}
}

func TestLinkedListQueue_Dequeue(t *testing.T) {
	type testcase struct {
		test          string
		size          int
		enqueue       int
		dequeue       int
		expectedError error
	}

	testcases := []testcase{
		{
			test:          "Empty queue",
			size:          0,
			enqueue:       0,
			dequeue:       1,
			expectedError: queue.ErrUnderflow,
		},
		{
			test:          "Filled Queue",
			size:          1,
			enqueue:       1,
			dequeue:       1,
			expectedError: nil,
		},
		{
			test:          "Filled Queue",
			size:          2,
			enqueue:       2,
			dequeue:       1,
			expectedError: nil,
		},
		{
			test:          "Dequeue too much",
			size:          2,
			enqueue:       1,
			dequeue:       2,
			expectedError: queue.ErrUnderflow,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			s, err := queue.NewLinkedListQueue(tc.size)
			if err != nil {
				t.Fatal(err)
			}
			for i := 0; i < tc.enqueue; i++ {
				err = s.Enqueue(i)
				if err != nil {
					t.Fatal(err)
				}
			}
			for i := 0; i < tc.dequeue; i++ {
				_, err = s.Dequeue()
				if tc.expectedError != nil && err != nil && tc.expectedError != err {
					t.Errorf("Expected get %v, but got %v", tc.expectedError, err)
				}
				if tc.expectedError == nil && err != nil {
					t.Fatal(err)
				}
			}
			if tc.expectedError != nil && err != nil && tc.expectedError != err {
				t.Errorf("Expected get %v, but got %v", tc.expectedError, err)
			}
		})
	}
}
