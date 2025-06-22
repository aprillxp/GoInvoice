package main

import (
	"log"
	"net/http"

	"api/database"
	"api/models"
	"api/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	database.ConnectDB()
	database.DB.AutoMigrate(&models.User{}, &models.Invoice{})

	router := mux.NewRouter()
	routes.Router(router)
	_ = godotenv.Load()

	log.Println("I Love You :8080")
	http.ListenAndServe(":8080", router)
}
