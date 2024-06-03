package services

import (
	"ethereum-parser/models"
	storage "ethereum-parser/storage"
	utils "ethereum-parser/utils"
	"reflect"
	"testing"
)

func TestGetCurrentBlock(t *testing.T) {
	rp := &RestfulParser{
		GetLatestBlock: func() ([]utils.Transaction, string, error) {
			return nil, "0x10", nil
		},
	}

	block, err := rp.GetCurrentBlock()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// assert block is int of 0x10
	if block != 16 {
		t.Errorf("Expected 16, got %v", block)
	}
}

func TestSubscribe(t *testing.T) {
	models.MemoryInitialize()
	storage.NewMemoryStorage()
	rp := &RestfulParser{}

	_, err := rp.Subscribe("testaddress")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if models.MemoryInstance.Addresses["testaddress"] != true {
		t.Errorf("Expected true, got %v", models.MemoryInstance.Addresses["testaddress"])
	}
}

func TestGetTransactions(t *testing.T) {
	models.MemoryInitialize()
	storage.NewMemoryStorage()

	rp := &RestfulParser{}

	models.MemoryInstance.Addresses["testaddress"] = true
	expectedTransaction := []models.Transaction{{From: "a", To: "b", Value: "1", TcType: "type", Hash: "hash"}}
	models.MemoryInstance.Transactions["testaddress"] = expectedTransaction

	transactions, err := rp.GetTransactions("testaddress")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(transactions, expectedTransaction) {
		t.Errorf("Expected %v, got %v", expectedTransaction, transactions)
	}
}

func TestFilterTransactionsForSubscribers(t *testing.T) {
	models.MemoryInitialize()
	storage.NewMemoryStorage()

	transactions := []utils.Transaction{
		{From: "a", To: "b", Value: "1", Hash: "hash1"},
		{From: "b", To: "c", Value: "2", Hash: "hash2"},
	}
	subscribers := map[string]bool{
		"a": true,
		"c": true,
	}

	filterTransactionsForSubscribers(transactions, subscribers)

	if models.MemoryInstance.Transactions["a"][0].Hash != "hash1" {
		t.Errorf("Expected %v, got %v", "hash1", models.MemoryInstance.Transactions["a"][0].Hash)
	}
	if models.MemoryInstance.Transactions["c"][0].Hash != "hash2" {
		t.Errorf("Expected %v, got %v", "hash2", models.MemoryInstance.Transactions["c"][0].Hash)
	}
	if models.MemoryInstance.Transactions["a"][0].TcType != "outbound" {
		t.Errorf("Expected %v, got %v", "outbound", models.MemoryInstance.Transactions["a"][0].TcType)
	}
	if models.MemoryInstance.Transactions["c"][0].TcType != "inbound" {
		t.Errorf("Expected %v, got %v", "inbound", models.MemoryInstance.Transactions["c"][0].TcType)
	}
}
