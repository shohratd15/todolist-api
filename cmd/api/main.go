package main

import (
	"net/http"

	"github.com/shohratd15/todolist-api/internal/config"
	"github.com/shohratd15/todolist-api/internal/db"
	"github.com/shohratd15/todolist-api/internal/handlers"
	"github.com/shohratd15/todolist-api/internal/logger"
	"github.com/shohratd15/todolist-api/internal/middleware"
)

func main() {

	logger.InitLogger()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Log.Fatalf("Failed to load config: %v", err)
	}

	db.Connect(cfg)

	logger.Log.Infof("Server running on port: %s", cfg.Port)

	mux := http.NewServeMux()
	mux.HandleFunc("/register", handlers.RegisterHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)

	protectedMux := http.NewServeMux()
	protectedMux.HandleFunc("/tasks", handlers.TasksHandler)
	protectedMux.HandleFunc("/tasks/", handlers.TaskHandler)

	// Применяем middleware для логирования
	muxWithLogging := middleware.LoggingMiddleware(mux)
	protectedMuxWithLogging := middleware.LoggingMiddleware(protectedMux)

	// Защищаем маршруты
	authenticatedRoutes := middleware.AuthMiddleware(protectedMuxWithLogging)

	
	// Основной роутер
	mainMux := http.NewServeMux()
	mainMux.Handle("/register", muxWithLogging)
	mainMux.Handle("/login", muxWithLogging)
	mainMux.Handle("/", authenticatedRoutes) // Все маршруты, кроме регистрации и логина, требуют авторизации

	if err := http.ListenAndServe(":"+cfg.Port, mainMux); err != nil {
		logger.Log.Fatalf("Error starting server: %v", err)
	}
}