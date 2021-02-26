package src

import (
	"bytes"
	"encoding/json"
)

type TodoList struct {

	todos map[string]*Todo
}

func NewTodoList(todos []*Todo) *TodoList {

	var NewTodoList TodoList

	NewTodoList.todos = make(map[string]*Todo)

	for _, td := range todos {

		NewTodoList.todos[td.Id] = td

	}

	return &NewTodoList

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

	t.todos[todo.Id] = todo

}

func (t *TodoList) ReadTodo(id string) []byte {
	buf := &bytes.Buffer{}

	json.NewEncoder(buf).Encode(t.todos[id])

	return buf.Bytes()
}

func (t *TodoList) ReadAllTodos() []byte {

	buf := &bytes.Buffer{}

	json.NewEncoder(buf).Encode(t.todos)

	return buf.Bytes()

}

func (t *TodoList) UpdateTodo(newStatus *StatusChange) {

	t.todos[newStatus.Id].Status = newStatus.Status

}

func (t *TodoList) DeleteTodo(id string) {

	delete(t.todos, id)

}
