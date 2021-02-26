package src

import (
	"bytes"
	"encoding/json"
)

type Edge struct {
	From string //id
	To string
}

type DAG struct {
	nodes *TodoList
	Edges []*Edge
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