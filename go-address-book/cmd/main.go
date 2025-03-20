package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/rushi/address-book-cli/internal/config"
	"github.com/rushi/address-book-cli/internal/generator"
	"github.com/rushi/address-book-cli/internal/models"
	"github.com/rushi/address-book-cli/internal/storage"
)

func main() {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	store := storage.NewCSVStorage(cfg.CSVPath)

	addressBook, err := store.Load()
	if err != nil {
		fmt.Printf("Error loading address book: %v\n", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nAddress Book CLI")
		fmt.Println("1. Add Contact")
		fmt.Println("2. List Contacts")
		fmt.Println("3. Search Contacts")
		fmt.Println("4. Update Contact")
		fmt.Println("5. Delete Contact")
		fmt.Println("6. Generate Test Data")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice (1-7): ")

		if !scanner.Scan() {
			break
		}
		choice := scanner.Text()

		switch choice {
		case "1":
			addContact(scanner, addressBook)
		case "2":
			listContacts(addressBook)
		case "3":
			searchContacts(scanner, addressBook)
		case "4":
			updateContact(scanner, addressBook)
		case "5":
			deleteContact(scanner, addressBook)
		case "6":
			generateTestData(addressBook)
		case "7":
			if err := store.Save(addressBook); err != nil {
				fmt.Printf("Error saving address book: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func addContact(scanner *bufio.Scanner, addressBook *models.AddressBook) {
	fmt.Print("Enter first name: ")
	scanner.Scan()
	firstName := scanner.Text()

	fmt.Print("Enter last name: ")
	scanner.Scan()
	lastName := scanner.Text()

	fmt.Print("Enter email: ")
	scanner.Scan()
	email := scanner.Text()

	fmt.Print("Enter phone: ")
	scanner.Scan()
	phone := scanner.Text()

	fmt.Print("Enter address: ")
	scanner.Scan()
	address := scanner.Text()

	contact := models.NewContact(firstName, lastName, email, phone, address)
	if err := addressBook.AddContact(contact); err != nil {
		fmt.Printf("Error adding contact: %v\n", err)
		return
	}

	fmt.Println("Contact added successfully!")
}

func listContacts(addressBook *models.AddressBook) {
	contacts := addressBook.GetAllContacts()
	if len(contacts) == 0 {
		fmt.Println("No contacts found.")
		return
	}

	fmt.Println("\nContacts:")
	for _, contact := range contacts {
		printContact(contact)
	}
}

func searchContacts(scanner *bufio.Scanner, addressBook *models.AddressBook) {
	fmt.Print("Enter search query: ")
	scanner.Scan()
	query := scanner.Text()

	contacts := addressBook.SearchContacts(query)
	if len(contacts) == 0 {
		fmt.Println("No contacts found matching your search.")
		return
	}

	fmt.Println("\nSearch Results:")
	for _, contact := range contacts {
		printContact(contact)
	}
}

func updateContact(scanner *bufio.Scanner, addressBook *models.AddressBook) {
	fmt.Print("Enter contact ID to update: ")
	scanner.Scan()
	id := scanner.Text()

	contact, err := addressBook.GetContact(id)
	if err != nil {
		fmt.Printf("Error finding contact: %v\n", err)
		return
	}

	fmt.Printf("Current contact: %s %s\n", contact.FirstName, contact.LastName)
	fmt.Print("Enter new first name (or press Enter to keep current): ")
	scanner.Scan()
	if firstName := scanner.Text(); firstName != "" {
		contact.FirstName = firstName
	}

	fmt.Print("Enter new last name (or press Enter to keep current): ")
	scanner.Scan()
	if lastName := scanner.Text(); lastName != "" {
		contact.LastName = lastName
	}

	fmt.Print("Enter new email (or press Enter to keep current): ")
	scanner.Scan()
	if email := scanner.Text(); email != "" {
		contact.Email = email
	}

	fmt.Print("Enter new phone (or press Enter to keep current): ")
	scanner.Scan()
	if phone := scanner.Text(); phone != "" {
		contact.Phone = phone
	}

	fmt.Print("Enter new address (or press Enter to keep current): ")
	scanner.Scan()
	if address := scanner.Text(); address != "" {
		contact.Address = address
	}

	if err := addressBook.UpdateContact(contact); err != nil {
		fmt.Printf("Error updating contact: %v\n", err)
		return
	}

	fmt.Println("Contact updated successfully!")
}

func deleteContact(scanner *bufio.Scanner, addressBook *models.AddressBook) {
	fmt.Print("Enter contact ID to delete: ")
	scanner.Scan()
	id := scanner.Text()

	if err := addressBook.DeleteContact(id); err != nil {
		fmt.Printf("Error deleting contact: %v\n", err)
		return
	}

	fmt.Println("Contact deleted successfully!")
}

func generateTestData(addressBook *models.AddressBook) {
	gen := generator.NewGenerator()
	contacts := gen.GenerateContacts(10)

	for _, contact := range contacts {
		if err := addressBook.AddContact(contact); err != nil {
			fmt.Printf("Error adding test contact: %v\n", err)
			return
		}
	}

	fmt.Println("Generated 10 test contacts successfully!")
}

func printContact(contact *models.Contact) {
	fmt.Printf("\nID: %s\n", contact.ID)
	fmt.Printf("Name: %s %s\n", contact.FirstName, contact.LastName)
	fmt.Printf("Email: %s\n", contact.Email)
	fmt.Printf("Phone: %s\n", contact.Phone)
	fmt.Printf("Address: %s\n", contact.Address)
	fmt.Printf("Created: %s\n", contact.CreatedAt.Format(time.RFC3339))
	fmt.Printf("Updated: %s\n", contact.UpdatedAt.Format(time.RFC3339))
}
