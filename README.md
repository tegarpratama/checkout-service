# Project Installation

1. Clone the project
   `git clone https://github.com/tegarpratama/checkout-service`

2. Running docker compose
   `docker-compose up --build -d`

3. Migrate & seed data
   1. `docker exec -it ifortepay-checkout-service-app bash`
   2. `make migrate-up`
    3. `go run scripts/seeder/main.go `

# API Endpoints

- GET: `http://localhost:8080/api/check-health` - Check health system.

- POST: `http://localhost:8080/api/users` - Register a new user.
 payload:
     ```
     {
        "email": "user@gmail.com",
        "password": "password"
    }
     ```
     
- POST: `http://localhost:8080/api/transactions/checkouts` - Checkout products.
 payload:
     ```
     {
        "user_id": 1,
        "product_sku": ["43N23P", "234234"]
    }
     ```