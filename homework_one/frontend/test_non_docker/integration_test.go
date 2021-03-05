package test_non_docker

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

const defaultCount = 6

var nextID = defaultCount

func TestDefaultTodos(t *testing.T) {
	cmd := exec.Command("../todocli", "get-all")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		t.Error("Failed to execute todocli get-all")
	}

	stdOut := out.String()

	for i := 0; i < defaultCount; i++ {
		if !strings.Contains(stdOut, fmt.Sprintf("Id: %d", i)) {
			t.Error("TODO with Id", i, " not found in default todos")
		}
	}
}

func TestCanAddTodo(t *testing.T) {
	title := "write some very interesting CLI tests"
	cmd := exec.Command("../todocli", "todo", title)

	var out bytes.Buffer
	// for some reason output goes to stderr here
	cmd.Stderr = &out

	err := cmd.Run()

	if err != nil {
		t.Error("Failed to execute todocli todo", title)
	}

	stdOut := out.String()

	if !strings.Contains(stdOut, fmt.Sprintf("Successfully added a TODO. It should be visible in your list now. New ID: %d", nextID)) {
		t.Error("todocli todo didn't print success message")
	}

	cmd = exec.Command("../todocli", "get-all")

	out.Reset()
	cmd.Stdout = &out

	err = cmd.Run()

	if err != nil {
		t.Error("Failed to execute todocli get-all")
	}

	stdOut = out.String()
	search := fmt.Sprintf(
		`
Id: %d
Name: %s
Status: Unfinished
TODO before this:  []
`,
		nextID, title)

	if !strings.Contains(stdOut, search) {
		t.Error("TODO with Id", nextID, "not found in all todos")
	}

	// assume that it works here
	exec.Command("../todocli", "remove", strconv.Itoa(nextID)).Run()

	// the next todo added will have a new ID
	nextID++
}

func TestCanDeleteTodo(t *testing.T) {
	// assume that this one works now
	exec.Command("../todocli", "todo", "test").Run()

	cmd := exec.Command("../todocli", "remove", strconv.Itoa(nextID))

	var out bytes.Buffer
	cmd.Stderr = &out

	err := cmd.Run()

	if err != nil {
		t.Error("Failed to execute todocli remove", nextID)
	}

	stdOut := out.String()

	if !strings.Contains(stdOut, fmt.Sprintf("successfully deleted a TODO with ID %d", nextID)) {
		t.Error("todocli remove didn't print success message")
	}

	cmd = exec.Command("../todocli", "get-all")
	out.Reset()
	cmd.Stdout = &out
	err = cmd.Run()

	if err != nil {
		t.Error("Failed to execute todocli get-all")
	}

	stdOut = out.String()

	search := fmt.Sprintf(
		`
Id: %d
Name: %s
Status: Unfinished
TODO before this:  []
`,
		nextID, "test")

	if strings.Contains(stdOut, search) {
		t.Error("TODO with Id", nextID, "found in all todos after deleting it")
	}

	// the next todo added will have a new ID
	nextID++
}

func TestCanMarkTodoAsCompleted(t *testing.T) {
	title := "write some very interesting CLI tests"
	exec.Command("../todocli", "todo", title).Run()

	exec.Command("../todocli", "check", strconv.Itoa(nextID), "do").Run()

	cmd := exec.Command("../todocli", "get-all")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		t.Error("Failed to execute todocli get-all")
	}

	stdOut := out.String()
	search := fmt.Sprintf(
		`
Id: %d
Name: %s
Status: Finished
TODO before this:  []
`,
		nextID, title)

	if !strings.Contains(stdOut, search) {
		t.Error("TODO with Id", nextID, "not found in all todos or doesn't have correct status")
	}

	// assume that it works here
	exec.Command("../todocli", "remove", strconv.Itoa(nextID)).Run()

	// the next todo added will have a new ID
	nextID++
}

/*
* Feel free to add more tests (testing dependency graph and stuff)
 */
