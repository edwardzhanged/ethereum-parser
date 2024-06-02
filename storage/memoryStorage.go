package storage

import (
	models "ethereum-parser/models"
	"fmt"
	"log"
)

// MemoryStorage is an in-memory implementation of the Storage interface.
type MemoryStorage struct {
	data *models.Memory
}

var MemoryStorageInstance *MemoryStorage

func NewMemoryStorage() {
	MemoryStorageInstance = &MemoryStorage{
		data: models.MemoryInstance,
	}
}

func (m *MemoryStorage) SaveSubscriber(subscriber string) error {
	if _, exists := m.data.Addresses[subscriber]; exists {
		return fmt.Errorf("address %s already exists", subscriber)
	}
	m.data.Addresses[subscriber] = true
	log.Printf("Saved subscriber %s", subscriber)
	return nil
}

func (m *MemoryStorage) GetSubscribers() (map[string]bool, error) {
	return m.data.Addresses, nil
}

func (m *MemoryStorage) SaveTransaction(transaction models.Transaction, address string) error {
	m.data.Transactions[address] = append(m.data.Transactions[address], transaction)
	log.Printf("Saved transaction %v for address %s", transaction, address)
	return nil
}

func (m *MemoryStorage) GetTransactions(address string) ([]models.Transaction, error) {
	if _, exists := m.data.Transactions[address]; !exists {
		return []models.Transaction{}, nil
	}
	return m.data.Transactions[address], nil
}
