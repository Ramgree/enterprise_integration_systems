package src

import (
	"bytes"
	"encoding/json"
)

type TodoList struct {
	Todos map[string]*Todo
}

func NewTodoListAndDag(todos []*Todo, edges []*Edge) (*DAG, *TodoList) {

	var NewTodoList TodoList

	var NewDAG DAG

	NewTodoList.Todos = make(map[string]*Todo)

	for _, td := range todos {

		NewTodoList.Todos[td.Id] = td

	}

	NewDAG.nodes = &NewTodoList
	NewDAG.AdjacencyList = make(map[string]map[string]bool)

	for _, td := range todos {

		NewDAG.AdjacencyList[td.Id] = make(map[string]bool)

	}

	for _, edge := range edges {

		if (NewDAG.AdjacencyList[edge.From])[edge.To] != true {

			NewDAG.Edges = append(NewDAG.Edges, edge)
			(NewDAG.AdjacencyList[edge.From])[edge.To] = true
		}
	}

	return &NewDAG, &NewTodoList

}

type Todo struct {
	Id     string
	Title  string
	Status string
}

type StatusChange struct {
	Id     string
	Status string
}

func (t *TodoList) CreateTodo(todo *Todo) {

	t.Todos[todo.Id] = todo

}

func (t *TodoList) ReadTodo(id string) []byte {
	buf := &bytes.Buffer{}

	json.NewEncoder(buf).Encode(t.Todos[id])

	return buf.Bytes()
}

func (t *TodoList) ReadAllTodos() []byte {

	buf := &bytes.Buffer{}

	json.NewEncoder(buf).Encode(t.Todos)

	return buf.Bytes()

}

func (t *TodoList) UpdateTodo(newStatus *StatusChange) {

	t.Todos[newStatus.Id].Status = newStatus.Status

}

func (t *TodoList) DeleteTodo(id string) {

	delete(t.Todos, id)

}
