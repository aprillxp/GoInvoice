package controllers

import (
	"api/database"
	"api/models"
	"api/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func GetInvoices(w http.ResponseWriter, r *http.Request) {
	userID, ok := context.Get(r, "user_id").(uint)
	if !ok || userID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var invoices []models.Invoice
	database.DB.Where("user_id = ?", userID).Find(&invoices)
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
	userID, ok := context.Get(r, "user_id").(uint)
	if !ok || userID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var result []models.Invoice
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	for i := range result {
		result[i].UserID = userID
	}

	if err := database.DB.Create(&result).Error; err != nil {
		http.Error(w, "Failed to create invoice", http.StatusInternalServerError)
		return
	}

	// Send email
	var user models.User
	if err := database.DB.First(&user, userID).Error; err == nil {
		for _, invoice := range result {
			body := fmt.Sprintf("An invoice of Rp%.2f has been created.", invoice.Amount)
			utils.SendInvoiceMail(user.Email, "New invoice", body)
		}
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
