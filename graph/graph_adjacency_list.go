package graph

// AdjList is a graph represented by an adjacency list.
type AdjList struct {
	vertices [][]int
}

// NewAdjList creates a new graph represented by an adjacency list.
func NewAdjList() *AdjList {
	return &AdjList{
		vertices: make([][]int, 0),
	}
}

// AddVertex adds a vertex to the graph.
// Returns the index of the new vertex.
func (g *AdjList) AddVertex() int {
	n := len(g.vertices)

	g.vertices = append(g.vertices, make([]int, 0))

	return n
}

// AddEdge adds an edge to the graph.
func (g *AdjList) AddEdge(from, to int) error {
	n := len(g.vertices)

	if from < 0 || from >= n {
		return ErrIndexOutOfRange
	}

	if to < 0 || to >= n {
		return ErrIndexOutOfRange
	}

	addToFrom := true
	for _, v := range g.vertices[from] {
		if v == to {
			addToFrom = false
			break
		}
	}
	if addToFrom {
		g.vertices[from] = append(g.vertices[from], to)
	}

	return nil
}

// RemoveEdge removes an edge from the graph.
func (g *AdjList) RemoveEdge(from, to int) error {
	n := len(g.vertices)

	if from < 0 || from >= n {
		return ErrIndexOutOfRange
	}

	if to < 0 || to >= n {
		return ErrIndexOutOfRange
	}

	for i, v := range g.vertices[from] {
		if v == to {
			g.vertices[from] = append(g.vertices[from][:i], g.vertices[from][i+1:]...)
			break
		}
	}

	return nil
}
