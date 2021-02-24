package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"todo_backend/src"
)

var todos = []*src.Todo{
	{
		Id:     "0",
		Title:  "foo",
		Status: "Unfinished",
	},
	{
		Id:     "1",
		Title:  "bar",
		Status: "Finished",
	},
}

var globalState = src.NewTodoList(todos)

func PostCreateTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("creating new todo")

	reqBody, _ := ioutil.ReadAll(r.Body)
	todo := src.Todo{}
	json.Unmarshal(reqBody, &todo)
	globalState.CreateTodo(&todo)

	json.NewEncoder(w).Encode(todo)

}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("getting todos")
	vars := mux.Vars(r)
	key := vars["id"]

	w.Write(globalState.ReadTodo(key))

}

func PostUpdateTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("updating an existing todo")

	reqBody, _ := ioutil.ReadAll(r.Body)
	statusChange := src.StatusChange{}
	json.Unmarshal(reqBody, &statusChange)

	log.Println(statusChange)
	globalState.UpdateTodo(&statusChange)

	json.NewEncoder(w).Encode(statusChange)

}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("deleting todos")
	vars := mux.Vars(r)
	key := vars["id"]

	globalState.DeleteTodo(key)

}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/todo", PostCreateTodo).Methods(http.MethodPost)
	router.HandleFunc("/todo/{id}", PostUpdateTodo).Methods(http.MethodPost)
	router.HandleFunc("/todo/{id}", GetTodo).Methods(http.MethodGet)
	router.HandleFunc("/todo/{id}", DeleteTodo).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8000", router))

}
