package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAdjList(t *testing.T) {
	n := NewAdjList()

	assert.NotNil(t, n)
	assert.Equal(t, 0, len(n.vertices))
}

func TestAdjList_AddVertex(t *testing.T) {
	n := NewAdjList()

	assert.Equal(t, 0, n.AddVertex())
	assert.Equal(t, 1, n.AddVertex())
	assert.Equal(t, 2, n.AddVertex())
	assert.Equal(t, 3, n.AddVertex())
	assert.Equal(t, 4, n.AddVertex())

	assert.Equal(t, 5, len(n.vertices))

	for i := 0; i < 5; i++ {
		assert.Equal(t, 0, len(n.vertices[i]))
	}
}

func TestAdjList_AddEdge(t *testing.T) {
	n := NewAdjList()

	assert.Equal(t, ErrIndexOutOfRange, n.AddEdge(0, 1))

	for i := 0; i < 5; i++ {
		n.AddVertex()
	}

	assert.Equal(t, ErrIndexOutOfRange, n.AddEdge(0, 5))

	assert.Nil(t, n.AddEdge(0, 1))
	assert.Equal(t, 1, len(n.vertices[0]))
	assert.Equal(t, 1, n.vertices[0][0])

	assert.Nil(t, n.AddEdge(0, 2))
	assert.Equal(t, 2, len(n.vertices[0]))
	assert.Equal(t, 2, n.vertices[0][1])
}

func TestAdjList_RemoveEdge(t *testing.T) {
	n := NewAdjList()

	assert.Equal(t, ErrIndexOutOfRange, n.RemoveEdge(0, 1))

	for i := 0; i < 5; i++ {
		n.AddVertex()
	}

	assert.Equal(t, ErrIndexOutOfRange, n.RemoveEdge(0, 5))

	assert.Nil(t, n.AddEdge(0, 1))
	assert.Equal(t, 1, len(n.vertices[0]))
	assert.Equal(t, 1, n.vertices[0][0])

	assert.Nil(t, n.AddEdge(0, 2))
	assert.Equal(t, 2, len(n.vertices[0]))
	assert.Equal(t, 2, n.vertices[0][1])

	assert.Nil(t, n.RemoveEdge(0, 1))
	assert.Equal(t, 1, len(n.vertices[0]))
	assert.Equal(t, 2, n.vertices[0][0])

	assert.Nil(t, n.RemoveEdge(0, 2))
	assert.Equal(t, 0, len(n.vertices[0]))
}
