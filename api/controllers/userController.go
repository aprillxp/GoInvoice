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

func CreateUsers(w http.ResponseWriter, r *http.Request) {
	var u models.User
	json.NewDecoder(r.Body).Decode(&u)
	database.DB.Create(&u)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}
