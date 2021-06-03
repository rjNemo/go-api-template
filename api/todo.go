package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rjNemo/go-api-template/error"
)

const todoURI = "https://jsonplaceholder.typicode.com/todos/1"

// Todo is a task object
type Todo struct {
	UserID, ID int
	Title      string
	Completed  bool
}

// GetTodo fetches a Todo item from API
func GetTodo() *Todo {
	// perform request
	resp, err := http.Get(todoURI)
	error.Handle(err)
	// read data
	data, err := ioutil.ReadAll(resp.Body)
	error.Handle(err)
	// create Todo
	var t Todo
	err = json.Unmarshal(data, &t)
	error.Handle(err)
	return &t
}

// HandleGetTodo does this ...
func HandleGetTodo(w http.ResponseWriter, r *http.Request) {
	t := GetTodo()
	fmt.Fprintf(w, "Todo: %+v", t)
}
