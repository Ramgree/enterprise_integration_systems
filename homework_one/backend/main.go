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
		Title:  "Turn cluster on",
		Status: "Unfinished",
	},
	{
		Id:     "1",
		Title:  "Process data",
		Status: "Finished",
	},
	{
		Id:     "3",
		Title:  "Turn another cluster on",
		Status: "Finished",
	},
	{
		Id:     "4",
		Title:  "Pay the server bills",
		Status: "Finished",
	},
	{
		Id:     "6",
		Title:  "Process more data",
		Status: "Unfinished",
	},
	{
		Id:     "7",
		Title:  "Shut down cluster",
		Status: "Unfinished",
	},
}

var edges = []*src.Edge{
	{
		From: "4",
		To:   "3",
	},
	{
		From: "3",
		To:   "6",
	},
	{
		From: "6",
		To:   "7",
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

	json.NewEncoder(w).Encode(todo)

}

func PostCreateEdge(w http.ResponseWriter, r *http.Request) {
	log.Println("adding new edge")

	vars := mux.Vars(r)
	from := vars["from"]
	to := vars["to"]

	log.Println("checking the request body for adding new edge: ", from, to)
	edge := src.Edge{}
	edge.From = from
	edge.To = to

	dependencyDAG.AddEdge(&edge)

	json.NewEncoder(w).Encode(edge)

}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("getting todos")
	vars := mux.Vars(r)
	key := vars["id"]

	w.Write(globalState.ReadTodo(key))

}

func GetAllTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("getting all todos")
	w.Write(globalState.ReadAllTodos())
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
}

func GetDownstream(w http.ResponseWriter, r *http.Request) {
	log.Println("fetching all dependencies")
	vars := mux.Vars(r)
	key := vars["id"]

	//testib := dependencyDAG.DepthFirstSearch(key)

	//log.Println("Everything downstream of ", key, " : ", testib)

	w.Write(dependencyDAG.DepthFirstSearch(key))

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
