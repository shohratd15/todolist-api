package db

import (
	"context"
	"log"

	"github.com/shohratd15/todolist-api/internal/models"
)

func CreateTask(task *models.Task) error {
	query := `INSERT INTO tasks(title, description, completed) VALUES($1,$2,$3) RETURNING id, created_at`
	err := DB.QueryRow(context.Background(), query, task.Title, task.Description, task.Completed).Scan(&task.ID,&task.CreatedAt)
	if err != nil {
		log.Println("Error creating task:", err)
		return err
	}

	return nil
}

func GetAllTasks() ([]models.Task, error) {
	query := `SELECT id, title, description, completed, created_at FROM tasks`
	rows, err := DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GetTaskByID(id int) (*models.Task, error){
	query := `SELECT id, title, description, completed, created_at FROM tasks WHERE id = $1`
	var task models.Task
	err := DB.QueryRow(context.Background(), query, id).Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func UpdateTask(task *models.Task) error {
	query := `UPDATE tasks SET title=$1, description=$2, completed=$3 WHERE id = $4`
	_, err := DB.Exec(context.Background(), query, task.Title, task.Description, task.Completed, task.ID)
	return err
}

func DeleteTask(id int) error {
	query := `DELETE FROM tasks WHERE id=$1`
	_, err := DB.Exec(context.Background(), query, id)
	return err
}