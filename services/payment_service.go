package services

import (
	"payment-collection/models"
	"payment-collection/repositories"

	"github.com/google/uuid"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

func CreatePayment(orderID uuid.UUID, amount float64, currency string) (*stripe.PaymentIntent, error) {
	// Convert the amount to the smallest unit (e.g. cents)
	amountInCents := int64(amount * 100) // Assuming amount is in dollars

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amountInCents), // Changed to Int64
		Currency: stripe.String(currency),
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		return nil, err
	}

	payment := models.Payment{
		OrderID:         orderID,
		Amount:          amount,
		Currency:        currency,
		Status:          "PENDING",
		PaymentIntentID: pi.ID,
	}

	err = repositories.SavePayment(&payment)
	return pi, err
}
