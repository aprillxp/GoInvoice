package handlers

import (
	"api/database"
	"api/models"
	"api/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/context"
)

func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := context.Get(r, "user_id").(uint)
	if !ok || userID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	var req struct {
		InvoiceID uint `json:"invoice_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var invoice models.Invoice
	if err := database.DB.Where("id = ? AND user_id = ?", req.InvoiceID, userID).First(&invoice).Error; err != nil {
		http.Error(w, "Invoice not found or unauthorized", http.StatusNotFound)
		return
	}

	amountInCents := int64(invoice.Amount * 100)
	url, err := utils.StripeSession(amountInCents, invoice.ID)
	if err != nil || url == "" {
		http.Error(w, "Stripe session creation failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"url": url})
}
