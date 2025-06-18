package routes

import (
	"api/controllers"
	"github.com/gorilla/mux"
)

func Router(router *mux.Router) {
	// User
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")

	// Authentication
	router.HandleFunc("/users", controllers.Register).Methods("POST")
	// router.HandleFunc("/login", controllers.Login).Methods("POST")

	// Invoice
	// router.HandleFunc("/invoices", controllers.GetInvoices).Methods("GET")

}
