package routes

import (
	"api/controllers"
	"api/middleware"

	"github.com/gorilla/mux"
)

func Router(router *mux.Router) {
	// Public routes
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")

	// Middleware / Protected routes
	secured := router.PathPrefix("/api").Subrouter()
	secured.Use(middleware.Authorization)

	// User (secured)
	secured.HandleFunc("/users", controllers.GetUsers).Methods("GET")

	// Invoice (secured)
	router.HandleFunc("/invoices", controllers.GetInvoices).Methods("GET")
	router.HandleFunc("/invoices/{id}", controllers.GetInvoiceByID).Methods("GET")
	router.HandleFunc("/invoices/", controllers.CreateInvoice).Methods("POST")
	router.HandleFunc("/invoices/{id}", controllers.UpdateInvoice).Methods("PUT")
	router.HandleFunc("/invoices/{id}", controllers.DeleteInvoice).Methods("DELETE")

}
