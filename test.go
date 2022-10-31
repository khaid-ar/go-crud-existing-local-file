package main

import (
	"go-trial/entity"
	"go-trial/repository"
	"go-trial/util"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	var data entity.Models
	util.ParseToString(&data)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", repository.Home)
	router.HandleFunc("/model", repository.Create).Methods("POST")
	router.HandleFunc("/models", repository.GetAll).Methods("GET")
	router.HandleFunc("/models/{id}", repository.GetOne).Methods("GET")
	router.HandleFunc("/models/{id}", repository.Update).Methods("PATCH")
	router.HandleFunc("/models/{id}", repository.DeleteById).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":6666", router))

	// tes := []int{1, 2, 3, 4, 5}
	// tes = append(tes, 6)
	// data.Models = append(data.Models, entity.Model{Id: "202"})
	// fmt.Println(tes)
	// fmt.Println(data.Models)

}
