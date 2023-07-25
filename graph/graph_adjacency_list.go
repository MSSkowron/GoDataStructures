package graph

// GraphAdjList is a graph represented by an adjacency list.
type GraphAdjList struct {
	vertices [][]int
}

// NewGraphAdjList creates a new graph represented by an adjacency list.
func NewGraphAdjList() *GraphAdjList {
	return &GraphAdjList{
		vertices: make([][]int, 0),
	}
}

// Vertices adds a vertex to the graph.
// Returns the index of the new vertex.
func (g *GraphAdjList) AddVertex() int {
	n := len(g.vertices)

	g.vertices = append(g.vertices, make([]int, 0))

	return n
}

// AddEdge adds an edge to the graph.
func (g *GraphAdjList) AddEdge(from, to int) error {
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
