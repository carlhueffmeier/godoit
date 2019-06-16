package main

import "fmt"

// A Todo item
type Todo struct {
	Title string
	Done  bool
}

func (t Todo) String() string {
	return fmt.Sprintf("%s (done=%v)", t.Title, t.Done)
}
