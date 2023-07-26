package binarysearchtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNode(t *testing.T) {
	n := NewNode[int](10)
	assert.NotNil(t, n, "bst node is not empty")
	assert.Equal(t, n.value, 10)
}

func TestInsert(t *testing.T) {
	n := buildIntTree()

	assert.NotNil(t, n.right)
	assert.Equal(t, 20, n.right.value)

	assert.NotNil(t, n.left)
	assert.Equal(t, 0, n.left.value)

	assert.NotNil(t, n.left.right)
	assert.Equal(t, 5, n.left.right.value)

	assert.NotNil(t, n.right.left)
	assert.Equal(t, 15, n.right.left.value)

	assert.NotNil(t, n.right.right)
	assert.Equal(t, 30, n.right.right.value)
}

func TestDeleteLeaf(t *testing.T) {
	n := buildIntTree()

	n = n.Delete(5)
	assert.NotNil(t, n)
	assert.Equal(t, 10, n.value)
	assert.Equal(t, 0, n.left.value)
	assert.Nil(t, n.left.right)
	assert.Equal(t, 20, n.right.value)
	assert.Equal(t, 30, n.right.right.value)
	assert.Equal(t, 15, n.right.left.value)
}

func TestDeleteRoot(t *testing.T) {
	n := buildIntTree()

	n = n.Delete(10)
	assert.NotNil(t, n)
	assert.Equal(t, 15, n.value)
	assert.Equal(t, 0, n.left.value)
	assert.Equal(t, 5, n.left.right.value)
	assert.Equal(t, 20, n.right.value)
	assert.Equal(t, 30, n.right.right.value)
	assert.Nil(t, n.right.left)
}

func TestDeleteNodeWithOneChild(t *testing.T) {
	n := buildIntTree()

	n = n.Delete(0)
	assert.NotNil(t, n)
	assert.Equal(t, 10, n.value)
	assert.Equal(t, 5, n.left.value)
	assert.Nil(t, n.left.right)
	assert.Nil(t, n.left.left)
	assert.Equal(t, 20, n.right.value)
	assert.Equal(t, 30, n.right.right.value)
	assert.Equal(t, 15, n.right.left.value)
}

func TestDeleteNotExistingNode(t *testing.T) {
	n := buildIntTree()

	n = n.Delete(100)
	assert.Equal(t, 20, n.right.value)
	assert.Equal(t, 0, n.left.value)
	assert.Equal(t, 5, n.left.right.value)
	assert.Equal(t, 15, n.right.left.value)
	assert.Equal(t, 30, n.right.right.value)
}

func TestDeleteNilNode(t *testing.T) {
	var n *Node[int] = nil

	n = n.Delete(100)
	assert.Nil(t, n)
}

func TestSearch(t *testing.T) {
	n := buildIntTree()

	foundNode := n.Search(10)
	assert.NotNil(t, foundNode)
	assert.Equal(t, n, foundNode)

	foundNode = n.Search(5)
	assert.NotNil(t, foundNode)
	assert.Equal(t, n.left.right, foundNode)

	foundNode = n.Search(15)
	assert.NotNil(t, foundNode)
	assert.Equal(t, n.right.left, foundNode)

	foundNode = n.Search(30)
	assert.NotNil(t, foundNode)
	assert.Equal(t, n.right.right, foundNode)
}

func buildIntTree() *Node[int] {
	var n *Node[int] = nil

	n = n.Insert(10)
	n = n.Insert(20)
	n = n.Insert(0)
	n = n.Insert(5)
	n = n.Insert(15)
	n = n.Insert(30)

	return n
}
