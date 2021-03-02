package service

import (
	"encoding/json"
	"log"
	"net/http"
	"todocli/model"
)

const baseURL = "http://localhost:8000"

//GetAllTodos is fetching all todos from the backend API
func GetAllTodos() model.TodoList {
	resp, err := http.Get(baseURL + "/todo")

	if err != nil {
		log.Println("an error occurred")
	 }

	 defer resp.Body.Close()

	 var data model.TodoList

	 if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Println("an error occurred")
	 }

	 return data
  	
}

//GetAllEdges is fetching all edges from the backend API
func GetAllEdges() model.EdgeList {
	resp, err := http.Get(baseURL + "/edges")

	if err != nil {
		log.Println("an error occurred")
	 }

	 defer resp.Body.Close()

	 var data model.EdgeList

	 if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Println("an error occurred")
	 }

	 return data
  	
}