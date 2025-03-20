package models

import (
	"testing"
	"time"
)

func TestNewContact(t *testing.T) {
	firstName := "John"
	lastName := "Doe"
	email := "john@example.com"
	phone := "1234567890"
	address := "123 Main St"

	contact := NewContact(firstName, lastName, email, phone, address)

	if contact.FirstName != firstName {
		t.Errorf("Expected FirstName %s, got %s", firstName, contact.FirstName)
	}
	if contact.LastName != lastName {
		t.Errorf("Expected LastName %s, got %s", lastName, contact.LastName)
	}
	if contact.Email != email {
		t.Errorf("Expected Email %s, got %s", email, contact.Email)
	}
	if contact.Phone != phone {
		t.Errorf("Expected Phone %s, got %s", phone, contact.Phone)
	}
	if contact.Address != address {
		t.Errorf("Expected Address %s, got %s", address, contact.Address)
	}
	if contact.ID == "" {
		t.Error("Expected non-empty ID")
	}
	if contact.CreatedAt.IsZero() {
		t.Error("Expected non-zero CreatedAt")
	}
	if contact.UpdatedAt.IsZero() {
		t.Error("Expected non-zero UpdatedAt")
	}
}

func TestContactJSON(t *testing.T) {
	contact := NewContact("Jane", "Smith", "jane@example.com", "0987654321", "456 Oak St")

	// Test ToJSON
	jsonStr, err := contact.ToJSON()
	if err != nil {
		t.Errorf("ToJSON failed: %v", err)
	}

	// Test FromJSON
	parsedContact, err := FromJSON(jsonStr)
	if err != nil {
		t.Errorf("FromJSON failed: %v", err)
	}

	// Verify all fields match
	if parsedContact.FirstName != contact.FirstName ||
		parsedContact.LastName != contact.LastName ||
		parsedContact.Email != contact.Email ||
		parsedContact.Phone != contact.Phone ||
		parsedContact.Address != contact.Address ||
		parsedContact.ID != contact.ID {
		t.Error("Parsed contact does not match original contact")
	}

	// Test invalid JSON
	_, err = FromJSON("invalid json")
	if err == nil {
		t.Error("Expected error for invalid JSON")
	}
}

func TestGenerateID(t *testing.T) {
	id1 := generateID()
	time.Sleep(time.Millisecond)
	id2 := generateID()

	if id1 == "" || id2 == "" {
		t.Error("Generated IDs should not be empty")
	}
	if id1 == id2 {
		t.Error("Generated IDs should be unique")
	}
	if len(id1) != len(id2) {
		t.Error("Generated IDs should have consistent length")
	}
}
