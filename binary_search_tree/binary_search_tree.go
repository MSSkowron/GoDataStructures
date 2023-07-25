package bst

import "golang.org/x/exp/constraints"

// Node is a generic binary search tree node
type Node[T constraints.Ordered] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

// NewNode creates a new binary search tree node
func NewNode[T constraints.Ordered](value T) *Node[T] {
	return &Node[T]{Value: value}
}

// Insert inserts a new node into the binary search tree
func (n *Node[T]) Insert(value T) *Node[T] {
	if n == nil {
		return NewNode[T](value)
	}

	if n.Value == value {
		return n
	}

	if value > n.Value {
		if n.Right == nil {
			n.Right = NewNode[T](value)
			return n
		}

		n.Right.Insert(value)
		return n
	}

	if n.Left == nil {
		n.Left = NewNode[T](value)
		return n
	}

	n.Left.Insert(value)
	return n
}

// Delete deletes a value from the binary search tree
func (n *Node[T]) Delete(value T) *Node[T] {
	if n == nil {
		return nil
	}

	if value < n.Value {
		n.Left = n.Left.Delete(value)
	} else if value > n.Value {
		n.Right = n.Right.Delete(value)
	} else {
		// Case 1: Node to be deleted has one or no children
		if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		}

		// Case 2: Node to be deleted has two children
		// Find the minimum value in the right subtree (successor)
		minRight := n.Right.findMinimum()
		// Replace the current node's value with the successor value
		n.Value = minRight.Value
		// Delete the successor node from the right subtree
		n.Right = n.Right.Delete(minRight.Value)
	}

	return n
}

func (n *Node[T]) findMinimum() *Node[T] {
	node := n
	for node.Left != nil {
		node = node.Left
	}
	return node
}

// Search searches a value in the binary search tree
func (n *Node[T]) Search(value T) *Node[T] {
	if n == nil {
		return nil
	}

	if n.Value == value {
		return n
	}

	if value > n.Value {
		return n.Right.Search(value)
	}

	return n.Left.Search(value)
}
