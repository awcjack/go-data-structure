package stack_test

import (
	"testing"

	"github.com/awcjack/go-data-structure/stack"
)

func TestNewStack(t *testing.T) {
	type testcase struct {
		test          string
		size          int
		expectedError error
	}

	testcases := []testcase{
		{
			test:          "Valid stack",
			size:          1,
			expectedError: nil,
		},
		{
			test:          "Valid stack",
			size:          100,
			expectedError: nil,
		},
		{
			test:          "Invalid size",
			size:          -1,
			expectedError: stack.ErrUnderflow,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := stack.New(tc.size)
			if err != tc.expectedError {
				t.Errorf("Expected get %v, but got %v", tc.expectedError, err)
			}
		})
	}
}

func TestStack_Full(t *testing.T) {
	type testcase struct {
		test           string
		size           int
		push           int
		expectedResult bool
	}

	testcases := []testcase{
		{
			test:           "Empty stack",
			size:           0,
			push:           0,
			expectedResult: true,
		},
		{
			test:           "Stack filled with elements",
			size:           1,
			push:           1,
			expectedResult: true,
		},
		{
			test:           "Stack with element but not filled",
			size:           2,
			push:           1,
			expectedResult: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			s, err := stack.New(tc.size)
			if err != nil {
				t.Fatal(err)
			}
			for i := 0; i < tc.push; i++ {
				err := s.Push(i)
				if err != nil {
					t.Fatal(err)
				}
			}
			result := s.IsFull()
			if result != tc.expectedResult {
				t.Errorf("Expected get %v, but got %v", tc.expectedResult, result)
			}
		})
	}
}

func TestStack_Empty(t *testing.T) {
	type testcase struct {
		test           string
		size           int
		push           int
		expectedResult bool
	}

	testcases := []testcase{
		{
			test:           "Empty stack",
			size:           0,
			push:           0,
			expectedResult: true,
		},
		{
			test:           "Stack filled element",
			size:           1,
			push:           1,
			expectedResult: false,
		},
		{
			test:           "Stack with element but not filled",
			size:           2,
			push:           1,
			expectedResult: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			s, err := stack.New(tc.size)
			if err != nil {
				t.Fatal(err)
			}
			for i := 0; i < tc.push; i++ {
				err := s.Push(i)
				if err != nil {
					t.Fatal(err)
				}
			}
			result := s.IsEmpty()
			if result != tc.expectedResult {
				t.Errorf("Expected get %v, but got %v", tc.expectedResult, result)
			}
		})
	}
}

func TestStack_Size(t *testing.T) {
	type testcase struct {
		test           string
		size           int
		push           int
		expectedResult int
	}

	testcases := []testcase{
		{
			test:           "Empty stack",
			size:           0,
			push:           0,
			expectedResult: 0,
		},
		{
			test:           "Stack filled element",
			size:           1,
			push:           1,
			expectedResult: 1,
		},
		{
			test:           "Stack with element but not filled",
			size:           2,
			push:           1,
			expectedResult: 1,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			s, err := stack.New(tc.size)
			if err != nil {
				t.Fatal(err)
			}
			for i := 0; i < tc.push; i++ {
				err := s.Push(i)
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

func TestStack_Push(t *testing.T) {
	type testcase struct {
		test          string
		size          int
		push          int
		expectedError error
	}

	testcases := []testcase{
		{
			test:          "Empty stack",
			size:          0,
			push:          1,
			expectedError: stack.ErrOverflow,
		},
		{
			test:          "Fill Stack",
			size:          1,
			push:          1,
			expectedError: nil,
		},
		{
			test:          "Fill Stack too much",
			size:          2,
			push:          4,
			expectedError: stack.ErrOverflow,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			s, err := stack.New(tc.size)
			if err != nil {
				t.Fatal(err)
			}
			for i := 0; i < tc.push; i++ {
				err = s.Push(i)
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

func TestStack_Pop(t *testing.T) {
	type testcase struct {
		test          string
		size          int
		push          int
		pop           int
		expectedError error
	}

	testcases := []testcase{
		{
			test:          "Empty stack",
			size:          0,
			push:          0,
			pop:           1,
			expectedError: stack.ErrUnderflow,
		},
		{
			test:          "Filled Stack",
			size:          1,
			push:          1,
			pop:           1,
			expectedError: nil,
		},
		{
			test:          "Pop Stack too much",
			size:          2,
			push:          1,
			pop:           2,
			expectedError: stack.ErrUnderflow,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			s, err := stack.New(tc.size)
			if err != nil {
				t.Fatal(err)
			}
			for i := 0; i < tc.push; i++ {
				err = s.Push(i)
				if err != nil {
					t.Fatal(err)
				}
			}
			for i := 0; i < tc.pop; i++ {
				_, err = s.Pop()
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

func TestStack_Peek(t *testing.T) {
	type testcase struct {
		test          string
		size          int
		push          int
		expectedError error
	}

	testcases := []testcase{
		{
			test:          "Empty Stack",
			size:          0,
			push:          0,
			expectedError: stack.ErrUnderflow,
		},
		{
			test:          "Empty Stack",
			size:          1,
			push:          0,
			expectedError: stack.ErrUnderflow,
		},
		{
			test:          "Fill Stack",
			size:          1,
			push:          1,
			expectedError: nil,
		},
		{
			test:          "Larger Fill Stack",
			size:          2,
			push:          2,
			expectedError: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			s, err := stack.New(tc.size)
			if err != nil {
				t.Fatal(err)
			}
			for i := 0; i < tc.push; i++ {
				err := s.Push(i)
				if err != nil {
					t.Fatal(err)
				}
			}
			temp, err := s.Peek()
			if err != tc.expectedError {
				t.Errorf("Expected get %v, but got %v", tc.expectedError, err)
			}
			temp2, err := s.Peek()
			if err != tc.expectedError {
				t.Errorf("Expected get %v, but got %v", tc.expectedError, err)
			}
			if temp != temp2 {
				t.Errorf("Expected get %v and %v, but got %v and %v", temp, temp, temp, temp2)
			}
		})
	}
}
