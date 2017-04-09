package srest

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome ya m3lm !")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId, err := strconv.Atoi(vars["todoId"])
	if err != nil {
		panic(err)
	}
	t := RepoFindTodo(todoId)
	if t.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusFound)
		if err := json.NewEncoder(w).Encode(t); err != nil {
			panic(err)
		}
	}
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	fmt.Println(body)
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		panic(err)
	}

	t := RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}

}


func TodoDelete(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	todoId, err := strconv.Atoi(vars["todoId"])
	if err != nil {
		panic(err)
	}
	t := RepoDestroyTodo(todoId)

	if t.code == 0 {
		w.WriteHeader(http.StatusAccepted)
	}else{
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(t); err != nil {
			panic(err)
		}
	}
}
