package models

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"time"
)

// Contact represents a person in the address book
type Contact struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// NewContact creates a new contact with the given information
func NewContact(firstName, lastName, email, phone, address string) *Contact {
	now := time.Now()
	return &Contact{
		ID:        generateID(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
		Address:   address,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// ToJSON converts the contact to a JSON string
func (c *Contact) ToJSON() (string, error) {
	bytes, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// FromJSON creates a contact from a JSON string
func FromJSON(jsonStr string) (*Contact, error) {
	var contact Contact
	err := json.Unmarshal([]byte(jsonStr), &contact)
	if err != nil {
		return nil, err
	}
	return &contact, nil
}

// generateID creates a unique ID for the contact
func generateID() string {
	timestamp := time.Now().Format("20060102150405")
	randomBytes := make([]byte, 8)
	rand.Read(randomBytes)
	return timestamp + "-" + hex.EncodeToString(randomBytes)
}
