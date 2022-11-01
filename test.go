package main

import (
	"go-trial/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", repository.Home)
	router.HandleFunc("/model", repository.Create).Methods("POST")
	router.HandleFunc("/models", repository.GetAll).Methods("GET")
	router.HandleFunc("/models/{id}", repository.GetOne).Methods("GET")
	router.HandleFunc("/models/{id}", repository.Update).Methods("PATCH")
	router.HandleFunc("/models/{id}", repository.DeleteById).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":6666", router))
}
