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

	err := MemoryStorageInstance.SaveSubscriber("testaddress")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	err = MemoryStorageInstance.SaveSubscriber("testaddress")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestGetSubscribers(t *testing.T) {
	models.MemoryInitialize()
	NewMemoryStorage()

	MemoryStorageInstance.data.Addresses["testaddress"] = true

	subscribers, err := MemoryStorageInstance.GetSubscribers()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if subscribers["testaddress"] != true {
		t.Errorf("Expected true, got %v", subscribers["testaddress"])
	}
}

func TestSaveTransaction(t *testing.T) {
	models.MemoryInitialize()
	NewMemoryStorage()

	transaction := models.Transaction{
		From:   "a",
		To:     "b",
		Value:  "1",
		TcType: "inbound",
		Hash:   "hashxx",
	}

	err := MemoryStorageInstance.SaveTransaction(transaction, "testaddress")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if !reflect.DeepEqual(transaction, MemoryStorageInstance.data.Transactions["testaddress"][0]) {
		t.Errorf("Expected %v, got %v", transaction, MemoryStorageInstance.data.Transactions["testaddress"][0])
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
	MemoryStorageInstance.SaveTransaction(transaction, "testaddress2")

	transactions, err := MemoryStorageInstance.GetTransactions("testaddress2")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if !reflect.DeepEqual(transaction, transactions[0]) {
		t.Errorf("Expected %v, got %v", transaction, transactions[0])
	}
}
