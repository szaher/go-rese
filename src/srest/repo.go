package srest

import "fmt"

var currentId int
var todos Todos

// Give us some seed data
func init() {
	RepoCreateTodo(Todo{Name: "Write Presentation"})
	RepoCreateTodo(Todo{Name: "Host Meetup!"})
}


func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return Todo{}
}

func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}


func RepoDestroyTodo(id int) MyError {
	for i, t := range todos{
		if t.Id == id{
			todos = append(todos[:1], todos[i+1:]...)
			return MyError{code: 0}
		}
	}

	return MyError{code: 404, Message: fmt.Sprintf("Couldnot find Todo with id %d to delete", id)}
}