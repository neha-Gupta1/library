package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Controller() {

	r := mux.NewRouter()
	r.HandleFunc("/books", GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", GetBook).Methods("GET")
	r.HandleFunc("/books", CreateBook).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))

}
