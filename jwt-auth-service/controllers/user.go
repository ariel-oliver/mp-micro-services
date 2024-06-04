package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ariel-oliver/mp-micro-services/jwt-auth-service/config"
	"github.com/ariel-oliver/mp-micro-services/jwt-auth-service/models"
	"github.com/ariel-oliver/mp-micro-services/jwt-auth-service/utils"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        log.Println("Error decoding request body:", err)
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        log.Println("Error generating password hash:", err)
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }
    user.Password = string(hashedPassword)

    stmt, err := config.DB.Prepare("INSERT INTO users (username, password) VALUES (?, ?)")
    if err != nil {
        log.Println("Error preparing SQL statement:", err)
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }
    _, err = stmt.Exec(user.Username, user.Password)
    if err != nil {
        log.Println("Error executing SQL statement:", err)
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
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	// Verificar se o usuário existe
	var existingUser models.User
	err := config.DB.QueryRow("SELECT id, username, password FROM users WHERE id = ?", user.ID).Scan(&existingUser.ID, &existingUser.Username, &existingUser.Password)
	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Atualizar o nome de usuário e a senha do usuário no banco de dados
	stmt, err := config.DB.Prepare("UPDATE users SET username = ?, password = ? WHERE id = ?")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	_, err = stmt.Exec(user.Username, string(hashedPassword), user.ID)
	if err != nil {
		http.Error(w, "Could not update user", http.StatusInternalServerError)
		return
	}

	// Retornar o usuário atualizado como resposta
	json.NewEncoder(w).Encode(user)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	// Verificar se o usuário existe
	var existingUser models.User
	err := config.DB.QueryRow("SELECT id, username, password FROM users WHERE id = ?", user.ID).Scan(&existingUser.ID, &existingUser.Username, &existingUser.Password)
	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Deletar o usuário do banco de dados
	stmt, err := config.DB.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	_, err = stmt.Exec(user.ID)
	if err != nil {
		http.Error(w, "Could not delete user", http.StatusInternalServerError)
		return
	}

	// Retornar uma resposta de sucesso
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}