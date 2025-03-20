package generator

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rushi/address-book-cli/internal/models"
)

// Generator generates random contact data for testing
type Generator struct {
	firstNames []string
	lastNames  []string
	domains    []string
}

// NewGenerator creates a new generator instance
func NewGenerator() *Generator {
	return &Generator{
		firstNames: []string{
			"John", "Jane", "Michael", "Sarah", "David", "Emma", "James", "Olivia",
			"William", "Sophia", "Robert", "Ava", "Joseph", "Isabella", "Thomas",
		},
		lastNames: []string{
			"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller",
			"Davis", "Rodriguez", "Martinez", "Hernandez", "Lopez", "Gonzalez",
			"Wilson", "Anderson",
		},
		domains: []string{
			"gmail.com", "yahoo.com", "hotmail.com", "outlook.com", "example.com",
		},
	}
}

// GenerateContact generates a random contact
func (g *Generator) GenerateContact() *models.Contact {
	firstName := g.firstNames[rand.Intn(len(g.firstNames))]
	lastName := g.lastNames[rand.Intn(len(g.lastNames))]
	domain := g.domains[rand.Intn(len(g.domains))]
	email := firstName + "." + lastName + "@" + domain
	phone := g.generatePhoneNumber()
	address := g.generateAddress()

	return models.NewContact(firstName, lastName, email, phone, address)
}

// GenerateContacts generates n random contacts
func (g *Generator) GenerateContacts(n int) []*models.Contact {
	contacts := make([]*models.Contact, n)
	for i := 0; i < n; i++ {
		contacts[i] = g.GenerateContact()
	}
	return contacts
}

// generatePhoneNumber generates a random phone number
func (g *Generator) generatePhoneNumber() string {
	digits := make([]byte, 10)
	for i := range digits {
		digits[i] = byte(rand.Intn(10) + '0')
	}
	return fmt.Sprintf("(%s) %s-%s",
		string(digits[:3]),
		string(digits[3:6]),
		string(digits[6:]))
}

// generateAddress generates a random address
func (g *Generator) generateAddress() string {
	streets := []string{
		"Main St", "Oak Ave", "Maple Dr", "Cedar Ln", "Pine Rd",
		"Elm St", "Washington Ave", "Lake Dr", "Hill Rd", "River Ln",
	}
	cities := []string{
		"New York", "Los Angeles", "Chicago", "Houston", "Phoenix",
		"Philadelphia", "San Antonio", "San Diego", "Dallas", "San Jose",
	}
	states := []string{
		"NY", "CA", "IL", "TX", "AZ", "PA", "TX", "CA", "TX", "CA",
	}

	street := streets[rand.Intn(len(streets))]
	number := rand.Intn(9999) + 1
	city := cities[rand.Intn(len(cities))]
	state := states[rand.Intn(len(states))]
	zip := rand.Intn(90000) + 10000

	return fmt.Sprintf("%d %s, %s, %s %d", number, street, city, state, zip)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
