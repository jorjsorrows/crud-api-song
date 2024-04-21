package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/songs", getSongs).Methods("GET")
	r.HandleFunc("/songs", createSongs).Methods("POST")
	r.HandleFunc("/songs/{id}", getSong).Methods("GET")
	r.HandleFunc("/songs/{id}", deleteSongs).Methods("DELETE")
	r.HandleFunc("/songs/{id}", updateSongs).Methods("PUT")

	log.Fatal(http.ListenAndServe(":9000", r))

}

func main() {
	InitialMigration()
	InitializeRouter()
}
