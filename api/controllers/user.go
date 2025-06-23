package controllers

import (
	"api/database"
	"api/models"
	"encoding/json"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	database.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var admin models.User
	json.NewDecoder(r.Body).Decode(&admin)
	database.DB.Create(admin)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(admin)
}
