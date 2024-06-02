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

	_, err := rp.Subscribe("testAddress")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if models.MemoryInstance.Addresses["testAddress"] != true {
		t.Errorf("Expected true, got %v", models.MemoryInstance.Addresses["testAddress"])
	}
}

func TestGetTransactions(t *testing.T) {
	models.MemoryInitialize()
	storage.NewMemoryStorage()

	rp := &RestfulParser{}

	models.MemoryInstance.Addresses["testAddress"] = true
	expectedTransaction := []models.Transaction{{From: "A", To: "B", Value: "1", TcType: "type", Hash: "hash"}}
	models.MemoryInstance.Transactions["testAddress"] = expectedTransaction

	transactions, err := rp.GetTransactions("testAddress")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(transactions, expectedTransaction) {
		t.Errorf("Expected %v, got %v", expectedTransaction, transactions)
	}
}
