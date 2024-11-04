package handlers

import (
	"encoding/json"
	"golang-beginner-21/practice/database"
	"golang-beginner-21/practice/models"
	"golang-beginner-21/practice/repositories"
	"golang-beginner-21/practice/services"
	"golang-beginner-21/practice/utils"
	"net/http"
	"strconv"
)

// LoginHandler handles user login requests
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	db, err := database.InitDB()
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "Failed to connect to database", nil)
		return
	}
	defer db.Close()

	userRepo := repositories.NewUserRepositoryDB(db)
	userService := services.NewUserService(*userRepo)

	userFound, err := userService.Login(user)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusUnauthorized, "Invalid credentials", nil)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "User logged in successfully", userFound)
}

func GetUserByIdHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	idInput := r.URL.Query().Get("id")
	if idInput == "" {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	db, err := database.InitDB()
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "Failed to connect to database", nil)
		return
	}
	defer db.Close()

	userRepo := repositories.NewUserRepositoryDB(db)
	userService := services.NewUserService(*userRepo)
	id, _ := strconv.Atoi(idInput)
	user, err := userService.GetUserById(id)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusNotFound, "User not found", nil)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, "User found", user)
}
