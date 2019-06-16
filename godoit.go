package main

import (
	"log"
	"net/http"
)

var user User

func main() {
	api := &APIHandler{new(TodoController)}
	http.Handle("/api/", http.StripPrefix("/api/", api))
	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/", fs)

	log.Fatal(http.ListenAndServe("localhost:4440", nil))
}
