package graph

// AdjMatrix is a graph represented by an adjacency matrix.
type AdjMatrix struct {
	vertices [][]bool
}

// NewAdjMatrix creates a new graph represented by an adjacency matrix.
func NewAdjMatrix() *AdjMatrix {
	return &AdjMatrix{
		vertices: make([][]bool, 0),
	}
}

// AddVertex adds a vertex to the graph.
// Returns the index of the new vertex.
func (g *AdjMatrix) AddVertex() int {
	n := len(g.vertices)
	newRow := make([]bool, n+1)
	g.vertices = append(g.vertices, newRow)
	for i := 0; i < n; i++ {
		g.vertices[i] = append(g.vertices[i], false)
	}

	return n
}

// AddEdge adds an edge to the graph.
func (g *AdjMatrix) AddEdge(from, to int) error {
	n := len(g.vertices)

	if from < 0 || from >= n {
		return ErrIndexOutOfRange
	}

	if to < 0 || to >= n {
		return ErrIndexOutOfRange
	}

	g.vertices[from][to] = true

	return nil
}

// RemoveEdge removes an edge from the graph.
func (g *AdjMatrix) RemoveEdge(from, to int) error {
	n := len(g.vertices)

	if from < 0 || from >= n {
		return ErrIndexOutOfRange
	}

	if to < 0 || to >= n {
		return ErrIndexOutOfRange
	}

	g.vertices[from][to] = false

	return nil
}
