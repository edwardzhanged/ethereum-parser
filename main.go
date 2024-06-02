package main

import (
	global "ethereum-parser/global"
	models "ethereum-parser/models"
	servers "ethereum-parser/servers"
	services "ethereum-parser/services"
	"ethereum-parser/storage"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile("parser.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer file.Close()
	log.SetOutput(file)

	// Initalize
	global.GlobalInitialize()
	models.MemoryInitialize()
	storage.NewMemoryStorage()
	services.RestfulParserInitialize()

	fmt.Println("Server is running on port http://127.0.0.1:8080")
	servers.RestfulServerInitialize()
}
