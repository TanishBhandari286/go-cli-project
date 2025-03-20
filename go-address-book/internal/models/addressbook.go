package models

import (
	"errors"
	"strings"
	"sync"
	"time"
)

// AddressBook represents a collection of contacts
type AddressBook struct {
	contacts map[string]*Contact
	mu       sync.RWMutex
}

// NewAddressBook creates a new empty address book
func NewAddressBook() *AddressBook {
	return &AddressBook{
		contacts: make(map[string]*Contact),
	}
}

// AddContact adds a new contact to the address book
func (ab *AddressBook) AddContact(contact *Contact) error {
	ab.mu.Lock()
	defer ab.mu.Unlock()

	if _, exists := ab.contacts[contact.ID]; exists {
		return errors.New("contact with this ID already exists")
	}

	ab.contacts[contact.ID] = contact
	return nil
}

// GetContact retrieves a contact by ID
func (ab *AddressBook) GetContact(id string) (*Contact, error) {
	ab.mu.RLock()
	defer ab.mu.RUnlock()

	contact, exists := ab.contacts[id]
	if !exists {
		return nil, errors.New("contact not found")
	}

	return contact, nil
}

// UpdateContact updates an existing contact
func (ab *AddressBook) UpdateContact(contact *Contact) error {
	ab.mu.Lock()
	defer ab.mu.Unlock()

	if _, exists := ab.contacts[contact.ID]; !exists {
		return errors.New("contact not found")
	}

	contact.UpdatedAt = time.Now()
	ab.contacts[contact.ID] = contact
	return nil
}

// DeleteContact removes a contact by ID
func (ab *AddressBook) DeleteContact(id string) error {
	ab.mu.Lock()
	defer ab.mu.Unlock()

	if _, exists := ab.contacts[id]; !exists {
		return errors.New("contact not found")
	}

	delete(ab.contacts, id)
	return nil
}

// GetAllContacts returns all contacts in the address book
func (ab *AddressBook) GetAllContacts() []*Contact {
	ab.mu.RLock()
	defer ab.mu.RUnlock()

	contacts := make([]*Contact, 0, len(ab.contacts))
	for _, contact := range ab.contacts {
		contacts = append(contacts, contact)
	}
	return contacts
}

// SearchContacts searches for contacts by name or email
func (ab *AddressBook) SearchContacts(query string) []*Contact {
	ab.mu.RLock()
	defer ab.mu.RUnlock()

	query = strings.ToLower(query)
	var results []*Contact

	for _, contact := range ab.contacts {
		if strings.Contains(strings.ToLower(contact.FirstName), query) ||
			strings.Contains(strings.ToLower(contact.LastName), query) ||
			strings.Contains(strings.ToLower(contact.Email), query) {
			results = append(results, contact)
		}
	}

	return results
}
