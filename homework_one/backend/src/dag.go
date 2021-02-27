package src

import (
	"bytes"
	"encoding/json"
)

type Edge struct {
	From string
	To   string
}

type DAG struct {
	nodes         *TodoList
	Edges         []*Edge
	AdjacencyList map[string]map[string]bool
}

func (d *DAG) AddEdge(edge *Edge) {

	if (d.AdjacencyList[edge.From])[edge.To] != true {

		d.Edges = append(d.Edges, edge)
		(d.AdjacencyList[edge.From])[edge.To] = true

	}
}

func (d *DAG) GetEdges() []byte {

	buf := &bytes.Buffer{}

	json.NewEncoder(buf).Encode(d.Edges)

	return buf.Bytes()

}

func (d *DAG) DepthFirstSearchIteration(todo string, visited *map[string]bool) {

	(*visited)[todo] = true

	for neighbor_id, _ := range d.AdjacencyList[todo] {

		if (*visited)[neighbor_id] == false {

			d.DepthFirstSearchIteration(neighbor_id, visited)

		}
	}

}

func (d *DAG) DepthFirstSearch(todo string) []byte {

	visited := make(map[string]bool)

	d.DepthFirstSearchIteration(todo, &visited)

	delete(visited, todo)

	downstream_nodes := make([]*Todo, 0)

	for todo_id, _ := range visited {

		downstream_nodes = append(downstream_nodes, d.nodes.Todos[todo_id])

	}

	buf := &bytes.Buffer{}

	json.NewEncoder(buf).Encode(downstream_nodes)

	return buf.Bytes()

}
