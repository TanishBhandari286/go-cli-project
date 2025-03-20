package models

import (
	"sync"
	"testing"
)

func TestAddressBookOperations(t *testing.T) {
	ab := NewAddressBook()

	// Test adding a contact
	contact := NewContact("John", "Doe", "john@example.com", "1234567890", "123 Main St")
	err := ab.AddContact(contact)
	if err != nil {
		t.Errorf("Failed to add contact: %v", err)
	}

	// Test adding duplicate contact
	err = ab.AddContact(contact)
	if err == nil {
		t.Error("Expected error when adding duplicate contact")
	}

	// Test getting a contact
	retrieved, err := ab.GetContact(contact.ID)
	if err != nil {
		t.Errorf("Failed to get contact: %v", err)
	}
	if retrieved.ID != contact.ID {
		t.Error("Retrieved contact does not match original")
	}

	// Test getting non-existent contact
	_, err = ab.GetContact("non-existent-id")
	if err == nil {
		t.Error("Expected error when getting non-existent contact")
	}

	// Test updating a contact
	contact.FirstName = "Jane"
	err = ab.UpdateContact(contact)
	if err != nil {
		t.Errorf("Failed to update contact: %v", err)
	}

	// Verify update
	updated, _ := ab.GetContact(contact.ID)
	if updated.FirstName != "Jane" {
		t.Error("Contact update failed")
	}
	if !updated.UpdatedAt.After(updated.CreatedAt) {
		t.Error("UpdatedAt should be after CreatedAt")
	}

	// Test updating non-existent contact
	nonExistentContact := NewContact("Non", "Existent", "none@example.com", "0000000000", "Nowhere")
	err = ab.UpdateContact(nonExistentContact)
	if err == nil {
		t.Error("Expected error when updating non-existent contact")
	}

	// Test getting all contacts
	allContacts := ab.GetAllContacts()
	if len(allContacts) != 1 {
		t.Errorf("Expected 1 contact, got %d", len(allContacts))
	}

	// Test searching contacts
	searchResults := ab.SearchContacts("Jane")
	if len(searchResults) != 1 {
		t.Errorf("Expected 1 search result, got %d", len(searchResults))
	}
	searchResults = ab.SearchContacts("john@example.com")
	if len(searchResults) != 1 {
		t.Errorf("Expected 1 search result, got %d", len(searchResults))
	}
	searchResults = ab.SearchContacts("nonexistent")
	if len(searchResults) != 0 {
		t.Errorf("Expected 0 search results, got %d", len(searchResults))
	}

	// Test deleting a contact
	err = ab.DeleteContact(contact.ID)
	if err != nil {
		t.Errorf("Failed to delete contact: %v", err)
	}

	// Verify deletion
	allContacts = ab.GetAllContacts()
	if len(allContacts) != 0 {
		t.Error("Contact was not deleted")
	}

	// Test deleting non-existent contact
	err = ab.DeleteContact("non-existent-id")
	if err == nil {
		t.Error("Expected error when deleting non-existent contact")
	}
}

func TestConcurrentOperations(t *testing.T) {
	ab := NewAddressBook()
	var wg sync.WaitGroup
	wg.Add(2)

	// Test concurrent reads and writes
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			contact := NewContact("Test", "User", "test@example.com", "1234567890", "Test Address")
			if err := ab.AddContact(contact); err != nil {
				t.Errorf("Failed to add contact: %v", err)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			_ = ab.GetAllContacts()
			_ = ab.SearchContacts("Test")
		}
	}()

	// Wait for both goroutines to complete
	wg.Wait()

	// Verify final state
	contacts := ab.GetAllContacts()
	if len(contacts) != 100 {
		t.Errorf("Expected 100 contacts, got %d", len(contacts))
	}
}
