package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGraphAdjMatrix(t *testing.T) {
	n := NewGraphAdjMatrix()

	assert.NotNil(t, n)
	assert.Equal(t, 0, len(n.vertices))
}

func TestGraphAdjMatrix_AddVertex(t *testing.T) {
	n := NewGraphAdjMatrix()

	assert.Equal(t, 0, n.AddVertex())
	assert.Equal(t, 1, n.AddVertex())
	assert.Equal(t, 2, n.AddVertex())
	assert.Equal(t, 3, n.AddVertex())
	assert.Equal(t, 4, n.AddVertex())

	assert.Equal(t, 5, len(n.vertices))

	for i := 0; i < 5; i++ {
		assert.Equal(t, 5, len(n.vertices[i]))
	}
}

func TestGraphAdjMatrix_AddEdge(t *testing.T) {
	n := NewGraphAdjMatrix()

	assert.Equal(t, ErrIndexOutOfRange, n.AddEdge(0, 1))

	for i := 0; i < 5; i++ {
		n.AddVertex()
	}

	assert.Equal(t, ErrIndexOutOfRange, n.AddEdge(0, 5))

	assert.Nil(t, n.AddEdge(0, 1))
	assert.True(t, n.vertices[0][1])
	assert.False(t, n.vertices[1][0])

	assert.Nil(t, n.AddEdge(0, 2))
	assert.True(t, n.vertices[0][2])
	assert.False(t, n.vertices[2][0])

	assert.False(t, n.vertices[1][2])
}

func TestGraphAdjMatrix_RemoveEdge(t *testing.T) {
	n := NewGraphAdjMatrix()

	assert.Equal(t, ErrIndexOutOfRange, n.RemoveEdge(0, 1))

	for i := 0; i < 5; i++ {
		n.AddVertex()
	}

	assert.Equal(t, ErrIndexOutOfRange, n.RemoveEdge(0, 5))

	assert.Nil(t, n.AddEdge(0, 1))
	assert.Nil(t, n.AddEdge(0, 2))

	assert.Nil(t, n.RemoveEdge(0, 1))
	assert.False(t, n.vertices[0][1])
	assert.False(t, n.vertices[1][0])

	assert.Nil(t, n.RemoveEdge(0, 2))
	assert.False(t, n.vertices[0][2])
	assert.False(t, n.vertices[2][0])
}
