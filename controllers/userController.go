package controllers

import (
	"encoding/json"
	"net/http"
	"simple-rest-api-golang/helper"
	"simple-rest-api-golang/initializers"
	"simple-rest-api-golang/models"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	// Parsing the body
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var body models.RegisterInput
	if err := decoder.Decode(&body); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Validate the input
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusBadRequest)
		return
	}

	// Create new user
	user := models.User{Name: body.Name, Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user) // pass pointer of data to Create

	if result.Error != nil {
		http.Error(w, "Filed to create user", http.StatusBadRequest)
		return
	}

	// Respond
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("New user registered successfully!"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	// Parsing the body
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var body models.LoginInput
	if err := decoder.Decode(&body); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Validate the input
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Look up for users
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		http.Error(w, "Invalid email/password", http.StatusBadRequest)
		return
	}

	// Compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		http.Error(w, "Invalid email/password", http.StatusBadRequest)
		return
	}

	// Generate jwt token
	token, err := helper.GenerateToken(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// respond
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}
