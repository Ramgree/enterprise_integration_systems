package src

import (
	"bytes"
	"encoding/json"
)

type TodoList struct {
	todos []*Todo
}

func NewTodoList(todos []*Todo) *TodoList {

	return &TodoList{

		todos: todos,
	}

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

	t.todos = append(t.todos, todo)

}

func (t *TodoList) ReadTodo(id string) []byte {
	buf := &bytes.Buffer{}
	for _, todo := range t.todos {

		if todo.Id == id {

			json.NewEncoder(buf).Encode(todo)

		}
	}
	return buf.Bytes()
}

func (t *TodoList) ReadAllTodos() []byte {

	buf := &bytes.Buffer{}
	json.NewEncoder(buf).Encode(t.todos)

	return buf.Bytes()

}

func (t *TodoList) UpdateTodo(newStatus *StatusChange) {

	for _, todo := range t.todos {

		if todo.Id == newStatus.Id {

			todo.Status = newStatus.Status

		}

	}

}

func (t *TodoList) DeleteTodo(id string) {

	for index, todo := range t.todos {

		if todo.Id == id {

			t.todos = append(t.todos[:index], t.todos[index+1:]...)

		}

	}

}
