package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// TodoController knows how to read and update todos
type TodoController struct{}

func (TodoController) getAllTodos(w http.ResponseWriter, r *http.Request) {
	user.Todos()
	data, err := json.Marshal(user.Todos())
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	w.Write(data)
}

func (TodoController) addTodo(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	t := Todo{}
	if err := dec.Decode(&t); err != nil {
		log.Fatal(err)
		return
	}
	user.AddTodo(t)
}
