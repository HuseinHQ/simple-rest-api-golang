package models

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required"`
}

var Users = []User{}

type UserHandler struct{}

func (uh *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Users)
}

func (uh *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var userInput RegisterUser

	if err := decoder.Decode(&userInput); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	newUser := User{
		ID:       GenerateUserID(),
		Name:     userInput.Name,
		Email:    userInput.Email,
		Password: userInput.Password,
	}

	Users = append(Users, newUser)

	w.Write([]byte("New user registered successfully!"))
}

func GenerateUserID() int {
	if len(Users) == 0 {
		return 1
	}
	return Users[len(Users)-1].ID + 1
}
