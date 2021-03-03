package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"todo_backend/src"

	"github.com/gorilla/mux"
)

var todos = []*src.Todo{
	{
		Title:  "Turn cluster on",
		Status: "Unfinished",
	},
	{
		Title:  "Process data",
		Status: "Finished",
	},
	{
		Title:  "Turn another cluster on",
		Status: "Finished",
	},
	{
		Title:  "Pay the server bills",
		Status: "Finished",
	},
	{
		Title:  "Process more data",
		Status: "Unfinished",
	},
	{
		Title:  "Shut down cluster",
		Status: "Unfinished",
	},
}

var edges = []*src.Edge{
	{
		From: "3",
		To:   "2",
	},
	{
		From: "4",
		To:   "5",
	},
	{
		From: "0",
		To:   "3",
	},
}

var dependencyDAG, globalState = src.NewTodoListAndDag(todos, edges)

func PostCreateTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("creating new todo")

	reqBody, _ := ioutil.ReadAll(r.Body)
	todo := src.Todo{}
	json.Unmarshal(reqBody, &todo)
	globalState.CreateTodo(&todo)

	log.Println("Let's see 2", todo)

	log.Println(globalState.Todos)

	for _, address := range globalState.Todos {

		log.Println(*address)

	}

	if len(dependencyDAG.AdjacencyList[todo.Id]) == 0 {

		dependencyDAG.AdjacencyList[todo.Id] = make(map[string]bool)

	}

	json.NewEncoder(w).Encode(todo.Id)

	w.WriteHeader(200)

}

func PostCreateEdge(w http.ResponseWriter, r *http.Request) {
	log.Println("adding new edge")

	vars := mux.Vars(r)
	from := vars["from"]
	to := vars["to"]

	from_todo, to_todo := globalState.Todos[from], globalState.Todos[to]

	if from_todo == nil || to_todo == nil {

		w.WriteHeader(400)
		return

	} else {

		log.Println("checking the request body for adding new edge: ", from, to)
		edge := src.Edge{}
		edge.From = from
		edge.To = to

		dependencyDAG.AddEdge(&edge)

		json.NewEncoder(w).Encode(edge)
		w.WriteHeader(200)

	}

}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("getting todos")
	vars := mux.Vars(r)
	key := vars["id"]

	key_value := globalState.Todos[key]

	if key_value == nil {

		w.WriteHeader(400)
		return

	} else {

		w.Write(globalState.ReadTodo(key))
		w.WriteHeader(200)

	}

}

func GetAllTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("getting all todos")
	w.Write(globalState.ReadAllTodos())
	w.WriteHeader(200)
	log.Println(globalState.Todos)

}

func GetAllEdges(w http.ResponseWriter, r *http.Request) {
	log.Println("getting all edges")
	log.Println(dependencyDAG.Edges)
	log.Println(dependencyDAG.AdjacencyList)

	for _, address := range dependencyDAG.Edges {

		log.Println(*address)

	}
	w.Write(dependencyDAG.GetEdges())
	w.WriteHeader(200)
}

func GetDownstream(w http.ResponseWriter, r *http.Request) {
	log.Println("fetching all dependencies")
	vars := mux.Vars(r)
	key := vars["id"]

	key_value := globalState.Todos[key]

	if key_value == nil {

		w.WriteHeader(400)
		return

	} else {

		w.Write(dependencyDAG.DepthFirstSearch(key))
		w.WriteHeader(200)

	}
}

func PostUpdateTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("updating an existing todo")

	reqBody, _ := ioutil.ReadAll(r.Body)
	statusChange := src.StatusChange{}
	json.Unmarshal(reqBody, &statusChange)

	key_value := globalState.Todos[statusChange.Id]

	if key_value == nil {

		w.WriteHeader(400)
		return

	} else {
		log.Println(statusChange)
		globalState.UpdateTodo(&statusChange)
		json.NewEncoder(w).Encode(statusChange)
		w.WriteHeader(200)
	}

}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("deleting todos")
	vars := mux.Vars(r)
	key := vars["id"]

	key_value := globalState.Todos[key]

	if key_value == nil {

		w.WriteHeader(400)
		return

	} else {

		globalState.DeleteTodo(key)
		w.WriteHeader(200)
	}

}

func main() {
	router := mux.NewRouter()

	log.Println("Server started")

	// Nodes
	router.HandleFunc("/todo", PostCreateTodo).Methods(http.MethodPost)
	router.HandleFunc("/todo", GetAllTodo).Methods(http.MethodGet)
	router.HandleFunc("/todo/{id}", GetTodo).Methods(http.MethodGet)
	router.HandleFunc("/todo/{id}", PostUpdateTodo).Methods(http.MethodPost)
	router.HandleFunc("/todo/{id}", DeleteTodo).Methods(http.MethodDelete)

	// Edges
	router.HandleFunc("/todo/{from}/{to}", PostCreateEdge).Methods(http.MethodPost)
	router.HandleFunc("/edges", GetAllEdges).Methods(http.MethodGet)
	router.HandleFunc("/dependencies/{id}", GetDownstream).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8000", router))

}
