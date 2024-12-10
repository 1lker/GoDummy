package types

// Address represents a physical address
type Address struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	ZipCode    string `json:"zip_code"`
	Country    string `json:"country"`
}

// CreditCard represents credit card information
type CreditCard struct {
	Number     string `json:"number"`
	Expiry     string `json:"expiry"`
	CVV        string `json:"cvv"`
	Type       string `json:"type"`
}

// Company represents company information
type Company struct {
	Name        string `json:"name"`
	Industry    string `json:"industry"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Website     string `json:"website"`
}

// Additional complex types

// SSN represents a Social Security Number
type SSN struct {
	Number        string `json:"number"`
	IsValid       bool   `json:"is_valid"`
	MaskedNumber  string `json:"masked_number"`
}

// ISBN represents an International Standard Book Number
type ISBN struct {
	ISBN10        string `json:"isbn_10"`
	ISBN13        string `json:"isbn_13"`
	IsValid       bool   `json:"is_valid"`
}

// UUID represents a Universally Unique Identifier
type UUID struct {
	Value     string `json:"value"`
	Version   int    `json:"version"`
}

// Export configurations
type ExportConfig struct {
	Format     string            `json:"format" binding:"required,oneof=csv json xml"`
	FilePath   string           `json:"file_path,omitempty"`
	Headers    []string         `json:"headers,omitempty"`
	Mapping    map[string]string `json:"mapping,omitempty"`
}

// ExportRequest represents a request to export generated data
type ExportRequest struct {
	GenerationRequest GenerationRequest `json:"generation_request"`
	ExportConfig     ExportConfig      `json:"export_config"`
}