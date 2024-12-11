# integer generation API test
curl -X POST http://localhost:8080/api/v1/generate/int \
  -H "Content-Type: application/json" \
  -d '{"min": 1, "max": 100, "count": 5}'



# string generation API test
curl -X POST http://localhost:8080/api/v1/generate/string \
  -H "Content-Type: application/json" \
  -d '{"length": 10, "count": 3}'



# boolean generation API test
curl -X POST http://localhost:8080/api/v1/generate/bool \
  -H "Content-Type: application/json" \
  -d '{"count": 5}'


# float generation API test
curl -X POST http://localhost:8080/api/v1/generate/float \
  -H "Content-Type: application/json" \
  -d '{"min": 0.0, "max": 1.0, "count": 5}'


# date generation API test
curl -X POST http://localhost:8080/api/v1/generate/date \
  -H "Content-Type: application/json" \
  -d '{"start_year": 2000, "end_year": 2024, "count": 3}'



# email generation API test
curl -X POST http://localhost:8080/api/v1/generate/email \
  -H "Content-Type: application/json" \
  -d '{"count": 3}'

# health check API test
curl http://localhost:8080/health



# address generation API test
curl -X POST http://localhost:8080/api/v1/generate/address \
  -H "Content-Type: application/json" \
  -d '{
    "count": 2,
    "country": "USA",
    "state": "CA"
  }'


# credit card generation API test
curl -X POST http://localhost:8080/api/v1/generate/creditcard \
  -H "Content-Type: application/json" \
  -d '{
    "count": 2,
    "card_type": "VISA"
  }'


# company generation API test
curl -X POST http://localhost:8080/api/v1/generate/company \
  -H "Content-Type: application/json" \
  -d '{
    "count": 2,
    "industry": "Technology"
  }'



# batch generation API test
curl -X POST http://localhost:8080/api/v1/generate/batch \
  -H "Content-Type: application/json" \
  -d '{
    "requests": [
      {
        "type": "address",
        "options": {
          "count": 2,
          "country": "USA"
        }
      },
      {
        "type": "company",
        "options": {
          "count": 1,
          "industry": "Technology"
        }
      }
    ]
  }'



curl -X POST http://localhost:8080/api/v1/generate/person \
  -H "Content-Type: application/json" \
  -d '{
    "count": 2,
    "gender": "female",
    "min_age": 20,
    "max_age": 30
  }'


curl --location 'http://localhost:8080/api/v1/generate/creditcard' \
--header 'Content-Type: application/json' \
--data '{
    "count": 2,
    "card_type": "VISA"
  }'



curl --location 'http://localhost:8080/api/v1/generate/transaction/batch' \
--header 'Content-Type: application/json' \
--data '{
    "batch_count": 2,
    "batch_size": 3,
    "base_request": {
      "min_amount": 1000,
      "max_amount": 5000,
      "currency": "USD",
      "type": "transfer"
    }
  }'



curl --location 'http://localhost:8080/api/v1/generate/transaction' \
--header 'Content-Type: application/json' \
--data '{
    "count": 5,
    "min_amount": 100,
    "max_amount": 1000,
    "currency": "USD",
    "status": "completed",
    "type": "payment"
  }'

