package utils

import (
	"encoding/json"
	"golang-beginner-21/practice/models"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(models.Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	})
}
