package types

// IntegerRequest represents a request for generating integers
type IntegerRequest struct {
	Min   int `json:"min" binding:"required"`
	Max   int `json:"max" binding:"required,gtfield=Min"`
	Count int `json:"count" binding:"required,gt=0"`
}

// StringRequest represents a request for generating strings
type StringRequest struct {
	Length int `json:"length" binding:"required,gt=0"`
	Count  int `json:"count" binding:"required,gt=0"`
}

// BooleanRequest represents a request for generating booleans
type BooleanRequest struct {
	Count int `json:"count" binding:"required,gt=0"`
}

// FloatRequest represents a request for generating floats
type FloatRequest struct {
	Min   float64 `json:"min" binding:"required"`
	Max   float64 `json:"max" binding:"required,gtfield=Min"`
	Count int     `json:"count" binding:"required,gt=0"`
}

// DateRequest represents a request for generating dates
type DateRequest struct {
	StartYear int `json:"start_year" binding:"required"`
	EndYear   int `json:"end_year" binding:"required,gtfield=StartYear"`
	Count     int `json:"count" binding:"required,gt=0"`
}

// EmailRequest represents a request for generating email addresses
type EmailRequest struct {
	Count int `json:"count" binding:"required,gt=0"`
}

// Complex requests
type AddressRequest struct {
	Count     int    `json:"count" binding:"required,gt=0"`
	Country   string `json:"country,omitempty"`   // Optional: filter by country
	State     string `json:"state,omitempty"`     // Optional: filter by state
}

type CreditCardRequest struct {
	Count    int    `json:"count" binding:"required,gt=0"`
	CardType string `json:"card_type,omitempty"` // Optional: specify card type
}

type CompanyRequest struct {
	Count    int    `json:"count" binding:"required,gt=0"`
	Industry string `json:"industry,omitempty"`  // Optional: filter by industry
}

// Batch generation request
type BatchRequest struct {
	Requests []GenerationRequest `json:"requests" binding:"required,dive"`
}

type GenerationRequest struct {
	Type    string      `json:"type" binding:"required"`
	Options interface{} `json:"options"`
}

// CountRequest is a simple request with just a count
type CountRequest struct {
	Count int `json:"count" binding:"required,gt=0"`
}

// PersonRequest represents a request for generating person data
type PersonRequest struct {
	Count    int    `json:"count" binding:"required,gt=0"`
	Gender   string `json:"gender,omitempty"` // male, female, or empty for any
	MinAge   int    `json:"min_age,omitempty"`
	MaxAge   int    `json:"max_age,omitempty"`
}

// ProductRequest represents a request for generating product data
type ProductRequest struct {
	Count     int     `json:"count" binding:"required,gt=0"`
	Category  string  `json:"category,omitempty"`
	MinPrice  float64 `json:"min_price,omitempty"`
	MaxPrice  float64 `json:"max_price,omitempty"`
}

// InternetRequest represents a request for generating internet data
type InternetRequest struct {
	Count        int    `json:"count" binding:"required,gt=0"`
	IncludeIPv6  bool   `json:"include_ipv6,omitempty"`
	Protocol     string `json:"protocol,omitempty"` // http, https
	DomainSuffix string `json:"domain_suffix,omitempty"`
}

// PaymentRequest represents a request for generating payment data
type PaymentRequest struct {
	Count     int    `json:"count" binding:"required,gt=0"`
	CardType  string `json:"card_type,omitempty"`  // visa, mastercard, etc.
	Currency  string `json:"currency,omitempty"`   // USD, EUR, etc.
}

// LocationRequest represents a request for generating location data
type LocationRequest struct {
	Count    int     `json:"count" binding:"required,gt=0"`
	Country  string  `json:"country,omitempty"`
	City     string  `json:"city,omitempty"`
	Latitude float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Radius   float64 `json:"radius,omitempty"`    // in kilometers
}

// TransactionRequest represents a request for generating transaction data
type TransactionRequest struct {
	Count    int     `json:"count" binding:"required,gt=0"`
	MinAmount float64 `json:"min_amount,omitempty"`
	MaxAmount float64 `json:"max_amount,omitempty"`
	Currency string  `json:"currency,omitempty"`

	FromAccount string `json:"from_account,omitempty"`
	ToAccount   string `json:"to_account,omitempty"`

	MinTimestamp string `json:"min_timestamp,omitempty"`
	MaxTimestamp string `json:"max_timestamp,omitempty"`

	Status string `json:"status,omitempty"`
	Type   string `json:"type,omitempty"`
}