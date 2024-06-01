package servers

import (
	"encoding/json"
	"ethereum-parser/services"
	"fmt"
	"net/http"
)

func InitrestfulServer() {
	http.HandleFunc("/currentBlock", func(w http.ResponseWriter, r *http.Request) {
		currentBlock, err := services.MyParser.GetCurrentBlock()
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			json.NewEncoder(w).Encode(map[string]any{"error": fmt.Sprintf("Error getting current block: %v", err), "success": false})
			return
		}
		json.NewEncoder(w).Encode(map[string]any{"currentBlock": currentBlock, "success": true})
	})

	http.HandleFunc("/subscribe", func(w http.ResponseWriter, r *http.Request) {
		topic := r.URL.Query().Get("address")
		if topic == "" {
			http.Error(w, "Missing adrress parameter", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		success, err := services.MyParser.Subscribe(topic)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]any{"error": fmt.Sprintf("Error subscribing: %v", err), "success": success})
			return
		}
		json.NewEncoder(w).Encode(map[string]any{"success": success})

	})

	http.HandleFunc("/getTransactions", func(w http.ResponseWriter, r *http.Request) {
		address := r.URL.Query().Get("address")
		if address == "" {
			http.Error(w, "Missing address parameter", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		transactions, err := services.MyParser.GetTransactions(address)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{"error": fmt.Sprintf("Error getting transactions: %v", err), "success": false})
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"transactions": transactions, "success": true})
	})

	http.ListenAndServe(":8080", nil)
}