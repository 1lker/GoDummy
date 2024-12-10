package generator

import (
	"fmt"
	"time"
	"strings"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/1lker/sd-gen-o2/internal/types"
)

type Generator struct {
    faker *gofakeit.Faker
}

func New() *Generator {
    return &Generator{
        faker: gofakeit.New(time.Now().UnixNano()),
    }
}

// Basic type generators
func (g *Generator) GenerateInt(min, max int) int {
	return g.faker.IntRange(min, max)
}

func (g *Generator) GenerateFloat(min, max float64) float64 {
	return g.faker.Float64Range(min, max)
}

func (g *Generator) GenerateString(length int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)
	for i := range result {
		result[i] = letters[g.faker.IntRange(0, len(letters)-1)]
	}
	return string(result)
}

func (g *Generator) GenerateBool() bool {
	return g.faker.Bool()
}

// Complex type generators
func (g *Generator) GenerateAddress() types.Address {
	return types.Address{
		Street:  fmt.Sprintf("%d %s %s", g.faker.IntRange(100, 9999), g.faker.LastName(), "Street"),
		City:    g.faker.City(),
		State:   g.faker.State(),
		ZipCode: fmt.Sprintf("%05d", g.faker.IntRange(10000, 99999)),
		Country: g.faker.Country(),
	}
}

func (g *Generator) GenerateCreditCard() types.CreditCard {
	ccTypes := []string{"VISA", "MASTERCARD", "AMEX", "DISCOVER"}
	cardType := ccTypes[g.faker.IntRange(0, len(ccTypes)-1)]
	
	return types.CreditCard{
		Number: generateCCNumber(g.faker, cardType),
		Expiry: fmt.Sprintf("%02d/%d", g.faker.IntRange(1, 12), g.faker.IntRange(23, 28)),
		CVV:    fmt.Sprintf("%03d", g.faker.IntRange(100, 999)),
		Type:   cardType,
	}
}

func (g *Generator) GenerateCompany() types.Company {
	industries := []string{"Technology", "Healthcare", "Finance", "Manufacturing", "Retail", "Education"}
	companyTypes := []string{"Corp.", "Inc.", "LLC", "Ltd."}
	
	return types.Company{
		Name:        fmt.Sprintf("%s %s", g.faker.LastName(), companyTypes[g.faker.IntRange(0, len(companyTypes)-1)]),
		Industry:    industries[g.faker.IntRange(0, len(industries)-1)],
		Type:        companyTypes[g.faker.IntRange(0, len(companyTypes)-1)],
		Description: g.faker.Sentence(10),
		Website:     g.faker.URL(),
	}
}

func (g *Generator) GenerateName() string {
	return fmt.Sprintf("%s %s", g.faker.FirstName(), g.faker.LastName())
}

func (g *Generator) GenerateEmail() string {
	return g.faker.Email()
}

func (g *Generator) GeneratePhone() string {
	return g.faker.Phone()
}

func (g *Generator) GenerateDate(startYear, endYear int) string {
	start := time.Date(startYear, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(endYear, 12, 31, 23, 59, 59, 0, time.UTC)
	return g.faker.DateRange(start, end).Format("2006-01-02")
}

// Additional data types
func (g *Generator) GenerateSSN() types.SSN {
	ssn := fmt.Sprintf("%03d-%02d-%04d", 
		g.faker.IntRange(100, 999),
		g.faker.IntRange(10, 99),
		g.faker.IntRange(1000, 9999))
	
	return types.SSN{
		Number:       ssn,
		IsValid:     true,
		MaskedNumber: "xxx-xx-" + ssn[7:],
	}
}

func (g *Generator) GenerateISBN() types.ISBN {
	// Generate ISBN-10
	isbn10 := generateISBN10(g.faker)
	// Generate ISBN-13
	isbn13 := generateISBN13(g.faker)
	
	return types.ISBN{
		ISBN10:  isbn10,
		ISBN13:  isbn13,
		IsValid: true,
	}
}

func (g *Generator) GenerateUsername() string {
	return g.faker.Username()
}

func (g *Generator) GeneratePassword(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[g.faker.IntRange(0, len(chars)-1)]
	}
	return string(result)
}

// Additional useful generators
func (g *Generator) GenerateJobInfo() map[string]string {
	return map[string]string{
		"title":       g.faker.JobTitle(),
		"description": g.faker.Sentence(15),
		"level":       g.faker.RandomString([]string{"Entry", "Junior", "Senior", "Lead", "Manager"}),
		"company":     g.GenerateCompany().Name,
	}
}

func (g *Generator) GenerateProduct() map[string]interface{} {
	return map[string]interface{}{
		"name":        fmt.Sprintf("%s %s", g.faker.Word(), g.faker.Word()),
		"category":    g.faker.RandomString([]string{"Electronics", "Clothing", "Food", "Books", "Tools"}),
		"description": g.faker.Sentence(10),
		"price":       g.faker.Float64Range(1, 1000),
		"sku":         g.faker.UUID(),
	}
}

func (g *Generator) GenerateCarInfo() map[string]string {
	carMakers := []string{"Toyota", "Honda", "Ford", "BMW", "Mercedes", "Audi"}
	fuelTypes := []string{"Gasoline", "Diesel", "Electric", "Hybrid"}
	transmissions := []string{"Automatic", "Manual", "CVT"}
	
	return map[string]string{
		"make":         carMakers[g.faker.IntRange(0, len(carMakers)-1)],
		"model":        g.faker.Word(),
		"year":         fmt.Sprintf("%d", g.faker.IntRange(1990, 2024)),
		"fuel":         fuelTypes[g.faker.IntRange(0, len(fuelTypes)-1)],
		"transmission": transmissions[g.faker.IntRange(0, len(transmissions)-1)],
	}
}

func (g *Generator) GenerateInternetInfo() map[string]string {
	return map[string]string{
		"url":      g.faker.URL(),
		"ipv4":     g.faker.IPv4Address(),
		"ipv6":     g.faker.IPv6Address(),
		"username": g.faker.Username(),
		"domain":   g.faker.DomainName(),
		"mac":      generateMACAddress(g.faker),
	}
}

func (g *Generator) GenerateFile() map[string]string {
	extensions := []string{".txt", ".pdf", ".doc", ".jpg", ".png", ".mp3"}
	mimeTypes := map[string]string{
		".txt": "text/plain",
		".pdf": "application/pdf",
		".doc": "application/msword",
		".jpg": "image/jpeg",
		".png": "image/png",
		".mp3": "audio/mpeg",
	}
	
	ext := extensions[g.faker.IntRange(0, len(extensions)-1)]
	return map[string]string{
		"name":      fmt.Sprintf("%s%s", g.faker.Word(), ext),
		"extension": ext,
		"mime":      mimeTypes[ext],
		"size":      fmt.Sprintf("%dKB", g.faker.IntRange(1, 10000)),
	}
}

func (g *Generator) GenerateColor() map[string]string {
	colors := []string{"red", "blue", "green", "yellow", "purple", "orange"}
	hexColors := []string{"#FF0000", "#0000FF", "#00FF00", "#FFFF00", "#800080", "#FFA500"}
	rgbColors := []string{"rgb(255,0,0)", "rgb(0,0,255)", "rgb(0,255,0)", "rgb(255,255,0)", "rgb(128,0,128)", "rgb(255,165,0)"}
	
	index := g.faker.IntRange(0, len(colors)-1)
	return map[string]string{
		"name": colors[index],
		"hex":  hexColors[index],
		"rgb":  rgbColors[index],
	}
}

// Helper functions
func generateCCNumber(f *gofakeit.Faker, cardType string) string {
	prefixes := map[string]string{
		"VISA": "4",
		"MASTERCARD": "5",
		"AMEX": "37",
		"DISCOVER": "6011",
	}
	
	prefix := prefixes[cardType]
	length := 16
	if cardType == "AMEX" {
		length = 15
	}
	
	number := prefix
	for i := len(prefix); i < length; i++ {
		number += fmt.Sprintf("%d", f.IntRange(0, 9))
	}
	
	return number
}

func generateMACAddress(f *gofakeit.Faker) string {
	mac := make([]string, 6)
	for i := range mac {
		mac[i] = fmt.Sprintf("%02X", f.IntRange(0, 255))
	}
	return strings.Join(mac, ":")
}

func generateISBN10(f *gofakeit.Faker) string {
	digits := make([]int, 9)
	for i := range digits {
		digits[i] = f.IntRange(0, 9)
	}
	
	// Calculate check digit
	sum := 0
	for i := 0; i < 9; i++ {
		sum += digits[i] * (10 - i)
	}
	checkDigit := (11 - (sum % 11)) % 11
	
	isbn := ""
	for _, d := range digits {
		isbn += fmt.Sprintf("%d", d)
	}
	if checkDigit == 10 {
		isbn += "X"
	} else {
		isbn += fmt.Sprintf("%d", checkDigit)
	}
	
	return isbn
}

func generateISBN13(f *gofakeit.Faker) string {
	prefix := "978"
	digits := make([]int, 9)
	for i := range digits {
		digits[i] = f.IntRange(0, 9)
	}
	
	// Calculate check digit
	sum := 0
	weights := []int{1, 3}
	
	// Add prefix digits
	for i, d := range prefix {
		digit := int(d - '0')
		sum += digit * weights[i%2]
	}
	
	// Add remaining digits
	for i, d := range digits {
		sum += d * weights[(i+3)%2]
	}
	
	checkDigit := (10 - (sum % 10)) % 10
	
	isbn := prefix
	for _, d := range digits {
		isbn += fmt.Sprintf("%d", d)
	}
	isbn += fmt.Sprintf("%d", checkDigit)
	
	return isbn
}

// GenerateMultiple generates multiple instances of the given generator function
func (g *Generator) GenerateMultiple(count int, generator func() interface{}) []interface{} {
    results := make([]interface{}, count)
    for i := 0; i < count; i++ {
        results[i] = generator()
    }
    return results
}