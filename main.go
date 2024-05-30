package main

import (
	"bytes"
	"encoding/json"
	global "ethereum-parser/global"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  struct {
		Transactions []any  `json:"transactions"`
		Hash         string `json:"hash"`
	} `json:"result"`
}

func main() {
	global.Initialize()

	go func() {
		for {
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
				log.Printf("Failed to send request: %v", err)
				continue
			}

			// TODO: Handle response here
			// resp.Body.Close()
			responseStruct := &Response{}
			err = json.NewDecoder(resp.Body).Decode(responseStruct)
			if err != nil {
				log.Printf("Failed to decode response: %v", err)
				continue
			}

			// Now you can access the fields in the response
			log.Printf("transactions: %s, hash: %s", responseStruct.Result.Transactions, responseStruct.Result.Hash)
			time.Sleep(5 * time.Second)
		}
	}()

	// TODO: Other code here
	select {}
}

// TODO:
//1. 启动一个http client,接受subscribe, get tranactions
//2. 实现parase 接口的函数，被http调用，是不是要controller。impl
//3.保证内存安全，多个goroutine访问内存
