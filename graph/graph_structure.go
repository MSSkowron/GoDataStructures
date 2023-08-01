package graph

import (
	"errors"
	"fmt"
)

var (
	// ErrVertexExists is returned when a vertex is added to a graph that already exists.
	ErrVertexAlreadyExists = errors.New("vertex with this key already exists")
	// ErrVertexDoesNotExist is returned when a vertex is referenced that does not exist.
	ErrVertexDoesNotExist = errors.New("vertex with this key does not exist")
)

// Key represents a unique identifier for a vertex in a graph.
type Key int

// Graph represents a set of vertices connected by edges.
type Graph[T any] struct {
	Vertices map[Key]*Vertex[T]
}

// Vertex represents a vertex in a graph.
type Vertex[T any] struct {
	Value T
	Edges map[Key]*Edge[T]
}

// Edge represents an edge in the graph and the destination vertex.
type Edge[T any] struct {
	Weight int
	Vertex *Vertex[T]
}

// GraphOption is a function that modifies a graph.
type GraphOption[T any] func(*Graph[T])

// NewGraph creates a new empty graph.
func NewGraph[T any](opts ...GraphOption[T]) *Graph[T] {
	g := &Graph[T]{
		Vertices: make(map[Key]*Vertex[T]),
	}

	for _, opt := range opts {
		opt(g)
	}

	return g
}

// WithVertices sets the vertices of the graph.
func WithVertices[T any](vertices map[Key]*Vertex[T]) GraphOption[T] {
	return func(g *Graph[T]) {
		g.Vertices = vertices
	}
}

// WithAdjacencyList sets the vertices of the graph using an adjacency list.
func WithAdjacencyList[T any](list map[int][]int) GraphOption[T] {
	return func(g *Graph[T]) {
		for k, v := range list {
			g.Vertices[Key(k)] = &Vertex[T]{
				Edges: make(map[Key]*Edge[T]),
			}

			for _, e := range v {
				g.Vertices[Key(k)].Edges[Key(e)] = &Edge[T]{
					Weight: 0,
					Vertex: g.Vertices[Key(e)],
				}
			}
		}
	}
}

// AddVertex adds a vertex to the graph.
// If a vertex with the same key already exists, ErrVertexAlreadyExists is returned.
func (g *Graph[T]) AddVertex(k Key, v T) error {
	if v := g.GetVertex(k); v != nil {
		return ErrVertexAlreadyExists
	}

	g.Vertices[k] = &Vertex[T]{
		Value: v,
		Edges: make(map[Key]*Edge[T]),
	}

	return nil
}

// RemoveVertex removes a vertex from the graph.
// If the vertex does not exist, ErrVertexDoesNotExist is returned.
func (g *Graph[T]) RemoveVertex(k Key) error {
	if v := g.GetVertex(k); v == nil {
		return ErrVertexDoesNotExist
	}

	delete(g.Vertices, k)

	for _, v := range g.Vertices {
		delete(v.Edges, k)
	}

	return nil
}

// AddEdge adds an edge between two vertices in the graph.
// If either of the vertices does not exist, ErrVertexDoesNotExist is returned.
func (g *Graph[T]) AddEdge(src, dest Key, weight int) error {
	v1, v2 := g.GetVertex(src), g.GetVertex(dest)
	if v1 == nil || v2 == nil {
		return ErrVertexDoesNotExist
	}

	v1.Edges[dest] = &Edge[T]{
		Weight: weight,
		Vertex: v2,
	}

	return nil
}

// RemoveEdge removes an edge between two vertices in the graph.
// If either of the vertices does not exist, ErrVertexDoesNotExist is returned.
func (g *Graph[T]) RemoveEdge(src, dest Key) error {
	v1, v2 := g.GetVertex(src), g.GetVertex(dest)
	if v1 == nil || v2 == nil {
		return ErrVertexDoesNotExist
	}

	delete(v1.Edges, dest)

	return nil
}

// GetVertex returns the vertex with the given key.
// If the vertex does not exist, nil is returned.
func (g *Graph[T]) GetVertex(k Key) *Vertex[T] {
	return g.Vertices[k]
}

// GetEdge returns the edge between two vertices.
// If either of the vertices does not exist, ErrVertexDoesNotExist is returned.
// If the edge does not exist, nil is returned.
func (g *Graph[T]) GetEdge(src, dest Key) (*Edge[T], error) {
	v1, v2 := g.GetVertex(src), g.GetVertex(dest)
	if v1 == nil || v2 == nil {
		return nil, ErrVertexDoesNotExist
	}

	return v1.Edges[dest], nil
}

// Print prints the graph.
func (g *Graph[T]) Print() {
	for k, v := range g.Vertices {
		fmt.Println(k, v.Value)
		for _, e := range v.Edges {
			fmt.Println("\t", e.Vertex.Value)
		}
	}
}
