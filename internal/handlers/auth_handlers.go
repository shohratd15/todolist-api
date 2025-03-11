package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/shohratd15/todolist-api/internal/db"
	"github.com/shohratd15/todolist-api/internal/logger"
	"github.com/shohratd15/todolist-api/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims {
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func RegisterHandler(w http.ResponseWriter, r *http.Request){
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Log.Error("Error decoding request body", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		logger.Log.Error("Failed to hash password", err)
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	user := &models.User{
		Username: req.Username,
		Password: hashedPassword,
	}

	if err := db.CreateUser(user); err != nil {
		logger.Log.Error("User creation failed", err)
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	logger.Log.Infof("User %s registered successfully", user.Username)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message":"User registered succesfully"}) 
}

func LoginHandler(w http.ResponseWriter, r *http.Request){
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Log.Error("Error decoding request body", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := db.GetUserByUsername(req.Username)
	if err != nil || !checkPasswordHash(req.Password, user.Password) {
		logger.Log.Error("Invalid credentials", err)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := generateToken(user)
	if err != nil {
		logger.Log.Error("Failed to generated token", err)
		http.Error(w, "Failed to generated token", http.StatusInternalServerError)
		return
	}

	logger.Log.Infof("User %s logined successfully", user.Username)

	json.NewEncoder(w).Encode(map[string]string {"token":token})
}