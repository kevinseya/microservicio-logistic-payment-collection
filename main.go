package main

import (
	"log"
	"payment-collection/config"
	"payment-collection/routes"
)

func main() {
	config.InitConfig()
	config.InitDB()
	config.InitStripe()
	r := routes.SetupRouter()
	log.Println("Payment microservice running on port 8080")
	r.Run(":8087")
}
