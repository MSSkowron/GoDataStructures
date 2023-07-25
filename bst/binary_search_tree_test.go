package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNode(t *testing.T) {
	n := NewNode[int](10)
	assert.NotNil(t, n, "bst node is not empty")
	assert.Equal(t, n.Value, 10)
}

func TestInsert(t *testing.T) {
	n := buildIntTree()

	assert.NotNil(t, n.Right)
	assert.Equal(t, 20, n.Right.Value)

	assert.NotNil(t, n.Left)
	assert.Equal(t, 0, n.Left.Value)

	assert.NotNil(t, n.Left.Right)
	assert.Equal(t, 5, n.Left.Right.Value)

	assert.NotNil(t, n.Right.Left)
	assert.Equal(t, 15, n.Right.Left.Value)

	assert.NotNil(t, n.Right.Right)
	assert.Equal(t, 30, n.Right.Right.Value)
}

func TestDeleteLeaf(t *testing.T) {
	n := buildIntTree()

	n = n.Delete(5)
	assert.NotNil(t, n)
	assert.Equal(t, 10, n.Value)
	assert.Equal(t, 0, n.Left.Value)
	assert.Nil(t, n.Left.Right)
	assert.Equal(t, 20, n.Right.Value)
	assert.Equal(t, 30, n.Right.Right.Value)
	assert.Equal(t, 15, n.Right.Left.Value)
}

func TestDeleteRoot(t *testing.T) {
	n := buildIntTree()

	n = n.Delete(10)
	assert.NotNil(t, n)
	assert.Equal(t, 15, n.Value)
	assert.Equal(t, 0, n.Left.Value)
	assert.Equal(t, 5, n.Left.Right.Value)
	assert.Equal(t, 20, n.Right.Value)
	assert.Equal(t, 30, n.Right.Right.Value)
	assert.Nil(t, n.Right.Left)
}

func TestDeleteNodeWithOneChild(t *testing.T) {
	n := buildIntTree()

	n = n.Delete(0)
	assert.NotNil(t, n)
	assert.Equal(t, 10, n.Value)
	assert.Equal(t, 5, n.Left.Value)
	assert.Nil(t, n.Left.Right)
	assert.Nil(t, n.Left.Left)
	assert.Equal(t, 20, n.Right.Value)
	assert.Equal(t, 30, n.Right.Right.Value)
	assert.Equal(t, 15, n.Right.Left.Value)
}

func TestDeleteNotExistingNode(t *testing.T) {
	n := buildIntTree()

	n = n.Delete(100)
	assert.Equal(t, 20, n.Right.Value)
	assert.Equal(t, 0, n.Left.Value)
	assert.Equal(t, 5, n.Left.Right.Value)
	assert.Equal(t, 15, n.Right.Left.Value)
	assert.Equal(t, 30, n.Right.Right.Value)
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
	assert.Equal(t, n.Left.Right, foundNode)

	foundNode = n.Search(15)
	assert.NotNil(t, foundNode)
	assert.Equal(t, n.Right.Left, foundNode)

	foundNode = n.Search(30)
	assert.NotNil(t, foundNode)
	assert.Equal(t, n.Right.Right, foundNode)
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
