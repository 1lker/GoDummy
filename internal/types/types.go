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

// Response represents a generic response structure
type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
	Meta    *MetaData   `json:"meta,omitempty"`
}
