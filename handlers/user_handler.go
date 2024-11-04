package handlers

import (
	"encoding/json"
	"net/http"

	"golang-beginner-21/database"
	"golang-beginner-21/models"
	"golang-beginner-21/repositories"
	"golang-beginner-21/services"
	"golang-beginner-21/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		utils.RespondWithJSON(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	// Decode the request body
	var login models.User
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	// Initialize the database
	db, err := database.InitDB()
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "Database connection error", nil)
		return
	}
	defer db.Close()

	// Create repository and service
	userRepo := repositories.NewUserRepositoryDB(db)
	userService := services.NewUserService(*userRepo)

	// Authenticate the user
	user, err := userService.LoginService(login.Email, login.Password)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusUnauthorized, "Invalid credentials", nil)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, "Login successful", user)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		utils.RespondWithJSON(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	var userInput models.User
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	db, err := database.InitDB()
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "Database connection error", nil)
		return
	}
	defer db.Close()

	userRepo := repositories.NewUserRepositoryDB(db)
	userService := services.NewUserService(*userRepo)
	newUser, err := userService.CreateUser(&userInput)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, "User created successfully", newUser)
}
