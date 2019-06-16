package main

import (
	"log"
	"net/http"
)

var user User

func main() {
	http.Handle("/api/todos", &APIHandler{new(TodoController)})
	http.Handle("/", new(StaticFileHandler))
	log.Fatal(http.ListenAndServe("localhost:4440", nil))
}
