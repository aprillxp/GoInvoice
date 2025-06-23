package handlers

import (
	"api/database"
	"api/models"
	"api/utils"
	"encoding/json"
	"net/http"
)

func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Amount int64 `json:"amount"`
		UserID uint  `json:"user_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	invoice := models.Invoice{
		UserID: req.UserID,
		Amount: req.Amount,
		Paid:   false,
	}
	if err := database.DB.Create(&invoice).Error; err != nil {
		http.Error(w, "Failed to create invoice", http.StatusInternalServerError)
		return
	}

	url, err := utils.StripeSession(int64(invoice.Amount*100), invoice.ID)
	if err != nil || url == "" {
		http.Error(w, "Stripe session creation failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"url": url})
}
