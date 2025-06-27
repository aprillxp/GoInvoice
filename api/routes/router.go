package routes

import (
	"api/controllers"
	handlers "api/handlers/stripe"
	"api/middleware"

	"github.com/gorilla/mux"
)

func Router(router *mux.Router) {

	// Public routes
	router.HandleFunc("/register", controllers.Register).Methods("POST") // ok
	router.HandleFunc("/login", controllers.Login).Methods("POST")       // ok

	router.HandleFunc("/webhook", handlers.StripeWebhook).Methods("POST") // ok

	// testing email
	router.HandleFunc("/sendmail", controllers.TestSendgrid).Methods("POST")

	// Middleware / Protected routes
	secured := router.PathPrefix("/api").Subrouter()
	secured.Use(middleware.Authorization)

	// User (secured)
	secured.HandleFunc("/users", controllers.GetUsers).Methods("GET") // ok

	// Invoice (secured)
	secured.HandleFunc("/invoices", controllers.GetInvoices).Methods("GET")           // ok
	secured.HandleFunc("/invoices/{id}", controllers.GetInvoiceByID).Methods("GET")   // ok
	secured.HandleFunc("/invoices", controllers.CreateInvoice).Methods("POST")        // ok
	secured.HandleFunc("/invoices/{id}", controllers.UpdateInvoice).Methods("PUT")    // ok
	secured.HandleFunc("/invoices/{id}", controllers.DeleteInvoice).Methods("DELETE") // ok

	// Webhook
	secured.HandleFunc("/pay", handlers.PaymentHandler).Methods("POST") // ok

}
