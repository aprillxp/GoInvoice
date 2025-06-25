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
		log.Println("❌ Error reading body:", err)
		http.Error(w, "Request body read error", http.StatusServiceUnavailable)
		return
	}

	endPointSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")
	if endPointSecret == "" {
		log.Println("❌ STRIPE_WEBHOOK_SECRET not set")
		http.Error(w, "Secret not configured", http.StatusBadRequest)
		return
	}

	sigHeader := r.Header.Get("Stripe-Signature")
	if sigHeader == "" {
		log.Println("❌ Missing Stripe-Signature header")
		http.Error(w, "Missing signature", http.StatusBadRequest)
		return
	}

	event, err := webhook.ConstructEventWithOptions(payload, sigHeader, endPointSecret, webhook.ConstructEventOptions{
		IgnoreAPIVersionMismatch: true,
	})

	if err != nil {
		log.Println("❌ Webhook verification failed:", err)
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
