package main

import (
	"log"
	"net/http"

	"github.com/shohratd15/todolist-api/internal/db"
	"github.com/shohratd15/todolist-api/internal/handlers"
)

func main() {
	db.Connect()

	port := "8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/tasks",handlers.TasksHandler)
	mux.HandleFunc("/tasks/",handlers.TaskHandler)

	log.Println("Starting server on port", port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}