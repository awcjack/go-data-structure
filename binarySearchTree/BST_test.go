package bst_test

import (
	"testing"

	bst "github.com/awcjack/go-data-structure/binarySearchTree"
)

func TestBST_IsEmpty(t *testing.T) {
	type testcase struct {
		test           string
		input          []int
		expectedResult bool
	}

	testcases := []testcase{
		{
			test:           "Empty BST",
			input:          []int{},
			expectedResult: true,
		},
		{
			test:           "One element BST",
			input:          []int{0},
			expectedResult: false,
		},
		{
			test:           "BST with child",
			input:          []int{1, 0, 2},
			expectedResult: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			tree := bst.New[int]()
			for _, v := range tc.input {
				tree.Insert(v, v)
			}
			result := tree.IsEmpty()
			if tc.expectedResult != result {
				t.Errorf("expected get %v, but got %v", tc.expectedResult, result)
			}
		})
	}
}

func TestBST_Clear(t *testing.T) {
	type testcase struct {
		test  string
		input []int
	}

	testcases := []testcase{
		{
			test:  "Empty BST",
			input: []int{},
		},
		{
			test:  "One element BST",
			input: []int{0},
		},
		{
			test:  "BST with child",
			input: []int{1, 0, 2},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			tree := bst.New[int]()
			for _, v := range tc.input {
				tree.Insert(v, v)
			}
			tree.Clear()
			result := tree.IsEmpty()
			if result != true {
				t.Errorf("expected get %v, but got %v", true, result)
			}
		})
	}
}

func TestBST_Insert(t *testing.T) {
	type testcase struct {
		test           string
		input          []int
		expectedResult []int
	}

	testcases := []testcase{
		{
			test:           "Empty BST",
			input:          []int{},
			expectedResult: []int{},
		},
		{
			test:           "One element BST",
			input:          []int{0},
			expectedResult: []int{0},
		},
		{
			test:           "BST with child",
			input:          []int{1, 0, 2},
			expectedResult: []int{1, 0, 2},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			tree := bst.New[int]()
			for _, v := range tc.input {
				tree.Insert(v, v)
			}
			result := tree.PreOrder()
			for i, element := range result {
				if element.(int) != tc.expectedResult[i] {
					t.Errorf("expected get %v at index %v, but got %v", tc.expectedResult[i], i, element)
				}
			}
		})
	}
}

func TestBST_Contain(t *testing.T) {
	type testcase struct {
		test           string
		input          []int
		expectedResult []int
	}

	testcases := []testcase{
		{
			test:           "Empty BST",
			input:          []int{},
			expectedResult: []int{},
		},
		{
			test:  "One element BST",
			input: []int{0},
		},
		{
			test:  "BST with child",
			input: []int{1, 0, 2},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			tree := bst.New[int]()
			for _, v := range tc.input {
				tree.Insert(v, v)
			}
			for i, v := range tc.input {
				result := tree.Contains(v)
				if result != true {
					t.Errorf("expected get %v at index %v, but got %v", true, i, result)
				}
			}
			result := tree.Contains(10000000000000)
			if result != false {
				t.Errorf("expected get %v, but got %v", false, result)
			}

			result = tree.Contains(-10000000000000)
			if result != false {
				t.Errorf("expected get %v, but got %v", false, result)
			}
		})
	}
}

func TestBST_Find(t *testing.T) {
	type testcase struct {
		test           string
		input          []int
		expectedResult []int
	}

	testcases := []testcase{
		{
			test:           "Empty BST",
			input:          []int{},
			expectedResult: []int{},
		},
		{
			test:  "One element BST",
			input: []int{0},
		},
		{
			test:  "BST with child",
			input: []int{1, 0, 2},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			tree := bst.New[int]()
			for _, v := range tc.input {
				tree.Insert(v, v)
			}
			for i, v := range tc.input {
				result, err := tree.Find(v)
				if err != nil {
					t.Fatal(err)
				}
				if result != v {
					t.Errorf("expected get %v at index %v, but got %v", v, i, result)
				}
			}
			_, err := tree.Find(10000000000000)
			if err != bst.ErrNodeNotFound {
				t.Errorf("expected get %v, but got %v", bst.ErrNodeNotFound, err)
			}

			_, err = tree.Find(-10000000000000)
			if err != bst.ErrNodeNotFound {
				t.Errorf("expected get %v, but got %v", bst.ErrNodeNotFound, err)
			}
		})
	}
}

func TestBST_FindMin(t *testing.T) {
	type testcase struct {
		test                string
		input               []int
		expectedFoundResult bool
		expectedResult      int
	}

	testcases := []testcase{
		{
			test:                "Empty BST",
			input:               []int{},
			expectedFoundResult: false,
			expectedResult:      0,
		},
		{
			test:                "One element BST",
			input:               []int{0},
			expectedFoundResult: true,
			expectedResult:      0,
		},
		{
			test:                "BST with child",
			input:               []int{1, 0, 2},
			expectedFoundResult: true,
			expectedResult:      0,
		},
		{
			test:                "BST with child",
			input:               []int{1, 0, 2, -100},
			expectedFoundResult: true,
			expectedResult:      -100,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			tree := bst.New[int]()
			for _, v := range tc.input {
				tree.Insert(v, v)
			}
			result := tree.FindMin()

			if result == nil && tc.expectedFoundResult {
				t.Errorf("Expected found node with value %v but got nothing", tc.expectedResult)
			}

			if !tc.expectedFoundResult && result != nil {
				t.Errorf("Expected no node found but got %v", result)
			}

			if tc.expectedFoundResult && tc.expectedResult != (result.Value()).(int) {
				t.Errorf("Expected get %v, but got %v", tc.expectedResult, result.Value())
			}
		})
	}
}

func TestBST_FindMax(t *testing.T) {
	type testcase struct {
		test                string
		input               []int
		expectedFoundResult bool
		expectedResult      int
	}

	testcases := []testcase{
		{
			test:                "Empty BST",
			input:               []int{},
			expectedFoundResult: false,
			expectedResult:      0,
		},
		{
			test:                "One element BST",
			input:               []int{0},
			expectedFoundResult: true,
			expectedResult:      0,
		},
		{
			test:                "BST with child",
			input:               []int{1, 0, 2},
			expectedFoundResult: true,
			expectedResult:      2,
		},
		{
			test:                "BST with child",
			input:               []int{1, 0, 2, 100},
			expectedFoundResult: true,
			expectedResult:      100,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			tree := bst.New[int]()
			for _, v := range tc.input {
				tree.Insert(v, v)
			}
			result := tree.FindMax()

			if result == nil && tc.expectedFoundResult {
				t.Errorf("Expected found node with value %v but got nothing", tc.expectedResult)
			}

			if !tc.expectedFoundResult && result != nil {
				t.Errorf("Expected no node found but got %v", result)
			}

			if tc.expectedFoundResult && tc.expectedResult != (result.Value()).(int) {
				t.Errorf("Expected get %v, but got %v", tc.expectedResult, result.Value())
			}
		})
	}
}

func TestBST_Delete(t *testing.T) {
	type testcase struct {
		test           string
		input          []int
		delete         []int
		expectedResult []int // inorder result
		expectedError  error
	}

	testcases := []testcase{
		{
			test:           "Empty BST",
			input:          []int{},
			delete:         []int{0},
			expectedResult: []int{},
			expectedError:  bst.ErrNodeNotFound,
		},

		{
			test:           "Node not found",
			input:          []int{0},
			delete:         []int{1},
			expectedResult: []int{0},
			expectedError:  bst.ErrNodeNotFound,
		},

		{
			test:           "Node not found",
			input:          []int{0},
			delete:         []int{-1},
			expectedResult: []int{0},
			expectedError:  bst.ErrNodeNotFound,
		},
		{
			test:           "One element BST delete root",
			input:          []int{0},
			delete:         []int{0},
			expectedResult: []int{},
			expectedError:  nil,
		},
		{
			test:           "Two elements BST delete root",
			input:          []int{0, 1},
			delete:         []int{0},
			expectedResult: []int{1},
			expectedError:  nil,
		},
		{
			test:           "Two elements BST delete root",
			input:          []int{0, -1},
			delete:         []int{0},
			expectedResult: []int{-1},
			expectedError:  nil,
		},
		{
			test:           "Three element BST delete left",
			input:          []int{1, 0, 2},
			delete:         []int{1},
			expectedResult: []int{0, 2},
			expectedError:  nil,
		},
		{
			test:           "Three element BST delete right",
			input:          []int{1, 0, 2},
			delete:         []int{2},
			expectedResult: []int{0, 1},
			expectedError:  nil,
		},
		{
			test:           "Three element BST delete root",
			input:          []int{1, 0, 2},
			delete:         []int{0},
			expectedResult: []int{1, 2},
			expectedError:  nil,
		},
		{
			test:           "Large BST delete subtree head",
			input:          []int{1, 0, 3, 2, 4},
			delete:         []int{3},
			expectedResult: []int{0, 1, 2, 4},
			expectedError:  nil,
		},
		{
			test:           "Large BST delete subtree left",
			input:          []int{1, 0, 3, 2, 4},
			delete:         []int{2},
			expectedResult: []int{0, 1, 3, 4},
			expectedError:  nil,
		},
		{
			test:           "Large BST delete subtree right",
			input:          []int{1, 0, 3, 2, 4},
			delete:         []int{4},
			expectedResult: []int{0, 1, 2, 3},
			expectedError:  nil,
		},

		{
			test:           "Large BST delete subtree left",
			input:          []int{0, -2, -1, -3},
			delete:         []int{-3},
			expectedResult: []int{-2, -1, 0},
			expectedError:  nil,
		},
		{
			test:           "Large BST delete subtree right",
			input:          []int{0, -2, -1, -3},
			delete:         []int{-1},
			expectedResult: []int{-3, -2, 0},
			expectedError:  nil,
		},
		{
			test:           "Large BST delete root",
			input:          []int{1, 0, 3, 2, 4},
			delete:         []int{1},
			expectedResult: []int{0, 2, 3, 4},
			expectedError:  nil,
		},
		{
			test:           "Larger BST delete root",
			input:          []int{1, 0, 5, 4, 3, 2},
			delete:         []int{1},
			expectedResult: []int{0, 2, 3, 4, 5},
			expectedError:  nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			tree := bst.New[int]()
			for _, v := range tc.input {
				tree.Insert(v, v)
			}
			for _, v := range tc.delete {
				err := tree.Delete(v)
				if err != tc.expectedError {
					t.Errorf("Expected get %v, but received %v", tc.expectedError, err)
				}
			}
			if tc.expectedError == nil {
				result := tree.InOrder()
				for i, v := range result {
					if v != tc.expectedResult[i] {
						t.Errorf("Expected get %v at index %v, but received %v", tc.expectedResult[i], i, v)
					}
				}
			}
		})
	}
}

func TestBST_PreOrder(t *testing.T) {
	type testcase struct {
		test           string
		input          []int
		expectedResult []int
	}

	testcases := []testcase{
		{
			test:           "Empty BST",
			input:          []int{},
			expectedResult: []int{},
		},
		{
			test:           "One element BST",
			input:          []int{0},
			expectedResult: []int{0},
		},
		{
			test:           "BST with child",
			input:          []int{1, 0, 2},
			expectedResult: []int{1, 0, 2},
		},
		{
			test:           "BST with child",
			input:          []int{1, 0, 2, 100},
			expectedResult: []int{1, 0, 2, 100},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			tree := bst.New[int]()
			for _, v := range tc.input {
				tree.Insert(v, v)
			}
			result := tree.PreOrder()
			for i, v := range result {
				if v != tc.expectedResult[i] {
					t.Errorf("Expected get %v at index %v, but got %v", tc.expectedResult[i], i, v)
				}
			}
		})
	}
}

func TestBST_InOrder(t *testing.T) {
	type testcase struct {
		test           string
		input          []int
		expectedResult []int
	}

	testcases := []testcase{
		{
			test:           "Empty BST",
			input:          []int{},
			expectedResult: []int{},
		},
		{
			test:           "One element BST",
			input:          []int{0},
			expectedResult: []int{0},
		},
		{
			test:           "BST with child",
			input:          []int{1, 0, 2},
			expectedResult: []int{0, 1, 2},
		},
		{
			test:           "BST with child",
			input:          []int{1, 0, 2, 100},
			expectedResult: []int{0, 1, 2, 100},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			tree := bst.New[int]()
			for _, v := range tc.input {
				tree.Insert(v, v)
			}
			result := tree.InOrder()
			for i, v := range result {
				if v != tc.expectedResult[i] {
					t.Errorf("Expected get %v at index %v, but got %v", tc.expectedResult[i], i, v)
				}
			}
		})
	}
}

func TestBST_PostOrder(t *testing.T) {
	type testcase struct {
		test           string
		input          []int
		expectedResult []int
	}

	testcases := []testcase{
		{
			test:           "Empty BST",
			input:          []int{},
			expectedResult: []int{},
		},
		{
			test:           "One element BST",
			input:          []int{0},
			expectedResult: []int{0},
		},
		{
			test:           "BST with child",
			input:          []int{1, 0, 2},
			expectedResult: []int{0, 2, 1},
		},
		{
			test:           "BST with child",
			input:          []int{1, 0, 2, 100},
			expectedResult: []int{0, 100, 2, 1},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			tree := bst.New[int]()
			for _, v := range tc.input {
				tree.Insert(v, v)
			}
			result := tree.PostOrder()
			for i, v := range result {
				if v != tc.expectedResult[i] {
					t.Errorf("Expected get %v at index %v, but got %v", tc.expectedResult[i], i, v)
				}
			}
		})
	}
}
