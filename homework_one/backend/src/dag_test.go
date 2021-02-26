package src

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestNewDAG(t *testing.T) {

	test_map := make(map[string][]string, 0)

	id_one := "1"

	example_adjacency_list := make([]string, 0, 5)

	example_adjacency_list = append(example_adjacency_list, "arara")

	test_map[id_one] = example_adjacency_list

}

func TestAddNewEdge(t *testing.T) {

	// Adding new edge

	address := "http://localhost:8000/todo/1/2"

	_, err := http.Post(address, "application/json", nil)

	if err != nil {

		t.Error("AAAAAAAAAAAAAAAAAAAAa")

	}

}

func TestReadAllEdges(t *testing.T) {

	var resp_struct []*Edge

	resp, err := http.Get("http://localhost:8000/edges")

	//defer resp.Body.Close()

	if err != nil {

		t.Error("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")

	}

	err = json.NewDecoder(resp.Body).Decode(&resp_struct)

	for _, address := range resp_struct {

		fmt.Println(*address)

	}

	if len(resp_struct) < 1 {

		t.Error("struct is empty :*(")

	}

}
