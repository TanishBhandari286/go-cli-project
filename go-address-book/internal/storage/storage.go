package storage

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/rushi/address-book-cli/internal/models" //local path of my macos machine
)

type Storage interface {
	Save(addressBook *models.AddressBook) error
	Load() (*models.AddressBook, error)
}

type CSVStorage struct {
	filepath string
	cache    *models.AddressBook
	mu       sync.RWMutex
}

func NewCSVStorage(filepath string) *CSVStorage {
	return &CSVStorage{
		filepath: filepath,
		cache:    models.NewAddressBook(),
	}
}

func (s *CSVStorage) Save(addressBook *models.AddressBook) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := os.MkdirAll(filepath.Dir(s.filepath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	file, err := os.Create(s.filepath)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %w", err)
	}
	defer file.Close()

	buffered := bufio.NewWriter(file)
	writer := csv.NewWriter(buffered)
	defer writer.Flush()
	defer buffered.Flush()

	if err := writer.Write([]string{"ID", "FirstName", "LastName", "Email", "Phone", "Address", "CreatedAt", "UpdatedAt"}); err != nil {
		return fmt.Errorf("failed to write CSV header: %w", err)
	}

	contacts := addressBook.GetAllContacts()
	batchSize := 100
	for i := 0; i < len(contacts); i += batchSize {
		end := i + batchSize
		if end > len(contacts) {
			end = len(contacts)
		}

		for _, contact := range contacts[i:end] {
			record := []string{
				contact.ID,
				contact.FirstName,
				contact.LastName,
				contact.Email,
				contact.Phone,
				contact.Address,
				contact.CreatedAt.Format(time.RFC3339),
				contact.UpdatedAt.Format(time.RFC3339),
			}
			if err := writer.Write(record); err != nil {
				return fmt.Errorf("failed to write contact to CSV: %w", err)
			}
		}
		writer.Flush()
		if err := writer.Error(); err != nil {
			return fmt.Errorf("failed to flush CSV writer: %w", err)
		}
	}

	s.cache = addressBook
	return nil
}

func (s *CSVStorage) Load() (*models.AddressBook, error) {
	s.mu.RLock()
	if s.cache != nil && len(s.cache.GetAllContacts()) > 0 {
		s.mu.RUnlock()
		return s.cache, nil
	}
	s.mu.RUnlock()

	s.mu.Lock()
	defer s.mu.Unlock()

	file, err := os.Open(s.filepath)
	if err != nil {
		if os.IsNotExist(err) {
			s.cache = models.NewAddressBook()
			return s.cache, nil
		}
		return nil, fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	buffered := bufio.NewReader(file)
	reader := csv.NewReader(buffered)

	_, err = reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV header: %w", err)
	}

	addressBook := models.NewAddressBook()
	batchSize := 100
	records := make([][]string, 0, batchSize)

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		records = append(records, record)

		if len(records) >= batchSize {
			if err := processBatch(addressBook, records); err != nil {
				return nil, err
			}
			records = records[:0]
		}
	}

	if len(records) > 0 {
		if err := processBatch(addressBook, records); err != nil {
			return nil, err
		}
	}

	s.cache = addressBook
	return addressBook, nil
}

func processBatch(addressBook *models.AddressBook, records [][]string) error {
	for _, record := range records {
		createdAt, err := time.Parse(time.RFC3339, record[6])
		if err != nil {
			return fmt.Errorf("failed to parse CreatedAt: %w", err)
		}

		updatedAt, err := time.Parse(time.RFC3339, record[7])
		if err != nil {
			return fmt.Errorf("failed to parse UpdatedAt: %w", err)
		}

		contact := &models.Contact{
			ID:        record[0],
			FirstName: record[1],
			LastName:  record[2],
			Email:     record[3],
			Phone:     record[4],
			Address:   record[5],
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}

		if err := addressBook.AddContact(contact); err != nil {
			return fmt.Errorf("failed to add contact: %w", err)
		}
	}
	return nil
}
