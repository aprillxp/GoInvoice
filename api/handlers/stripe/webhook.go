package handlers

import (
	"api/database"
	"api/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/webhook"
)

func StripeWebhook(w http.ResponseWriter, r *http.Request) {
	const MaxBodyBytes = int64(65536)
	r.Body = http.MaxBytesReader(w, r.Body, MaxBodyBytes)
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Request body read error", http.StatusServiceUnavailable)
		return
	}

	endPointSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")
	sigHeader := r.Header.Get("Stripe-Signature")

	event, err := webhook.ConstructEvent(payload, sigHeader, endPointSecret)
	if err != nil {
		http.Error(w, "Webhook verification failed", http.StatusBadRequest)
		return
	}

	switch event.Type {
	case "checkout.session.completed":
		var session stripe.CheckoutSession
		if err := json.Unmarshal(event.Data.Raw, &session); err == nil {
			invoiceIDStr := session.Metadata["invoice_id"]
			invoiceID, err := strconv.Atoi(invoiceIDStr)
			if err == nil {
				var invoice models.Invoice
				if err := database.DB.First(&invoice, invoiceID).Error; err == nil {
					invoice.Paid = true
					database.DB.Save(&invoice)
					log.Println("Payment success for session", session.ID)
				}
			}
		}
	default:
		fmt.Println("Unhandled event type:", event.Type)
	}

	w.WriteHeader(http.StatusOK)
}
