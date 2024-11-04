package handlers

import (
	"encoding/json"
	"golang-beginner-21/database"
	"golang-beginner-21/models"
	"golang-beginner-21/repositories"
	"golang-beginner-21/services"
	"golang-beginner-21/utils"
	"net/http"
)

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RespondWithJSON(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	var todoInput models.Todos
	if err := json.NewDecoder(r.Body).Decode(&todoInput); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}
	token := r.Header.Get("token")

	db, err := database.InitDB()
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "Error connecting to the database", nil)
		return
	}

	todoRepo := repositories.NewTodoRepositoryDB(db)
	todoService := services.NewTodoService(*todoRepo)
	newTodo, err := todoService.CreateTodo(&todoInput, token)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "Error creating todo", nil)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, "To do created successfully", newTodo)
}

func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RespondWithJSON(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	token := r.Header.Get("token")

	db, err := database.InitDB()
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "Error connecting to the database", nil)
		return
	}

	todoRepo := repositories.NewTodoRepositoryDB(db)
	todoService := services.NewTodoService(*todoRepo)
	todos, err := todoService.GetTodosByUserId(token)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, "Todos fetched successfully", todos)
}

func GetTodoCountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RespondWithJSON(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	token := r.Header.Get("token")
	db, err := database.InitDB()
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "Error connecting to the database", nil)
		return
	}

	todoRepo := repositories.NewTodoRepositoryDB(db)
	todoService := services.NewTodoService(*todoRepo)
	todos, err := todoService.GetTodoCount(token)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, "Todos fetched successfully", todos)
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RespondWithJSON(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	var todoInput models.Todos
	if err := json.NewDecoder(r.Body).Decode(&todoInput); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	db, err := database.InitDB()
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "Error connecting to the database", nil)
		return
	}
	todoRepo := repositories.NewTodoRepositoryDB(db)
	todoService := services.NewTodoService(*todoRepo)
	updatedTodo, err := todoService.UpdateTodo(&todoInput)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "Error updating todo", nil)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, "Todo updated successfully", updatedTodo)
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RespondWithJSON(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	var todo models.Todos
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	db, err := database.InitDB()
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "Error connecting to the database", nil)
		return
	}
	todoRepo := repositories.NewTodoRepositoryDB(db)
	todoService := services.NewTodoService(*todoRepo)
	err = todoService.DeleteTodo(todo.ID)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "Error updating todo", nil)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, "Todo updated successfully", nil)
}
