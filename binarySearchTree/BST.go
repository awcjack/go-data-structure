package bst

import (
	"errors"

	"golang.org/x/exp/constraints"
)

var (
	ErrNodeNotFound = errors.New("node not found in tree")
)

type node[T constraints.Ordered] struct {
	key   T
	value interface{}
	left  *node[T]
	right *node[T]
}

type BST[T constraints.Ordered] struct {
	root *node[T]
}

func New[T constraints.Ordered]() *BST[T] {
	return &BST[T]{
		root: nil,
	}
}

func (n node[T]) Value() interface{} {
	return n.value
}

func (t *BST[T]) Clear() {
	t.root = nil
}

func (t BST[T]) IsEmpty() bool {
	return t.root == nil
}

func (n *node[T]) Insert(key T, value interface{}) {
	if key < n.key {
		if n.left == nil {
			n.left = &node[T]{
				key:   key,
				value: value,
			}
		} else {
			n.left.Insert(key, value)
		}
	} else {
		if n.right == nil {
			n.right = &node[T]{
				key:   key,
				value: value,
			}
		} else {
			n.right.Insert(key, value)
		}
	}
}

func (t *BST[T]) Insert(key T, value interface{}) {
	if t.root == nil {
		t.root = &node[T]{
			key:   key,
			value: value,
		}
	} else {
		t.root.Insert(key, value)
	}
}

func (n node[T]) Contains(key T) bool {
	if key == n.key {
		return true
	}
	if key < n.key {
		if n.left == nil {
			return false
		} else {
			return n.left.Contains(key)
		}
	}
	if n.right == nil {
		return false
	} else {
		return n.right.Contains(key)
	}
}

func (t BST[T]) Contains(key T) bool {
	if t.root == nil {
		return false
	}
	return t.root.Contains(key)
}

func (n node[T]) Find(key T) (interface{}, error) {
	if key == n.key {
		return n.value, nil
	}
	if key < n.key {
		if n.left == nil {
			return nil, ErrNodeNotFound
		} else {
			return n.left.Find(key)
		}
	}
	if n.right == nil {
		return nil, ErrNodeNotFound
	} else {
		return n.right.Find(key)
	}
}

func (t BST[T]) Find(key T) (interface{}, error) {
	if t.root == nil {
		return nil, ErrNodeNotFound
	}
	return t.root.Find(key)
}

func (n node[T]) FindMin() *node[T] {
	if n.left == nil {
		return &n
	} else {
		return n.left.FindMin()
	}
}

func (t BST[T]) FindMin() *node[T] {
	if t.root == nil {
		return nil
	}

	return t.root.FindMin()
}

func (n node[T]) FindMax() *node[T] {
	if n.right == nil {
		return &n
	} else {
		return n.right.FindMax()
	}
}

func (t BST[T]) FindMax() *node[T] {
	if t.root == nil {
		return nil
	}

	return t.root.FindMax()
}

func (n *node[T]) FindMinWithParent() (parent, min *node[T]) {
	if n.left == nil {
		return nil, n
	}

	if n.left.left == nil {
		return n, n.left
	}

	return n.left.FindMinWithParent()
}

func (n node[T]) DeleteFromTree() *node[T] {
	if n.left == nil && n.right == nil {
		return nil
	}

	if n.left == nil {
		return n.right
	}

	if n.right == nil {
		return n.left
	}

	parent, min := n.right.FindMinWithParent()

	if parent == nil {
		min.left = n.left
		return min
	}

	parent.left = nil
	min.left = n.left
	min.right = n.right
	return min
}

func (n *node[T]) Delete(key T) error {
	if key > n.key {
		if n.right == nil {
			return ErrNodeNotFound
		}
		if n.right.value == key {
			newSubTree := n.right.DeleteFromTree()
			n.right = newSubTree
			return nil
		}
		return n.right.Delete(key)
	} else {
		if n.left == nil {
			return ErrNodeNotFound
		}
		if n.left.value == key {
			newSubTree := n.left.DeleteFromTree()
			n.left = newSubTree
			return nil
		}
		return n.left.Delete(key)
	}
}

func (t *BST[T]) Delete(key T) error {
	if t.root == nil {
		return ErrNodeNotFound
	}
	if t.root.key == key {
		t.root = t.root.DeleteFromTree()
		return nil
	}
	return t.root.Delete(key)
	// if t.root.key == key {
	// 	if t.root.left == nil && t.root.right == nil {
	// 		t.root = nil
	// 		return nil
	// 	}
	// 	if t.root.right == nil {
	// 		// only t.root.left exist
	// 		t.root = t.root.left
	// 		return nil
	// 	}
	// 	if t.root.left == nil {
	// 		t.root = t.root.right
	// 		return nil
	// 	}

	// 	// both left and right exist
	// 	// find minimum from right child and move to root
	// 	node := t.root.right.FindMin()

	// 	if node.right == nil {
	// 		left := t.root.left
	// 		right := t.root.right
	// 		t.root = node
	// 		t.root.left = left
	// 		t.root.right = right
	// 	} else {
	// 		maxChildNode := node.FindMax()
	// 		maxChildNode.right = t.root.right
	// 		left := t.root.left
	// 		t.root = node
	// 		t.root.left = left
	// 	}
	// }

	// return nil
}

func (n node[T]) PreOrder() []interface{} {
	var result []interface{}

	result = append(result, n.value)

	if n.left != nil {
		result = append(result, n.left.PreOrder()...)
	}

	if n.right != nil {
		result = append(result, n.right.PreOrder()...)
	}

	return result
}

func (t BST[T]) PreOrder() []interface{} {
	if t.root == nil {
		return []interface{}{}
	}

	return t.root.PreOrder()
}

func (n node[T]) InOrder() []interface{} {
	var result []interface{}

	if n.left != nil {
		result = append(result, n.left.InOrder()...)
	}

	result = append(result, n.value)

	if n.right != nil {
		result = append(result, n.right.InOrder()...)
	}

	return result
}

func (t BST[T]) InOrder() []interface{} {
	if t.root == nil {
		return []interface{}{}
	}

	return t.root.InOrder()
}

func (n node[T]) PostOrder() []interface{} {
	var result []interface{}

	if n.left != nil {
		result = append(result, n.left.PostOrder()...)
	}

	if n.right != nil {
		result = append(result, n.right.PostOrder()...)
	}

	result = append(result, n.value)

	return result
}

func (t BST[T]) PostOrder() []interface{} {
	if t.root == nil {
		return []interface{}{}
	}

	return t.root.PostOrder()
}
