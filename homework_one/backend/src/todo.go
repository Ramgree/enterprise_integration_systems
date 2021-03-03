package src

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

type TodoList struct {
	Todos map[string]*Todo
	GlobalCounter int
}

func NewTodoListAndDag(todos []*Todo, edges []*Edge) (*DAG, *TodoList) {

	var NewTodoList TodoList

	NewTodoList.GlobalCounter = 0

	var NewDAG DAG

	NewTodoList.Todos = make(map[string]*Todo)

	for _, td := range todos {

		NewTodoList.CreateTodo(td)

		//NewTodoList.Todos[td.Id] = td

	}

	NewDAG.nodes = &NewTodoList
	NewDAG.AdjacencyList = make(map[string]map[string]bool)

	for index, val := range NewDAG.nodes.Todos {

		NewDAG.AdjacencyList[index] = make(map[string]bool)

		fmt.Println("Index: ", index, val)

	}

	for _, edge := range edges {

		if (NewDAG.AdjacencyList[edge.From])[edge.To] != true {

			fmt.Println(edge.From, edge.To)

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

	newId := &t.GlobalCounter

	todo.Id = strconv.Itoa(*newId)

	t.Todos[todo.Id] = todo

	*newId = *newId + 1

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
