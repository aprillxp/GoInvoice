package routes

import (
	"api/controllers"
	"github.com/gorilla/mux"
)

func RegisterRouter(r *mux.Router) {
	// User
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/users", controllers.CreateUsers).Methods("POST")
	// r.HandleFunc("/login", controllers.Login).Methods("POST")

	// Invoice
	// r.HandleFunc("/invoices", controllers.GetInvoices).Methods("GET")
	// r.HandleFunc("/invoices/{id}", controllers.GetInvoiceByID).Methods("GET")
	// r.HandleFunc("/invoices", controllers.CreateInvoice).Methods("POST")
	// r.HandleFunc("/invoices/{id}", controllers.UpdateInvoice).Methods("PUT")
	// r.HandleFunc("/invoices/{id}", controllers.DeleteInvoice).Methods("DELETE")
}
