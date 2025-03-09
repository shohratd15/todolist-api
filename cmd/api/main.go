package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shohratd15/todolist-api/internal/db"
)

func main() {
	db.Connect()

	port := "8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, ToDo API!"))
	})

	log.Println("Starting server on port", port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}