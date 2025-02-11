package routes

import (
	"payment-collection/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("api/payment/create-intent", controllers.ProcessPayment)
	return r
}
