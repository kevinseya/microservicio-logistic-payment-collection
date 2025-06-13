package controllers

import (
	"net/http"
	"payment-collection/models" // Import models
	"payment-collection/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ProcessPayment processes a payment with Stripe
// @Summary Creates a payment intent
// @Description Processes a payment and returns a Payment Intent
// @Tags Payments
// @Accept json
// @Produce json
// @Param paymentRequest body models.PaymentRequest true "Datos del pago"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /payment/create-intent [post]
func ProcessPayment(c *gin.Context) {
	var req models.PaymentRequest // Use the `models` structure

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
