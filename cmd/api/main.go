package main

import (
	"log"
	"net/http"
	"os"

	"github.com/shohratd15/todolist-api/internal/db"
	"github.com/shohratd15/todolist-api/internal/handlers"
	"github.com/shohratd15/todolist-api/internal/middleware"
)

func main() {
	os.Setenv("JWT_SECRET", "supersecretkey") // В реальном проекте использовать ENV
	db.Connect()

	mux := http.NewServeMux()
	mux.HandleFunc("/register", handlers.RegisterHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)

	protectedMux := http.NewServeMux()
	protectedMux.HandleFunc("/tasks", handlers.TasksHandler)
	protectedMux.HandleFunc("/tasks/", handlers.TaskHandler)

	// Защищаем маршруты
	authenticatedRoutes := middleware.AuthMiddleware(protectedMux)
	
	// Основной роутер
	mainMux := http.NewServeMux()
	mainMux.Handle("/register", mux)
	mainMux.Handle("/login", mux)
	mainMux.Handle("/", authenticatedRoutes) // Все маршруты, кроме регистрации и логина, требуют авторизации

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", mainMux)
}