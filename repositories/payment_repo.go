package repositories

import (
	"payment-collection/config"
	"payment-collection/models"
)

func SavePayment(payment *models.Payment) error {
	return config.DB.Create(payment).Error
}

func GetPaymentByIntentID(paymentIntentID string) (models.Payment, error) {
	var payment models.Payment
	err := config.DB.Where("payment_intent_id = ?", paymentIntentID).First(&payment).Error
	return payment, err
}

func UpdatePaymentStatus(payment *models.Payment) error {
	return config.DB.Save(payment).Error
}
