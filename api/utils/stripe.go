package utils

import (
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/checkout/session"

	"os"
)

func StripeSession(amount int64) (string, error) {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{{
			PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
				Currency: stripe.String("idr"),
				ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
					Name: stripe.String("Invoice payment"),
				},
				UnitAmount: stripe.Int64(amount),
			},
			Quantity: stripe.Int64(1),
		}},
		Mode:       stripe.String("payment"),
		SuccessURL: stripe.String("http://localhost:8080/success"),
		CancelURL:  stripe.String("http://localhost:8080/cancel"),
	}

	s, err := session.New(params)
	if err != nil {
		return "", err
	}
	return s.URL, nil
}
