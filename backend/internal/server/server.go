package server

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/ryan-tech/gameoflife/internal/handlers"
)

func Start() {
	log.Println("Starting server...")
	router := mux.NewRouter()
	router.HandleFunc("/step", handlers.Step).Methods("POST")
	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", router)
}