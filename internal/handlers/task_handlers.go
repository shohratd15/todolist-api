package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/shohratd15/todolist-api/internal/db"
	"github.com/shohratd15/todolist-api/internal/models"
)

func TasksHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	
	case http.MethodGet:
		tasks, err := db.GetAllTasks()
		if err != nil {
			http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(tasks)
	
	case http.MethodPost:
		var task models.Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if err := db.CreateTask(&task); err != nil {
			http.Error(w, "Failed to create task", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(task)

	default:
		http.Error(w, "Method not allowed",http.StatusMethodNotAllowed)
	}
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid request path", http.StatusBadRequest)
		return
	}
	
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	
	switch r.Method {
	case http.MethodGet:
		task, err := db.GetTaskByID(id)
		if err != nil {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(task)
	case http.MethodPut:
		var task models.Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		task.ID = id
		if err := db.UpdateTask(&task); err != nil {
			http.Error(w,"Failed to update task", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(task)
	case http.MethodDelete:
		if err := db.DeleteTask(id); err != nil {
			http.Error(w, "Failed to delete task", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}