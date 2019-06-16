package main

// User is a user with todos
type User struct {
	todos []Todo
}

// Todos returns all todo items
func (u User) Todos() []Todo {
	return u.todos
}

// AddTodo creates a new todo item for user
func (u *User) AddTodo(todo Todo) {
	(*u).todos = append(u.todos, todo)
}
