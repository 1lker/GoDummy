package generator

import (
	"fmt"
	"strings"
	"time"
	"github.com/1lker/sd-gen-o2/internal/types"
)

// GenerateAddress generates a random address
func (g *Generator) GenerateAddress() types.Address {
	// Additional data for realistic addresses
	streetTypes := []string{"Street", "Avenue", "Boulevard", "Road", "Lane", "Drive"}
	cities := []string{"New York", "Los Angeles", "Chicago", "Houston", "Phoenix", "Philadelphia"}
	states := map[string]string{
		"NY": "New York",
		"CA": "California",
		"IL": "Illinois",
		"TX": "Texas",
		"AZ": "Arizona",
		"PA": "Pennsylvania",
	}

	// Generate street number and name
	streetNum := g.GenerateInt(1, 9999)
	streetName := g.GenerateName()
	streetType := streetTypes[g.rand.Intn(len(streetTypes))]

	// Select random city and state
	city := cities[g.rand.Intn(len(cities))]
	stateAbbr := g.randomChoice(getKeys(states))

	// Generate ZIP code
	zipCode := fmt.Sprintf("%05d", g.GenerateInt(10000, 99999))

	return types.Address{
		Street:  fmt.Sprintf("%d %s %s", streetNum, streetName, streetType),
		City:    city,
		State:   stateAbbr,
		ZipCode: zipCode,
		Country: "USA",
	}
}

// GenerateCreditCard generates a random credit card
func (g *Generator) GenerateCreditCard() types.CreditCard {
	cardTypes := map[string]string{
		"VISA":       "4",
		"MASTERCARD": "5",
		"AMEX":       "37",
		"DISCOVER":   "6011",
	}

	// Select random card type
	cardType := g.randomChoice(getKeys(cardTypes))
	prefix := cardTypes[cardType]

	// Generate card number
	length := 16
	if cardType == "AMEX" {
		length = 15
	}

	number := prefix
	remainingLength := length - len(prefix)
	for i := 0; i < remainingLength; i++ {
		number += fmt.Sprintf("%d", g.rand.Intn(10))
	}

	// Generate expiry date (between now and 4 years from now)
	currentYear := time.Now().Year() % 100
	year := g.GenerateInt(currentYear, currentYear+4)
	month := g.GenerateInt(1, 12)
	expiry := fmt.Sprintf("%02d/%02d", month, year)

	// Generate CVV
	cvvLength := 3
	if cardType == "AMEX" {
		cvvLength = 4
	}
	cvv := ""
	for i := 0; i < cvvLength; i++ {
		cvv += fmt.Sprintf("%d", g.rand.Intn(10))
	}

	return types.CreditCard{
		Number: g.formatCardNumber(number, cardType),
		Expiry: expiry,
		CVV:    cvv,
		Type:   cardType,
	}
}

// GenerateCompany generates a random company
func (g *Generator) GenerateCompany() types.Company {
	industries := []string{
		"Technology", "Healthcare", "Finance", "Manufacturing",
		"Retail", "Education", "Media", "Construction",
	}
	
	companyTypes := []string{
		"LLC", "Inc.", "Corp.", "Ltd.", "Group",
	}

	// Generate company name components
	prefix := g.lastNames[g.rand.Intn(len(g.lastNames))]
	industry := industries[g.rand.Intn(len(industries))]
	companyType := companyTypes[g.rand.Intn(len(companyTypes))]

	// Construct company name
	var name string
	switch g.rand.Intn(3) {
	case 0:
		name = fmt.Sprintf("%s %s", prefix, companyType)
	case 1:
		name = fmt.Sprintf("%s %s %s", prefix, industry, companyType)
	case 2:
		suffixes := []string{"Tech", "Systems", "Solutions", "Global", "International"}
		suffix := suffixes[g.rand.Intn(len(suffixes))]
		name = fmt.Sprintf("%s %s %s", prefix, suffix, companyType)
	}

	// Generate website
	website := fmt.Sprintf("www.%s.com", 
		strings.ToLower(strings.ReplaceAll(
			strings.ReplaceAll(name, " ", ""),
			".", "",
		)),
	)

	// Generate description
	descriptions := []string{
		"Leading provider of %s solutions",
		"Innovative %s company",
		"Global leader in %s services",
		"Premier %s solutions provider",
	}
	description := fmt.Sprintf(
		descriptions[g.rand.Intn(len(descriptions))],
		strings.ToLower(industry),
	)

	return types.Company{
		Name:        name,
		Industry:    industry,
		Type:        companyType,
		Description: description,
		Website:     website,
	}
}

// Helper functions

func (g *Generator) formatCardNumber(number string, cardType string) string {
	if cardType == "AMEX" {
		return fmt.Sprintf("%s-%s-%s-%s",
			number[0:4],
			number[4:10],
			number[10:15],
			"",
		)
	}
	return fmt.Sprintf("%s-%s-%s-%s",
		number[0:4],
		number[4:8],
		number[8:12],
		number[12:16],
	)
}

func getKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}