// di controller/controllers_test.go atau sejenis:
package controllers

import (
	"api/utils"
	"encoding/json"
	"net/http"
)

type SendMailPayload struct {
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func TestSendgrid(w http.ResponseWriter, r *http.Request) {
	var payload SendMailPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	utils.SendInvoiceMail(payload.Email, payload.Subject, payload.Body)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("✅ Email sent (or attempted)"))
}
