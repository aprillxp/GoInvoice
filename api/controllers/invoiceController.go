package controllers

import (
	"api/database"
	"api/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetInvoices(w http.ResponseWriter, r *http.Request) {
	var invoices []models.Invoice
	database.DB.Find(invoices)
	json.NewEncoder(w).Encode(invoices)
}

func GetInvoiceByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var invoice []models.Invoice
	if err := database.DB.First(&invoice, id).Error; err != nil {
		http.Error(w, "Invoice not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(invoice)
}

func CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var result []models.Invoice
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	database.DB.Create(&result)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func UpdateInvoice(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var invoiceID []models.Invoice
	if err := database.DB.First(&invoiceID, id).Error; err != nil {
		http.Error(w, "Invoice not found", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&invoiceID); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	database.DB.Save(&invoiceID)
	json.NewEncoder(w).Encode(invoiceID)
}

func DeleteInvoice(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	database.DB.Delete(&models.Invoice{}, id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Invoice deleted"})
}
