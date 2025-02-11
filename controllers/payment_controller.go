package controllers

import (
	"net/http"
	"payment-collection/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ProcessPayment(c *gin.Context) {
	type PaymentRequest struct {
		OrderID  string  `json:"orderId"`
		Amount   float64 `json:"amount"` // Remains as float64
		Currency string  `json:"currency"`
	}

	var req PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Validate the OrderID format
	orderUUID, err := uuid.Parse(req.OrderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Order ID format"})
		return
	}

	// Create the payment
	pi, err := services.CreatePayment(orderUUID, req.Amount, req.Currency)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment: " + err.Error()})
		return
	}

	// Reply with the result
	c.JSON(http.StatusOK, gin.H{
		"paymentIntent": pi.ID,
		"clientSecret":  pi.ClientSecret,
	})
}
