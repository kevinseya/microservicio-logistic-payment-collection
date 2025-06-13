# microservicio-logistic-payment-collection

## 📌 Description

This microservice is developed in Go and is responsible for managing payment collection on the logistics platform, integrating the Stripe API to process transactions.

## 📂 Project Structure

- config/ - System configuration.

- controllers/payment_controller.go - Controller that manages payment operations.

- models/payment.go - Data model definition for payments.

- repositories/payment_repo.go - Persistence layer for payment operations.

- routes/routes.go - Microservice route definitions.

- services/payment_service.go - Business logic related to payments.

- websocket/websocket.go - Module for real-time communication (optional).

- main.go - Main entry point of the microservice.

- go.mod and go.sum - Dependency management in Go.

- .gitignore - Files and folders ignored by Git.

- README.md - Project documentation.

## 🛠 Requirements

- Go 1.18 or higher.

- Stripe API Key.

- Docker (optional, for containerized execution).

## 🚀 Installation

Clone the repository:
```bash
git clone https://github.com/kevinseya/microservice-logistic-payment-collection.git
```

Navigate to the project directory:
```bash
cd microservice-logistic-payment-collection
```

Install dependencies:
```bash
go mod tidy
```

Set up environment variables:
Create a .env file in the root directory with the following content:
```bash
STRIPE_SECRET_KEY=your_secret_key
DB_HOST=localhost
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=logistic_db
```

Run the server:
```bash
go run main.go
```
The microservice will run on http://localhost:8080

📌 Response Codes

- 200 OK - Successful operation.

- 400 Bad Request - Incorrect or missing data.

- 500 Internal Server Error - Server error.

