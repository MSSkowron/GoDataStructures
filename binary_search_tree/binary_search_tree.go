package binarysearchtree

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Node represents a binary search tree node of type T.
type Node[T constraints.Ordered] struct {
	value T
	left  *Node[T]
	right *Node[T]
}

// NewNode creates a new binary search tree node with the given value of type T.
func NewNode[T constraints.Ordered](value T) *Node[T] {
	return &Node[T]{value: value}
}

// Insert inserts a new node into the binary search tree.
func (n *Node[T]) Insert(value T) *Node[T] {
	if n == nil {
		return NewNode[T](value)
	}

	if n.value == value {
		return n
	}

	if value > n.value {
		if n.right == nil {
			n.right = NewNode[T](value)
			return n
		}

		n.right.Insert(value)
		return n
	}

	if n.left == nil {
		n.left = NewNode[T](value)
		return n
	}

	n.left.Insert(value)
	return n
}

// Delete deletes a value from the binary search tree.
func (n *Node[T]) Delete(value T) *Node[T] {
	if n == nil {
		return nil
	}

	if value < n.value {
		n.left = n.left.Delete(value)
	} else if value > n.value {
		n.right = n.right.Delete(value)
	} else {
		// Case 1: Node to be deleted has one or no children
		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		}

		// Case 2: Node to be deleted has two children
		// Find the minimum value in the right subtree (successor)
		minright := n.right.findMinimum()
		// Replace the current node's value with the successor value
		n.value = minright.value
		// Delete the successor node from the right subtree
		n.right = n.right.Delete(minright.value)
	}

	return n
}

func (n *Node[T]) findMinimum() *Node[T] {
	node := n
	for node.left != nil {
		node = node.left
	}
	return node
}

// Search searches a value in the binary search tree.
func (n *Node[T]) Search(value T) *Node[T] {
	if n == nil {
		return nil
	}

	if n.value == value {
		return n
	}

	if value > n.value {
		return n.right.Search(value)
	}

	return n.left.Search(value)
}

// InorderTraversal traverses the binary search tree in inorder fashion.
func (n *Node[T]) InorderTraversal() {
	if n == nil {
		return
	}

	n.left.InorderTraversal()
	fmt.Printf("%v ", n.value)
	n.right.InorderTraversal()
}

// PreorderTraversal traverses the binary search tree in preorder fashion.
func (n *Node[T]) PreorderTraversal() {
	if n == nil {
		return
	}

	fmt.Printf("%v ", n.value)
	n.left.InorderTraversal()
	n.right.InorderTraversal()
}

// PostorderTraversal traverses the binary search tree in postorder fashion.
func (n *Node[T]) PostorderTraversal() {
	if n == nil {
		return
	}

	n.left.InorderTraversal()
	n.right.InorderTraversal()
	fmt.Printf("%v ", n.value)
}
