package doublylinkedlist

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type node[T constraints.Ordered] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

// LinkedList represents a doubly linked list with nodes with values of type T.
type LinkedList[T constraints.Ordered] struct {
	head   *node[T]
	length int
}

// New creates a new, empty doubly linked list with nodes with values of type T.
func New[T constraints.Ordered]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Prepend adds a new element at the beginning of the list.
func (ll *LinkedList[T]) Prepend(value T) {
	ll.length++

	n := &node[T]{value: value}
	n.next = ll.head
	if ll.head != nil {
		ll.head.prev = n
	}

	ll.head = n
}

// Append adds a new element at the end of the list.
func (ll *LinkedList[T]) Append(value T) {
	ll.length++

	if ll.head == nil {
		ll.head = &node[T]{value: value}
		return
	}

	curr := ll.head
	for curr.next != nil {
		curr = curr.next
	}

	n := &node[T]{value: value}
	curr.next = n
	n.prev = curr
}

// Delete deletes the first occurrence of the given value from the list.
func (ll *LinkedList[T]) Delete(value T) {
	if ll.head == nil {
		return
	}

	if ll.head.value == value {
		ll.head = ll.head.next
		ll.length--

		if ll.head != nil {
			ll.head.prev = nil
		}

		return
	}

	curr := ll.head
	for curr.next != nil {
		if curr.next.value == value {
			curr.next = curr.next.next
			if curr.next != nil {
				curr.next.prev = curr
			}
			ll.length--
			return
		}

		curr = curr.next
	}
}

// Contains returns true if the list contains the given value, false otherwise.
func (ll *LinkedList[T]) Contains(value T) bool {
	curr := ll.head
	for curr != nil {
		if curr.value == value {
			return true
		}

		curr = curr.next
	}

	return false
}

// Length returns the length of the list.
func (ll *LinkedList[T]) Length() int {
	return ll.length
}

// Print prints the list.
func (ll *LinkedList[T]) Print() {
	if ll.head == nil {
		fmt.Println("nil")
		return
	}

	curr := ll.head
	for curr != nil {
		if curr.next == nil {
			fmt.Printf("%v -> ", curr.value)
			break
		}

		fmt.Printf("%v <-> ", curr.value)
		curr = curr.next
	}
	fmt.Printf("nil")
}
