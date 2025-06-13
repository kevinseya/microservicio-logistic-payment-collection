package main

import (
	"log"
	"net/http"
	"payment-collection/config"
	"payment-collection/routes"

	_ "payment-collection/docs" // Import Swagger documentation

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     //  Import with aliases
	ginSwagger "github.com/swaggo/gin-swagger" //  Import Swagger
)

// @title Payment Collection API
// @version 1.0
// @description API to handle payments with Stripe.
// @host localhost:8087
// @BasePath /api
var _ = swaggerFiles.Handler

func main() {
	config.InitConfig()
	config.InitDB()
	config.InitStripe()

	r := routes.SetupRouter()

	//  Redirect `/` to `/swagger/index.html`
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	// Add Swagger documentation to `/swagger`
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Payment microservice running on port 8087")
	r.Run(":8087")
}
