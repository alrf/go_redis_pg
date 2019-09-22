package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"main/controllers"
)

func main() {

	router := mux.NewRouter()
	port := "8080"

	router.HandleFunc("/humai/echoservice", controllers.ListInventory).Methods("GET")
	router.HandleFunc("/humai/echoservice", controllers.CreateInventory).Methods("POST")

	log.Println("Starting app using port " + port)
	log.Fatal(http.ListenAndServe(":" + port, router))

}
