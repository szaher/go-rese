package srest

import "time"

type Todo struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
	Due time.Time `json:"due"`
}

type Todos []Todo


type MyError struct {
	code int `json:"code"`
	Message string `json:"message"`
}