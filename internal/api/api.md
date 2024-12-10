API Structure:
└── internal/
    ├── api/              # API related code
    │   ├── server.go     # HTTP server setup
    │   ├── handlers.go   # Request handlers
    │   └── routes.go     # API route definitions
    └── types/            # Shared types/models
        └── types.go      # Request/Response structures


BASIC API ENDPOINTS
POST /api/v1/generate/int
Request:
{
    "min": 1,
    "max": 100,
    "count": 5
}
Response:
{
    "data": [45, 67, 23, 89, 12]
}

POST /api/v1/generate/string
Request:
{
    "length": 10,
    "count": 3
}
Response:
{
    "data": ["aB3kM9pQ2R", "xY7nL4jK8S", "pO2wQ9mN5T"]
}

* Technologies we'll use:

gin-gonic/gin: For HTTP routing and handling
go-playground/validator: For request validation

- Flow of an API request:

CopyClient Request → Router → Handler → Generator → Response
   (JSON)            (Gin)    (Parse)   (Generate)  (JSON)