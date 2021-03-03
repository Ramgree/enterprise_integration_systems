package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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
	req, err := http.NewRequest("DELETE", baseURL+"/todo/"+id, nil)

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

func AddTodo(title string) string {
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

	var id_s = ""

	if err != nil {
		log.Println("an error occurred")
	} else {

		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		} else {
			id_s = string(bodyBytes)
			id_s = id_s[1 : len(id_s)-2]
			log.Println("Successfully added a TODO. It should be visible in your list now. New ID:", id_s)
			return id_s

		}
	}

	return id_s
}

func AddDependency(id string, dependency string) {
	_, err := http.Post(baseURL+"/todo/"+dependency+"/"+id, "application/json",
		nil)

	if err != nil {
		log.Println("an error occurred")
	}
}

func AddDependencies(id string, dependencies []string) {
	log.Print(dependencies)
	for _, dep := range dependencies {
		AddDependency(id, dep)
	}
}

func Check(id string, action string) {
	var jsonString string = `{ "Id":"` + id + `", "Status":"` + action + `"}`

	log.Print(jsonString)
	// var processedString = []bytes(jsonString)
	_, err := http.Post(baseURL+"/todo/"+id, "application/json",
		bytes.NewBuffer([]byte(jsonString)))

	if err != nil {
		log.Println("Could not perform this action. ( Check that depending tasks are all checked! )")
	}

}
