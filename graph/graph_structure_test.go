package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGraph(t *testing.T) {
	n := NewGraph[int]()

	assert.NotNil(t, n)
	assert.NotNil(t, n.Vertices)
	assert.Equal(t, 0, len(n.Vertices))
}

func TestAddVertex(t *testing.T) {
	g := NewGraph[int]()
	key := Key(1)
	value := 10

	// Test adding a new vertex
	err := g.AddVertex(key, value)
	assert.NoError(t, err)

	v, ok := g.Vertices[key]
	assert.True(t, ok)
	assert.Equal(t, value, v.Value)

	// Test adding a vertex with the same key (should return ErrVertexAlreadyExists)
	err = g.AddVertex(key, value)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrVertexAlreadyExists)
}

func TestRemoveVertex(t *testing.T) {
	g := NewGraph[int]()
	key := Key(1)
	value := 10
	g.AddVertex(key, value)

	// Test removing an existing vertex
	err := g.RemoveVertex(key)
	assert.NoError(t, err)

	v, ok := g.Vertices[key]
	assert.False(t, ok)
	assert.Nil(t, v)

	// Test removing a non-existing vertex (should return ErrVertexDoesNotExist)
	err = g.RemoveVertex(key)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrVertexDoesNotExist)
}

func TestAddEdge(t *testing.T) {
	g := NewGraph[int]()
	srcKey := Key(1)
	destKey := Key(2)
	weight := 5
	g.AddVertex(srcKey, 10)
	g.AddVertex(destKey, 20)

	// Test adding an edge between existing vertices
	err := g.AddEdge(srcKey, destKey, weight)
	assert.NoError(t, err)

	v1, ok := g.Vertices[srcKey]
	assert.True(t, ok)
	assert.Equal(t, weight, v1.Edges[destKey].Weight)

	v2, ok := g.Vertices[destKey]
	assert.True(t, ok)
	assert.Equal(t, 0, len(v2.Edges))

	// Test adding existing edge (should return nil)
	err = g.AddEdge(srcKey, destKey, weight)
	assert.NoError(t, err)

	// Test adding an edge with a non-existing source vertex (should return ErrVertexDoesNotExist)
	err = g.AddEdge(Key(3), destKey, weight)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrVertexDoesNotExist)

	// Test adding an edge with a non-existing destination vertex (should return ErrVertexDoesNotExist)
	err = g.AddEdge(srcKey, Key(4), weight)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrVertexDoesNotExist)
}

func TestRemoveEdge(t *testing.T) {
	g := NewGraph[int]()
	srcKey := Key(1)
	destKey := Key(2)
	weight := 5
	g.AddVertex(srcKey, 10)
	g.AddVertex(destKey, 20)
	g.AddEdge(srcKey, destKey, weight)

	// Test removing an existing edge
	err := g.RemoveEdge(srcKey, destKey)
	assert.NoError(t, err)

	v1, ok := g.Vertices[srcKey]
	assert.True(t, ok)
	assert.Equal(t, 0, len(v1.Edges))

	// Test removing a non-existing edge
	err = g.RemoveEdge(srcKey, destKey)
	assert.NoError(t, err)

	// Test removing an edge with a non-existing source vertex (should return ErrVertexDoesNotExist)
	err = g.RemoveEdge(Key(3), destKey)
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrVertexDoesNotExist)

	// Test removing an edge with a non-existing destination vertex (should return ErrVertexDoesNotExist)
	err = g.RemoveEdge(srcKey, Key(4))
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrVertexDoesNotExist)
}

func TestGetVertex(t *testing.T) {
	g := NewGraph[int]()
	key := Key(1)
	value := 10
	g.AddVertex(key, value)

	// Test getting an existing vertex
	vertex := g.GetVertex(key)
	assert.NotNil(t, vertex)
	assert.Equal(t, value, vertex.Value)

	// Test getting a non-existing vertex
	vertex = g.GetVertex(Key(2))
	assert.Nil(t, vertex)
}

func TestGetEdge(t *testing.T) {
	g := NewGraph[int]()
	srcKey := Key(1)
	destKey := Key(2)
	weight := 5
	g.AddVertex(srcKey, 10)
	g.AddVertex(destKey, 20)
	g.AddEdge(srcKey, destKey, weight)

	// Test getting an existing edge
	edge, err := g.GetEdge(srcKey, destKey)
	assert.NoError(t, err)
	assert.NotNil(t, edge)
	assert.Equal(t, weight, edge.Weight)

	// Test getting a non-existing edge
	edge, err = g.GetEdge(srcKey, Key(3))
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrVertexDoesNotExist)
	assert.Nil(t, edge)
}
