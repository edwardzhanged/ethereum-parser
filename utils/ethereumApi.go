package utils

import (
	"bytes"
	"encoding/json"
	"ethereum-parser/global"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  struct {
		Transactions []Transaction `json:"transactions"`
		Hash         string        `json:"hash"`
		Number       string        `json:"number"`
	} `json:"result"`
}

type Transaction struct {
	BlockHash   string `json:"blockHash"`
	BlockNumber string `json:"blockNumber"`
	Hash        string `json:"hash"`
	Value       string `json:"value"`
	From        string `json:"from"`
	To          string `json:"to"`
}

func GetLatestBlock() (transactions []Transaction, number string, err error) {
	reqBody := bytes.NewBuffer([]byte(`{
		"jsonrpc": "2.0",
		"method": "eth_getBlockByNumber",
		"params": [
			"latest",
			true
		],
		"id": 1
	}`))

	resp, err := http.Post(global.GlobalConfig.Endpoint, "application/json", reqBody)
	if err != nil {
		return nil, "", fmt.Errorf("failed to send request: %v", err)
	}

	// TODO: Handle response here
	// resp.Body.Close()
	responseStruct := &Response{}
	err = json.NewDecoder(resp.Body).Decode(responseStruct)
	if err != nil {
		log.Printf("Failed to decode response: %v", err)
	}

	// Now you can access the fields in the response
	// log.Printf("transactions: %s, hash: %s", responseStruct.Result.Transactions, responseStruct.Result.Hash)
	transactions, number = responseStruct.Result.Transactions, responseStruct.Result.Number

	return transactions, number, nil
}
