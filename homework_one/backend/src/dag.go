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
	adjacencyList map[string][]string
}

func (d *DAG) AddEdge(edge *Edge) {

	d.Edges = append(d.Edges, edge)

}

func (d *DAG) GetEdges() []byte {

	buf := &bytes.Buffer{}

	json.NewEncoder(buf).Encode(d.Edges)

	return buf.Bytes()


}