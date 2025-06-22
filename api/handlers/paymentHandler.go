package handlers

import (
	"api/utils"
	"encoding/json"
	"net/http"
)

func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Amount int64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	url, err := utils.StripeSession(req.Amount)
	if err != nil || url == "" {
		http.Error(w, "Stripe session creation failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"url": url})
}
