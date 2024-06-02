package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/ariel-oliver/mp-micro-services/jwt-auth-service/config"
	"github.com/ariel-oliver/mp-micro-services/jwt-auth-service/models"
	"github.com/ariel-oliver/mp-micro-services/jwt-auth-service/utils"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	stmt, err := config.DB.Prepare("INSERT INTO users (username, password) VALUES (?, ?)")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	_, err = stmt.Exec(user.Username, user.Password)
	if err != nil {
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var input models.User
	var user models.User
	json.NewDecoder(r.Body).Decode(&input)

	err := config.DB.QueryRow("SELECT id, username, password FROM users WHERE username = ?", input.Username).Scan(&user.ID, &user.Username, &user.Password)
	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
