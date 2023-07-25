package graph

import "errors"

var (
	ErrIndexOutOfRange = errors.New("index out of range")
)

type GraphAList struct {
	vertices [][]int
}

func NewGraphAList() *GraphAList {
	return &GraphAList{
		vertices: make([][]int, 0),
	}
}

func (g *GraphAList) AddVertex() {
	g.vertices = append(g.vertices, make([]int, 0))
}

func (g *GraphAList) AddEdge(from, to int) error {
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
