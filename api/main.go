package main

import (
	"log"
	"net/http"

	"api/database"
	"api/models"
	"api/routes"
	"github.com/gorilla/mux"
)

func main() {
	database.ConnectDB()
	database.DB.AutoMigrate(&models.User{}, &models.Invoice{})

	router := mux.NewRouter()
	routes.Router(router)

	log.Println("Server is running at PORT:8080")
	http.ListenAndServe(":8080", router)
}
