package services

import (
	"context"
	"ethereum-parser/models"
	storage "ethereum-parser/storage"
	utils "ethereum-parser/utils"
	"strconv"
	"strings"
	"sync"

	"fmt"
	"log"
	"time"
)

type RestfulParser struct {
	ctx context.Context
}

var (
	MyParser *RestfulParser
	mu       sync.Mutex
)

func (rp *RestfulParser) GetCurrentBlock() (int, error) {
	rp.ctx = context.Background()
	_, currentBlockStr, err := utils.GetLatestBlock()
	if err != nil {
		log.Println(err)
		return 0, fmt.Errorf("failed to request latest block: %v", err)
	}

	var currentBlockHex string
	if strings.HasPrefix(currentBlockStr, "0x") {
		currentBlockHex = currentBlockStr[2:]
	}

	currentBlockInt, err := strconv.ParseInt(currentBlockHex, 16, 64)
	if err != nil {
		log.Printf("failed to parse block number: %v", err)
		return 0, fmt.Errorf("failed to parse block number: %v", err)
	}
	return int(currentBlockInt), nil
}

func (rp *RestfulParser) Subscribe(address string) (bool, error) {
	err := storage.MemoryStorageInstance.SaveSubscriber(address)
	if err != nil {
		log.Printf("failed to save subscriber: %v", err)
		return false, err
	}
	return true, nil
}

func (rp *RestfulParser) GetTransactions(address string) ([]models.Transaction, error) {
	subscribers, _ := storage.MemoryStorageInstance.GetSubscribers()
	if _, exists := subscribers[address]; !exists {
		return nil, fmt.Errorf("address %s is not subscribed", address)
	}
	transactions, _ := storage.MemoryStorageInstance.GetTransactions(address)
	return transactions, nil
}

func (rp *RestfulParser) startListeners() {
	const numListeners = 5
	const delay = time.Second

	for i := 0; i <= numListeners; i++ {
		go func(id int) {
			for {
				singleListener()

				time.Sleep(1 * time.Second)
			}
		}(i)
		time.Sleep(delay)
	}
}

func singleListener() {
	log.Println("Starting latest transactions listener")

	transactions, number, err := utils.GetLatestBlock()
	if err != nil {
		log.Printf("failed to get latest block: %v", err)
	}
	log.Printf("On latest block, number: %s", number)
	log.Printf("On latest block, Transactions: %v", transactions)

	subscribers, _ := storage.MemoryStorageInstance.GetSubscribers()
	for _, transaction := range transactions {
		mu.Lock()
		if _, exists := models.MyMemory.RecordedTxHashes[transaction.Hash]; exists {
			mu.Unlock()
			continue
		}
		models.MyMemory.RecordedTxHashes[transaction.Hash] = true
		mu.Unlock()

		if _, exists := subscribers[transaction.From]; exists {
			saveTransaction := models.Transaction{
				From:   transaction.From,
				To:     transaction.To,
				Value:  transaction.Value,
				TcType: "outbound",
				Hash:   transaction.Hash,
			}
			mu.Lock()
			storage.MemoryStorageInstance.SaveTransaction(saveTransaction, transaction.From)
			mu.Unlock()
		}
		if _, exists := subscribers[transaction.To]; exists {
			saveTransaction := models.Transaction{
				From:   transaction.From,
				To:     transaction.To,
				Value:  transaction.Value,
				TcType: "inbound",
				Hash:   transaction.Hash,
			}
			mu.Lock()
			storage.MemoryStorageInstance.SaveTransaction(saveTransaction, transaction.To)
			mu.Unlock()
		}
	}
	time.Sleep(time.Second)

}

func InitRestfulParser() {
	MyParser = &RestfulParser{}
	go MyParser.startListeners()
}
