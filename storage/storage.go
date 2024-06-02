package storage

import "ethereum-parser/models"

type Storage interface {
	SaveSubscriber(subscriber string) error
	SaveTransaction(transactions []string) error
	GetSubscribers() ([]string, error)
	GetTransactions(address string) ([]models.Transaction, error)
}
