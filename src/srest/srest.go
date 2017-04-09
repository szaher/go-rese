package srest

import (
	"log"
	"net/http"
)

func Srest(){
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
