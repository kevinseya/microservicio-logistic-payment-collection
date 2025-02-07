package models

import (
	"github.com/google/uuid"
)

type Payment struct {
	ID              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	OrderID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()" json:"orderId"`
	Amount          int64     `json:"amount"`
	Currency        string    `json:"currency"`
	Status          string    `json:"status"`
	PaymentIntentID string    `json:"paymentIntentId"`
}
