package main

import (
	"log"
	"net/http"
	"payment-collection/config"
	"payment-collection/routes"

	_ "payment-collection/docs" // ✅ Importar la documentación de Swagger

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // ✅ Importar con alias
	ginSwagger "github.com/swaggo/gin-swagger" // ✅ Importar Swagger
)

// @title Payment Collection API
// @version 1.0
// @description API para manejar pagos con Stripe.
// @host localhost:8087
// @BasePath /api
var _ = swaggerFiles.Handler

func main() {
	config.InitConfig()
	config.InitDB()
	config.InitStripe()

	r := routes.SetupRouter()

	// 🔹 Redirigir `/` a `/swagger/index.html`
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	// 🔹 Agregar la documentación Swagger en `/swagger`
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Payment microservice running on port 8087")
	r.Run(":8087")
}
