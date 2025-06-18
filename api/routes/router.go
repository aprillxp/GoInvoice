package routes

import (
	"api/controllers"
	"api/middleware"

	"github.com/gorilla/mux"
)

func Router(router *mux.Router) {
	// Authentication
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")

	// Middleware
	secured := router.PathPrefix("/api").Subrouter()
	secured.Use(middleware.Authorization)

	// User
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")

	// Invoice
	// router.HandleFunc("/invoices", controllers.GetInvoices).Methods("GET")

}
