package graph

// GraphAdjMatrix is a graph represented by an adjacency matrix.
type GraphAdjMatrix struct {
	vertices [][]bool
}

// NewGraphAdjMatrix creates a new graph represented by an adjacency matrix.
func NewGraphAdjMatrix() *GraphAdjMatrix {
	return &GraphAdjMatrix{
		vertices: make([][]bool, 0),
	}
}

// Vertices adds a vertex to the graph.
// Returns the index of the new vertex.
func (g *GraphAdjMatrix) AddVertex() int {
	n := len(g.vertices)
	newRow := make([]bool, n+1)
	g.vertices = append(g.vertices, newRow)
	for i := 0; i < n; i++ {
		g.vertices[i] = append(g.vertices[i], false)
	}

	return n
}

// AddEdge adds an edge to the graph.
func (g *GraphAdjMatrix) AddEdge(from, to int) error {
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
