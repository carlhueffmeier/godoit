package main

import "net/http"

// APIHandler is responsible for routing all api requests
type APIHandler struct {
	todoController *TodoController
}

func (h *APIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/api/todos":
		switch r.Method {
		case "GET":
			h.todoController.getAllTodos(w, r)
		case "POST":
			h.todoController.addTodo(w, r)
		}
	}
}
