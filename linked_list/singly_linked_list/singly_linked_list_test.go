package singlylinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	ll := New[int]()
	assert.NotNil(t, ll)
	assert.Nil(t, ll.head)
	assert.Equal(t, 0, ll.length)
}

func TestPrepend(t *testing.T) {
	ll := New[int]()

	// Test Prepend 1
	ll.Prepend(1)
	assertLinkedListEqual(t, ll, []int{1})

	// Test Prepend 2
	ll.Prepend(2)
	assertLinkedListEqual(t, ll, []int{2, 1})

	// Test Prepend 3
	ll.Prepend(3)
	assertLinkedListEqual(t, ll, []int{3, 2, 1})
}

func TestAppend(t *testing.T) {
	ll := New[int]()

	// Test Append 1
	ll.Append(1)
	assertLinkedListEqual(t, ll, []int{1})

	// Test Append 2
	ll.Append(2)
	assertLinkedListEqual(t, ll, []int{1, 2})

	// Test Append 3
	ll.Append(3)
	assertLinkedListEqual(t, ll, []int{1, 2, 3})
}

func TestDelete(t *testing.T) {
	ll := buildSinglyLinkedList()

	// Test deleting a value in the middle of the list
	ll.Delete(3)
	assertLinkedListEqual(t, ll, []int{1, 2, 4, 5})

	// Test deleting the first value
	ll.Delete(1)
	assertLinkedListEqual(t, ll, []int{2, 4, 5})

	// Test deleting the last value
	ll.Delete(5)
	assertLinkedListEqual(t, ll, []int{2, 4})

	// Test deleting a non-existing value
	ll.Delete(10)
	assertLinkedListEqual(t, ll, []int{2, 4})
}

func TestContains(t *testing.T) {
	ll := buildSinglyLinkedList()

	for i := 1; i <= 5; i++ {
		assert.True(t, ll.Contains(i))
	}

	assert.False(t, ll.Contains(10))
}

func TestLength(t *testing.T) {
	ll := buildSinglyLinkedList()

	assert.Equal(t, 5, ll.Length())

	for i := 5; i >= 1; i-- {
		ll.Delete(i)
		assert.Equal(t, i-1, ll.Length())
	}
}

func assertLinkedListEqual(t *testing.T, ll *LinkedList[int], expected []int) {
	assert.Equal(t, len(expected), ll.length)

	curr := ll.head
	for _, val := range expected {
		assert.NotNil(t, curr)
		assert.Equal(t, val, curr.value)
		curr = curr.next
	}
}

func buildSinglyLinkedList() *LinkedList[int] {
	ll := New[int]()
	for i := 1; i <= 5; i++ {
		ll.Append(i)
	}

	return ll
}
