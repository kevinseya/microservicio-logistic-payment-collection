package config

import (
	"database/sql"
	"log"
	"os"
	"strings"

	"payment-collection/models"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/stripe/stripe-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var StripePublicKey string

func InitConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: No .env file found or cannot be loaded")
	}
}

func createDatabaseIfNotExists() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN is not set in .env file")
	}

	parts := strings.Split(dsn, " ")
	var dbName string
	for _, part := range parts {
		if strings.HasPrefix(part, "dbname=") {
			dbName = strings.TrimPrefix(part, "dbname=")
			break
		}
	}

	if dbName == "" {
		log.Fatal("Could not extract database name from DB_DSN")
	}

	baseDSN := strings.Replace(dsn, "dbname="+dbName, "dbname=postgres", 1)
	conn, err := sql.Open("postgres", baseDSN)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL for database creation:", err)
	}
	defer conn.Close()

	_, err = conn.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		log.Println("Database may already exist or could not be created:", err)
	}
}

func InitDB() {
	createDatabaseIfNotExists()

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN is not set in .env file")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Connected to database successfully.")

	// Auto Migration of the Payment table
	err = DB.AutoMigrate(&models.Payment{})
	if err != nil {
		log.Fatal("Failed to migrate database schema:", err)
	}

	log.Println("Database migration completed successfully.")
}

func InitStripe() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	StripePublicKey = os.Getenv("STRIPE_PUBLIC_KEY")

	if stripe.Key == "" {
		log.Fatal("STRIPE_SECRET_KEY is not set in .env file")
	}

	if StripePublicKey == "" {
		log.Println("Warning: STRIPE_PUBLIC_KEY is not set in .env file")
	}
}
