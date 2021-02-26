package src

type Edge struct {
	from string //id
	to string
}

type DAG struct {
	nodes TodoList
	edges []*Edge
	adjacencyList map[string][]string
}