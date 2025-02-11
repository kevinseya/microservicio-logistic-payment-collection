package services

import (
	"payment-collection/models"
	"payment-collection/repositories"

	"github.com/google/uuid"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

func CreatePayment(orderID uuid.UUID, amount float64, currency string) (*stripe.PaymentIntent, error) {
	// Convertir el amount a la unidad más pequeña (por ejemplo, centavos)
	amountInCents := int64(amount * 100) // Asumiendo que amount está en dólares

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amountInCents), // Cambiado a Int64
		Currency: stripe.String(currency),
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		return nil, err
	}

	payment := models.Payment{
		OrderID:         orderID, // Cambiado a orderID en lugar de uuid.New()
		Amount:          amount,  // Se mantiene como float64
		Currency:        currency,
		Status:          "PENDING",
		PaymentIntentID: pi.ID,
	}

	err = repositories.SavePayment(&payment)
	return pi, err
}
