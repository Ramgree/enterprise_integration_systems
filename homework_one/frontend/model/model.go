package model

type Todo struct {
	Id     *string
	Title  string
	Status string
}

type Edge struct {
	From string
	To   string
}

type TodoList map[string]*Todo

type EdgeList []*Edge
