package db

import (
	"context"
	"errors"
	"log"

	"github.com/shohratd15/todolist-api/internal/models"
)

func CreateUser(user *models.User) error {
	query := `INSERT INTO users(username, password) VALUES($1, $2) RETURNING id, created_at`
	err := DB.QueryRow(context.Background(), query, user.Username, user.Password).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		log.Println("Error creating user:",err)
		return err
	}
	return nil
}

func GetUserByUsername(username string) (*models.User, error){
	query := `SELECT id, username, password, created_at FROM users WHERE username=$1`
	var user models.User
	err := DB.QueryRow(context.Background(), query, username).Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

