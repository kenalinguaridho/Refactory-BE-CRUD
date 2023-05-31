package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kenalinguaridho/config"
	"github.com/kenalinguaridho/controller"
)

func main() {
	config.DBConnection()

	router := mux.NewRouter()

	router.HandleFunc("/item", controller.Create).Methods("POST")
	router.HandleFunc("/items", controller.GetAll).Methods("GET")
	router.HandleFunc("/item/{id}", controller.GetItemById).Methods("GET")
	router.HandleFunc("/item/{id}/update", controller.Update).Methods("PUT")
	router.HandleFunc("/item/{id}/delete", controller.Delete).Methods("DELETE")

	http.ListenAndServe(":8080", router)

}