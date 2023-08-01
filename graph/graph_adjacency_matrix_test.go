package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAdjMatrix(t *testing.T) {
	n := NewAdjMatrix()

	assert.NotNil(t, n)
	assert.Equal(t, 0, len(n.vertices))
}

func TestAdjMatrix_AddVertex(t *testing.T) {
	n := NewAdjMatrix()

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

func TestAdjMatrix_AddEdge(t *testing.T) {
	n := NewAdjMatrix()

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

func TestAdjMatrix_RemoveEdge(t *testing.T) {
	n := NewAdjMatrix()

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
