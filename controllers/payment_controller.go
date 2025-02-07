package controllers

import (
	"net/http"
	"payment-collection/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ProcessPayment(c *gin.Context) {
	type PaymentRequest struct {
		OrderID  string `json:"orderId"`
		Amount   int64  `json:"amount"`
		Currency string `json:"currency"`
	}

	var req PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderUUID, err := uuid.Parse(req.OrderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Order ID format"})
		return
	}

	pi, err := services.CreatePayment(orderUUID, req.Amount, req.Currency)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"paymentIntent": pi.ID,
		"clientSecret":  pi.ClientSecret,
	})
}

func HandleWebhook(c *gin.Context) {
	var payload map[string]interface{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, ok := payload["data"].(map[string]interface{})
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	obj, ok := data["object"].(map[string]interface{})
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid object format"})
		return
	}

	paymentIntentID, ok := obj["id"].(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment intent ID"})
		return
	}

	status, ok := obj["status"].(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status format"})
		return
	}

	if err := services.UpdatePaymentStatus(paymentIntentID, status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Webhook processed"})
}
