package src

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestCreateReadTodo(t *testing.T) {

	// Create

	address := "http://localhost:8000/todo"

	todo := Todo{
		Id:     "2",
		Title:  "foobar",
		Status: "Finished",
	}

	marshalled_todo, err := json.Marshal(todo)

	if err != nil {

		t.Error("Marshalling failure // create", err)

	}

	resp, err := http.Post(address, "application/json", bytes.NewBuffer(marshalled_todo))

	if err != nil {

		t.Error("Problem with the endpoint // create")

	}

	defer resp.Body.Close()

	// Read

	address = "http://localhost:8000/todo/2"

	resp, err = http.Get(address)

	if err != nil {
		t.Error("Problem with the endpoint // read", err)
	}

	defer resp.Body.Close()

	var resp_struct Todo

	err = json.NewDecoder(resp.Body).Decode(&resp_struct)

	if resp_struct.Id != "2" {

		t.Error("Problem reading todo // read", err)

	}

}

func TestUpdate(t *testing.T) {

	// Updating

	address := "http://localhost:8000/todo/2"

	status_change := StatusChange{
		Id:     "2",
		Status: "Unfinished",
	}

	marshalled_status_change, err := json.Marshal(status_change)

	if err != nil {

		t.Error("Marshalling failure // update", err)

	}

	resp, err := http.Post(address, "application/json", bytes.NewBuffer(marshalled_status_change))

	if err != nil {

		t.Error("Problem with the update endpoint // update")

	}

	defer resp.Body.Close()

	// Confirming the update

	resp, err = http.Get(address)

	if err != nil {
		t.Error("Problem with the get endpoint // update", err)
	}

	defer resp.Body.Close()

	var resp_struct Todo

	err = json.NewDecoder(resp.Body).Decode(&resp_struct)

	if resp_struct.Status != "Unfinished" {

		t.Error("Todo was not updated", resp_struct)

	}

}

func TestDelete(t *testing.T) {

	// Deleting

	address := "http://localhost:8000/todo/2"

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodDelete, address, nil)

	if err != nil {

		t.Error("Problem with the delete endpoint // delete")

	}

	resp, err := client.Do(req)

	resp.Body.Close()

	// Confirming the deletion

	resp, err = http.Get(address)

	var resp_struct Todo

	err = json.NewDecoder(resp.Body).Decode(&resp_struct)

	empty_struct := Todo{}

	if resp_struct != empty_struct {

		t.Error("Delete failed", resp_struct)

	}

}
