package generator

import (
	"fmt"
	"strings"
	"github.com/1lker/sd-gen-o2/internal/types"
	"github.com/brianvoe/gofakeit/v6"
)

// GeneratePerson generates a complete person profile
func (g *Generator) GeneratePerson() types.Person {
	gender := g.faker.Gender()
	firstName := g.faker.FirstName()
	lastName := g.faker.LastName()
	
	return types.Person{
		FirstName: firstName,
		LastName:  lastName,
		Email:     g.faker.Email(),
		Phone:     g.faker.Phone(),
		Gender:    gender,
		Age:       g.faker.IntRange(18, 80),
		Address:   g.GenerateAddress(),
		JobInfo:   g.GenerateJobDetails(),
	}
}

// GenerateJobDetails generates detailed job information
func (g *Generator) GenerateJobDetails() types.JobInfo {
	return types.JobInfo{
		Title:       g.faker.JobTitle(),
		Company:     g.faker.Company(),
		Level:       g.faker.JobLevel(),
		Description: g.faker.JobDescriptor(),
		Salary:      g.faker.IntRange(30000, 200000),
	}
}

// GenerateProductDetailed generates detailed product information
func (g *Generator) GenerateProductDetailed() types.Product {
	return types.Product{
		Name:        g.faker.ProductName(),
		Category:    g.faker.ProductCategory(),
		Description: g.faker.ProductDescription(),
		Price:       g.faker.Float64Range(1, 1000),
		SKU:         g.faker.UUID(),
		Barcode:     fmt.Sprintf("%d", g.faker.Int64()),  // Generate numeric barcode
		Brand:       g.faker.Company(),
		InStock:     g.faker.Bool(),
	}
}

// GenerateInternetDetailed generates detailed internet information
func (g *Generator) GenerateInternetDetailed() types.InternetInfo {
	browserList := []string{"Chrome", "Firefox", "Safari", "Edge", "Opera"}
	return types.InternetInfo{
		URL:       g.faker.URL(),
		IPv4:      g.faker.IPv4Address(),
		IPv6:      g.faker.IPv6Address(),
		Username:  g.faker.Username(),
		Domain:    g.faker.DomainName(),
		MAC:       g.faker.MacAddress(),
		UserAgent: g.faker.UserAgent(),
		Browser:   browserList[g.faker.IntRange(0, len(browserList)-1)],
	}
}

// GeneratePaymentInfo generates detailed payment information
func (g *Generator) GeneratePaymentInfo() types.PaymentInfo {
	return types.PaymentInfo{
		CreditCard: g.GenerateCreditCard(),
		Bitcoin:    g.faker.BitcoinAddress(),
		IBAN:      fmt.Sprintf("GB%02d%04d%06d%08d", 
			g.faker.IntRange(10, 99),
			g.faker.IntRange(1000, 9999),
			g.faker.IntRange(100000, 999999),
			g.faker.IntRange(10000000, 99999999)),
		BIC:       fmt.Sprintf("BKXX%04dX", g.faker.IntRange(1000, 9999)),
		Amount:    g.faker.Float64Range(10, 10000),
		Currency:  g.faker.CurrencyShort(),
	}
}

// GenerateLocation generates detailed location information
func (g *Generator) GenerateLocation() types.Location {
	lat := g.faker.Latitude()
	lon := g.faker.Longitude()
	
	return types.Location{
		Latitude:   lat,
		Longitude:  lon,
		City:      g.faker.City(),
		State:     g.faker.State(),
		Country:   g.faker.Country(),
		Timezone:  g.faker.TimeZone(),
		PostalCode: g.faker.Zip(),
	}
}

// GenerateVehicle generates detailed vehicle information
func (g *Generator) GenerateVehicle() types.Vehicle {
	return types.Vehicle{
		Make:         g.faker.CarMaker(),
		Model:        g.faker.CarModel(),
		Type:         g.faker.CarType(),
		Fuel:         g.faker.CarFuelType(),
		Transmission: g.faker.CarTransmissionType(),
		Year:         g.faker.IntRange(1990, 2024),
		Color:        g.faker.Color(),
		VIN:         generateVIN(g.faker),  // Custom VIN generator
		LicensePlate: generateLicensePlate(g.faker),  // Custom plate generator
	}
}

// GenerateFileInfo generates detailed file information
func (g *Generator) GenerateFileInfo() types.File {
	size := g.faker.IntRange(1000, 10000000)
	extension := generateFileExtension(g.faker)
	
	// MIME type mapping for common file extensions
	mimeTypes := map[string]string{
		"txt":  "text/plain",
		"pdf":  "application/pdf",
		"doc":  "application/msword",
		"docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		"xls":  "application/vnd.ms-excel",
		"xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		"jpg":  "image/jpeg",
		"png":  "image/png",
		"mp3":  "audio/mpeg",
		"mp4":  "video/mp4",
	}

	return types.File{
		Name:         generateFileName(g.faker) + "." + extension,
		Extension:    extension,
		Size:         int64(size),
		MimeType:     mimeTypes[extension],
		Path:         fmt.Sprintf("/path/to/%s.%s", generateFileName(g.faker), extension),
		ModifiedDate: g.faker.Date().Format("2006-01-02 15:04:05"),
	}
}


// Helper functions

func generateVIN(f *gofakeit.Faker) string {
	// VIN format: 17 characters
	chars := "ABCDEFGHJKLMNPRSTUVWXYZ0123456789"
	vin := make([]byte, 17)
	for i := range vin {
		vin[i] = chars[f.IntRange(0, len(chars)-1)]
	}
	return string(vin)
}

func generateLicensePlate(f *gofakeit.Faker) string {
	letters := "ABCDEFGHJKLMNPRSTUVWXYZ"
	numbers := "0123456789"
	
	plate := make([]byte, 7)
	// Format: ABC1234
	for i := 0; i < 3; i++ {
		plate[i] = letters[f.IntRange(0, len(letters)-1)]
	}
	for i := 3; i < 7; i++ {
		plate[i] = numbers[f.IntRange(0, len(numbers)-1)]
	}
	return string(plate)
}

func generateFileName(f *gofakeit.Faker) string {
	words := []string{
		f.Word(), f.Word(),
		fmt.Sprintf("%d", f.IntRange(1, 999)),
	}
	return strings.Join(words, "_")
}

func generateTransaction()

func generateFileExtension(f *gofakeit.Faker) string {
	extensions := []string{"txt", "pdf", "doc", "docx", "xls", "xlsx", "jpg", "png", "mp3", "mp4"}
	return extensions[f.IntRange(0, len(extensions)-1)]
}

// Helper functions for generating arrays of data
func GenerateMultiple[T any](count int, generator func() T) []T {
	results := make([]T, count)
	for i := 0; i < count; i++ {
		results[i] = generator()
	}
	return results
}