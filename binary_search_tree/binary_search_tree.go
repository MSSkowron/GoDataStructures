package tree

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
func (n *Node[T]) Insert(value T) {
	if n == nil {
		return
	}

	if n.Value == value {
		return
	}

	if value > n.Value {
		if n.Right == nil {
			n.Right = NewNode[T](value)
			return
		}

		n.Right.Insert(value)
		return
	}

	if n.Left == nil {
		n.Left = NewNode[T](value)
		return
	}

	n.Left.Insert(value)
}

// Delete deletes a value from the binary search tree
func (n *Node[T]) Delete(value T) {

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
