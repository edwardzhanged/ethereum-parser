package storage

import (
	models "ethereum-parser/models"
	"reflect"
	"testing"
)

func TestNewMemoryStorage(t *testing.T) {
	NewMemoryStorage()
	if MemoryStorageInstance == nil {
		t.Errorf("Expected MemoryStorageInstance not to be nil")
	}
}

func TestSaveSubscriber(t *testing.T) {
	models.MemoryInitialize()
	NewMemoryStorage()

	err := MemoryStorageInstance.SaveSubscriber("testAddress")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	err = MemoryStorageInstance.SaveSubscriber("testAddress")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestGetSubscribers(t *testing.T) {
	models.MemoryInitialize()
	NewMemoryStorage()

	MemoryStorageInstance.data.Addresses["testAddress"] = true

	subscribers, err := MemoryStorageInstance.GetSubscribers()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if subscribers["testAddress"] != true {
		t.Errorf("Expected true, got %v", subscribers["testAddress"])
	}
}

func TestSaveTransaction(t *testing.T) {
	models.MemoryInitialize()
	NewMemoryStorage()

	transaction := models.Transaction{
		From:   "A",
		To:     "B",
		Value:  "1",
		TcType: "inbound",
		Hash:   "hashxx",
	}

	err := MemoryStorageInstance.SaveTransaction(transaction, "testAddress")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if !reflect.DeepEqual(transaction, MemoryStorageInstance.data.Transactions["testAddress"][0]) {
		t.Errorf("Expected %v, got %v", transaction, MemoryStorageInstance.data.Transactions["testAddress"][0])
	}
}

func TestGetTransactions(t *testing.T) {
	models.MemoryInitialize()
	NewMemoryStorage()

	transaction := models.Transaction{
		From:   "c",
		To:     "d",
		Value:  "2",
		TcType: "inbound",
		Hash:   "hashyy",
	}
	MemoryStorageInstance.SaveTransaction(transaction, "testAddress2")

	transactions, err := MemoryStorageInstance.GetTransactions("testAddress2")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if !reflect.DeepEqual(transaction, transactions[0]) {
		t.Errorf("Expected %v, got %v", transaction, transactions[0])
	}
}
