package types

// Person represents a generated person
type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Gender    string `json:"gender"`
	Age       int    `json:"age"`
	Address   Address `json:"address"`
	JobInfo   JobInfo `json:"job"`
}

// JobInfo represents job information
type JobInfo struct {
	Title       string `json:"title"`
	Company     string `json:"company"`
	Level       string `json:"level"`
	Description string `json:"description"`
	Salary      int    `json:"salary"`
}

// Product represents a product
type Product struct {
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	SKU         string  `json:"sku"`
	Barcode     string  `json:"barcode"`
	Brand       string  `json:"brand"`
	InStock     bool    `json:"in_stock"`
}

// InternetInfo represents internet-related information
type InternetInfo struct {
	URL          string `json:"url"`
	IPv4         string `json:"ipv4"`
	IPv6         string `json:"ipv6,omitempty"`
	Username     string `json:"username"`
	Domain       string `json:"domain"`
	MAC          string `json:"mac"`
	UserAgent    string `json:"user_agent"`
	Browser      string `json:"browser"`
}

// PaymentInfo represents payment information
type PaymentInfo struct {
	CreditCard CreditCard `json:"credit_card"`
	Bitcoin    string     `json:"bitcoin"`
	IBAN       string     `json:"iban"`
	BIC        string     `json:"bic"`
	Amount     float64    `json:"amount"`
	Currency   string     `json:"currency"`
}

// Location represents a geographic location
type Location struct {
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	City       string  `json:"city"`
	State      string  `json:"state"`
	Country    string  `json:"country"`
	Timezone   string  `json:"timezone"`
	PostalCode string  `json:"postal_code"`
}

// Vehicle represents a vehicle
type Vehicle struct {
	Make         string `json:"make"`
	Model        string `json:"model"`
	Type         string `json:"type"`
	Fuel         string `json:"fuel"`
	Transmission string `json:"transmission"`
	Year         int    `json:"year"`
	Color        string `json:"color"`
	VIN         string `json:"vin"`
	LicensePlate string `json:"license_plate"`
}

// File represents a file
type File struct {
	Name          string `json:"name"`
	Extension     string `json:"extension"`
	Size          int64  `json:"size"`
	MimeType      string `json:"mime_type"`
	Path          string `json:"path"`
	ModifiedDate  string `json:"modified_date"`
}

// MetaData represents metadata about the generated data
type MetaData struct {
	Count       int         `json:"count"`
	TimeStamp   string      `json:"timestamp"`
	Types       []string    `json:"types,omitempty"`
	Parameters  interface{} `json:"parameters,omitempty"`
	Generation  string      `json:"generation_time"`
}