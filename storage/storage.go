package storage

import "ethereum-parser/models"

type Storage interface {
	SaveLatestBlock(number string) error
	SaveSubscriber(subscriber string) error
	SaveTransactions(transactions []string) error
	GetSubscribers() ([]string, error)
	GetTransactions(address string) ([]models.Transaction, error)
}
