package generator

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
)

type Generator struct {
	rand *rand.Rand
	// Pre-defined data for generating realistic values
	domains     []string
	firstNames  []string
	lastNames   []string
	tlds        []string
	areaCodes   []string
}

func New() *Generator {
	source := rand.NewSource(time.Now().UnixNano())
	
	return &Generator{
		rand: rand.New(source),
		domains: []string{"gmail.com", "yahoo.com", "hotmail.com", "outlook.com", "example.com"},
		firstNames: []string{"James", "John", "Robert", "Michael", "William", "David", "Mary", "Patricia", "Jennifer", "Linda"},
		lastNames: []string{"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez"},
		tlds: []string{"com", "net", "org", "io", "dev"},
		areaCodes: []string{"201", "202", "212", "213", "301", "302", "303", "304", "305", "310"},
	}
}

// Basic type generators

func (g *Generator) GenerateInt(min, max int) int {
	return g.rand.Intn(max-min+1) + min
}

func (g *Generator) GenerateFloat(min, max float64) float64 {
	return min + g.rand.Float64()*(max-min)
}

func (g *Generator) GenerateString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[g.rand.Intn(len(charset))]
	}
	return string(result)
}

func (g *Generator) GenerateBool() bool {
	return g.rand.Intn(2) == 1
}

// Complex type generators

func (g *Generator) GenerateDate(startYear, endYear int) string {
	if startYear >= endYear {
		startYear, endYear = endYear, startYear
	}
	
	year := g.rand.Intn(endYear-startYear+1) + startYear
	month := g.rand.Intn(12) + 1
	
	// Calculate max days for the month
	maxDays := 31
	if month == 2 {
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			maxDays = 29
		} else {
			maxDays = 28
		}
	} else if month == 4 || month == 6 || month == 9 || month == 11 {
		maxDays = 30
	}
	
	day := g.rand.Intn(maxDays) + 1
	
	return fmt.Sprintf("%04d-%02d-%02d", year, month, day)
}

func (g *Generator) GenerateEmail() string {
	// Generate username part
	firstName := g.firstNames[g.rand.Intn(len(g.firstNames))]
	lastName := g.lastNames[g.rand.Intn(len(g.lastNames))]
	number := g.rand.Intn(100)
	
	// Get random domain
	domain := g.domains[g.rand.Intn(len(g.domains))]
	
	// Create email
	return fmt.Sprintf("%s.%s%d@%s", 
		firstName,
		lastName,
		number,
		domain,
	)
}

func (g *Generator) GeneratePhone() string {
	// Get random area code
	areaCode := g.areaCodes[g.rand.Intn(len(g.areaCodes))]
	
	// Generate remaining digits
	prefix := g.rand.Intn(900) + 100  // 100-999
	lineNum := g.rand.Intn(10000)     // 0000-9999
	
	return fmt.Sprintf("+1-%s-%03d-%04d", areaCode, prefix, lineNum)
}

func (g *Generator) GenerateName() string {
	firstName := g.firstNames[g.rand.Intn(len(g.firstNames))]
	lastName := g.lastNames[g.rand.Intn(len(g.lastNames))]
	return fmt.Sprintf("%s %s", firstName, lastName)
}

func (g *Generator) GenerateUsername() string {
	firstName := g.firstNames[g.rand.Intn(len(g.firstNames))]
	number := g.rand.Intn(1000)
	return fmt.Sprintf("%s%d", strings.ToLower(firstName), number)
}

func (g *Generator) GeneratePassword(length int) string {
	if length < 8 {
		length = 8 // Minimum password length
	}
	
	const (
		lowerChars = "abcdefghijklmnopqrstuvwxyz"
		upperChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		numbers    = "0123456789"
		special    = "!@#$%^&*"
	)
	
	// Ensure at least one of each type
	password := make([]byte, length)
	password[0] = lowerChars[g.rand.Intn(len(lowerChars))]
	password[1] = upperChars[g.rand.Intn(len(upperChars))]
	password[2] = numbers[g.rand.Intn(len(numbers))]
	password[3] = special[g.rand.Intn(len(special))]
	
	// Fill the rest randomly
	for i := 4; i < length; i++ {
		allChars := lowerChars + upperChars + numbers + special
		password[i] = allChars[g.rand.Intn(len(allChars))]
	}
	
	// Shuffle the password
	for i := len(password) - 1; i > 0; i-- {
		j := g.rand.Intn(i + 1)
		password[i], password[j] = password[j], password[i]
	}
	
	return string(password)
}

// Helper functions

func (g *Generator) randomChoice(items []string) string {
	return items[g.rand.Intn(len(items))]
}