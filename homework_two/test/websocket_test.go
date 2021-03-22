package test

import (
	"encoding/json"
	"fmt"
	"rentit/pkg/domain"
	"testing"

	"github.com/gorilla/websocket"
)

const (
	wsPort = 8081
)

func TestGetAllWs(t *testing.T) {

	expectedCount := 8

	c, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://localhost:%d/ws", wsPort), nil)

	if err != nil {
		t.Error("Failed to connect to WebSocket: " + err.Error())
		return 
	}

	defer c.Close()

	var data []*domain.Plant

	err = c.WriteJSON(domain.SocketQuery{Command: "all"})

	
	if err != nil {
		t.Error("Failed to send a message: " + err.Error())
		return 
	}

	_, msg, err := c.ReadMessage()

	if err != nil {
		t.Error("Failed to receive a message back: " + err.Error())
		return 
	}

	if err := json.Unmarshal(msg, &data); err != nil {
		t.Error("Couldn't decode the response: " + err.Error())
		return 
	}

	if len(data) != expectedCount{
		t.Error(fmt.Sprintf("Expected %d results, got %d", expectedCount, len(data)))
		return 
	}

	for _, plant := range data {
		if plant == nil{
			t.Error("One plant was nil")
		}
	}
}


func TestEstimatePriceWs(t *testing.T) {
	cmd := "estimate"
	name := "bulldozer"
	fakeName := "Bruno Rucy"
	sd := "2020-01-01"
	ed := "2020-01-10"

	verifyPriceWs(t, domain.SocketQuery{Command: cmd, Name: &name, StartDate: &sd, EndDate: &ed}, 45000)
	verifyError(t, domain.SocketQuery{Command: cmd, Name: nil, StartDate: &sd, EndDate: &ed})
	verifyError(t, domain.SocketQuery{Command: cmd, Name: &fakeName, StartDate: &sd, EndDate: &ed})
	verifyError(t, domain.SocketQuery{Command: cmd, Name: &name, StartDate: nil, EndDate: &ed})
}

func TestAvailabilityWs(t *testing.T) {
	cmd := "availability"
	name := "crane"
	fakeName := "Bruno Rucy"
	sd := "2021-11-18"
	ed := "2021-11-20"

	verifyAvailabilityWs(t, domain.SocketQuery{Command: cmd, Name: &name, StartDate: &sd, EndDate: &ed}, false)
	verifyError(t, domain.SocketQuery{Command: cmd, Name: nil, StartDate: &sd, EndDate: &ed})
	verifyError(t, domain.SocketQuery{Command: cmd, Name: &fakeName, StartDate: &sd, EndDate: &ed})
	verifyError(t, domain.SocketQuery{Command: cmd, Name: &name, StartDate: nil, EndDate: &ed})
}


func verifyPriceWs(t *testing.T, query domain.SocketQuery, expected float32){
	c, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://localhost:%d/ws", wsPort), nil)

	if err != nil {
		t.Error("Failed to connect to WebSocket: " + err.Error())
		return 
	}

	defer c.Close()

	err = c.WriteJSON(query)


	if err != nil {
		t.Error("Failed to send a message: " + err.Error())
		return 
	}

	_, msg, err := c.ReadMessage()

	if err != nil {
		t.Error("Failed to receive a message back: " + err.Error())
		return 
	}

	var data map[string]float32

	if err := json.Unmarshal(msg, &data); err != nil {
		t.Error("Couldn't decode the response: " + err.Error())
		return 
	}

	for key, _ := range data {
		if key != "price"{
			t.Error("Invalid field in the response")
			return 
		}
	}

	if _, ok := data["price"]; ok {
		if !ok{
			t.Error("\"price\" field not present in response")
			return 
		}
	}

	if data["price"] != expected{
		t.Error("Wrong price returned")
		return 
	}
}

func verifyAvailabilityWs(t *testing.T, query domain.SocketQuery, expected bool){
	c, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://localhost:%d/ws", wsPort), nil)

	if err != nil {
		t.Error("Failed to connect to WebSocket: " + err.Error())
		return 
	}

	defer c.Close()

	err = c.WriteJSON(query)


	if err != nil {
		t.Error("Failed to send a message: " + err.Error())
		return 
	}

	_, msg, err := c.ReadMessage()

	if err != nil {
		t.Error("Failed to receive a message back: " + err.Error())
		return 
	}

	var data map[string]bool

	if err := json.Unmarshal(msg, &data); err != nil {
		t.Error("Couldn't decode the response: " + err.Error())
		return 
	}

	for key, _ := range data {
		if key != "isAvailable"{
			t.Error("Invalid field in the response")
			return 
		}
	}

	if _, ok := data["isAvailable"]; ok {
		if !ok{
			t.Error("\"isAvailable\" field not present in response")
			return 
		}
	}

	if data["isAvailable"] != expected{
		t.Error("Wrong price returned")
		return 
	}
}

func verifyError(t *testing.T, query domain.SocketQuery){
	c, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://localhost:%d/ws", wsPort), nil)

	if err != nil {
		t.Error("Failed to connect to WebSocket: " + err.Error())
		return 
	}

	defer c.Close()

	err = c.WriteJSON(query)


	if err != nil {
		t.Error("Failed to send a message: " + err.Error())
		return 
	}

	_, msg, err := c.ReadMessage()

	if err != nil {
		t.Error("Failed to receive a message back: " + err.Error())
		return 
	}

	var data map[string]string

	if err := json.Unmarshal(msg, &data); err != nil {
		t.Error("Couldn't decode the response: " + err.Error())
		return 
	}

	for key, _ := range data {
		if key != "error"{
			t.Error("Invalid field in the response")
			return 
		}
	}

	if _, ok := data["error"]; ok {
		if !ok{
			t.Error("\"error\" field not present in response")
		}
	}
}