package services

import "ethereum-parser/models"

type Parser interface {
	// last parsed block
	GetCurrentBlock() (int, error)
	// add address to observer
	Subscribe(address string) (bool, error)
	// list of inbound or outbound transactions for an address
	GetTransactions(address string) []models.Transaction
}
