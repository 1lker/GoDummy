package types

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