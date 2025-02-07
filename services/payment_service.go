package services

import (
	"payment-collection/models"
	"payment-collection/repositories"

	"github.com/google/uuid"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

func CreatePayment(orderID uuid.UUID, amount int64, currency string) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(currency),
	}
	pi, err := paymentintent.New(params)
	if err != nil {
		return nil, err
	}

	payment := models.Payment{
		OrderID:         uuid.New(),
		Amount:          amount,
		Currency:        currency,
		Status:          "PENDING",
		PaymentIntentID: pi.ID,
	}

	err = repositories.SavePayment(&payment)
	return pi, err
}

func UpdatePaymentStatus(paymentIntentID, status string) error {
	payment, err := repositories.GetPaymentByIntentID(paymentIntentID)
	if err != nil {
		return err
	}

	payment.Status = status
	return repositories.UpdatePaymentStatus(&payment)
}
