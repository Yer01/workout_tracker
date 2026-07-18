package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Yer01/workout_tracker/internal/services"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	auth_service services.AuthService
}

func NewAuthHandler(authservice services.AuthService) *AuthHandler {
	return &AuthHandler{
		auth_service: authservice,
	}
}

func (ah *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	existingUserID, err := ah.auth_service.GetSingle(email)

	if err == nil && existingUserID != 0 {
		log.Println("User with same email already exists")
		http.Error(w, "user already exists", http.StatusConflict)
		return
	}
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Printf("Couldn't create user: %v\n", err)
		http.Error(w, "failed to check user", http.StatusInternalServerError)
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Printf("Couldn't hash password: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := ah.auth_service.CreateUser(email, string(hashed), username)
	if id == 0 || err != nil {
		log.Printf("Couldn't create user: %v\n", err)
		http.Error(w, "failed to create user", http.StatusBadRequest)
		return
	}

	var response struct {
		Msg string `json:"msg"`
	}
	response.Msg = "User created succesfully"

	if err = json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Problem with encoding response: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("User created succesfully!")
}
