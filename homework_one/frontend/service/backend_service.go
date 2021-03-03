package service

import (
	"bytes"
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

//DeleteTodo sends a delete request to the backend API
func DeleteTodo(id string) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", baseURL+"/todo/"+id, nil)

	if err != nil {
		log.Println("an error occurred")
	}

	_, err = client.Do(req)
	if err != nil {
		log.Println("an error occurred")
	} else {
		log.Println("successfully deleted a TODO with ID", id)
	}

}

func AddTodo(title string) int {
	// client := &http.Client{}

	// I'll assume the id is generated on the backend, cuz its kinda bad if we do that on front-end

	entity := model.Todo{
		Title:  title,
		Status: "Unfinished", // Not sure if I should leave it like this
	}

	var jsonData []byte
	jsonData, err := json.Marshal(entity)

	resp, err := http.Post(baseURL+"/todo", "application/json",
		bytes.NewBuffer(jsonData))

	if err != nil {
		log.Println("an error occurred")
	} else {
		log.Println("Successfully added a TODO. It should be visible in your list now. Code:" + resp.Status)
	}

	defer resp.Body.Close()
	// log.Println(json.NewDecoder(resp.Body).Decode())
	return -1
}
