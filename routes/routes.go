package routes

import (
	"payment-collection/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/pay", controllers.ProcessPayment)
	r.POST("/pay/webhook", controllers.HandleWebhook)
	return r
}
